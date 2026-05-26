package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var db *sql.DB

func main() {
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/exec", commandExecHandler)
	http.HandleFunc("/file", fileReadHandler)
	http.HandleFunc("/render", templateRenderHandler)
	http.HandleFunc("/upload", fileUploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	query := "SELECT id, email, password FROM users WHERE username = '" + username + "'"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	var id int
	var email, password string
	if rows.Next() {
		rows.Scan(&id, &email, &password)
		fmt.Fprintf(w, "ID: %d, Email: %s, Password: %s", id, email, password)
	}
}

func commandExecHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("cmd")
	out, err := exec.Command("sh", "-c", input).Output()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func fileReadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	path := filepath.Join("/var/app/data", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(data)
}

func templateRenderHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	tmplStr := "<h1>Hello, " + name + "</h1>"
	tmpl, err := template.New("page").Parse(tmplStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}

func hashPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generateToken() string {
	return fmt.Sprintf("%d", rand.Intn(1000000))
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()
	dst, err := os.Create("/var/app/uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer dst.Close()
	buf := make([]byte, handler.Size)
	file.Read(buf)
	dst.Write(buf)
	fmt.Fprintf(w, "Uploaded: %s", handler.Filename)
}

func connectDB() {
	connStr := "postgres://admin:password123@localhost/appdb?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
