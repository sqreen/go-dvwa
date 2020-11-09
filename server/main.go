// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sqreen/go-agent/sdk/middleware/sqhttp"
)

func main() {
	bin, err := os.Executable()
	if err != nil {
		log.Panic("could not get the executable filename:", err)
	}
	templateDir := filepath.Join(filepath.Dir(bin), "template")

	router := NewRouter(templateDir)

	addr := ":8080"
	log.Println("Serving application on", addr)
	err = http.ListenAndServe(addr, sqhttp.Middleware(router))
	if err != nil {
		log.Fatalln(err)
	}
}
