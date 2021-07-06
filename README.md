# [Sqreen](https://www.sqreen.com/)'s Go Damn Vulnerable Web App

This Go web server is a vulnerable application demonstration, protected by
Sqreen.

It currently includes the following vulnerabilities:

- [x] SQL injection: `/products` accepts a URL-query parameter `category` that
  is injected into the SQL query (eg. `/products?category=all%27%20UNION%20SELECT%20*%20FROM%20user%27`)
  .
- [x] Shell injection
- [ ] NoSQL injection
- [ ] Server-Side Request Forgery

The web app comes with Sqreen for Go which can be enabled by running a valid
Sqreen configuration that can be obtained at <https://my.sqreen.com/>. Once
enabled, the agent should protect the application according to the application
security configuration you enabled.

## Quick Start

The pre-compiled [go-dvwa](https://github.com/sqreen/go-dvwa/packages/494452)
docker image can be used to simply run the web application. The HTTP server
listens the TCP address `0.0.0.0:8080` so you can expose it with docker:

```console
$ docker run -it -p 8080:8080 go-dvwa
```

The vulnerable web app starts regardless of Sqreen's agent. It will start when
having a valid configuration with Sqreen credentials you can get at
<https://my.sqree.com/>. You can pass them using container's environment
variables:

```console
$ docker run -it -p 8080:8080 -e SQREEN_TOKEN=<token> -e SQREEN_APP_NAME="Go DVWA" go-dvwa
```

The web app vulnerabilities should be now blocked by Sqreen :-)

<p align="center">
<img width="60%" src="./doc/images/blocking-page-with-gopher.png" alt="Sqreen for Go" title="Sqreen for Go" />
</p>

## Compile from sources

### With docker builder

The simplest way to build this repository is by using the latest docker builder
which can take a git repository source. Simply run the following command to
build the latest `go-dvwa` docker image of this repository:

```console
$ docker builder build github.com/sqreen/go-dvwa.git
```

Once built, you can simply run the image and pass the Sqreen configuration to
the container via environment variables:

```console
$ docker run -e SQREEN_TOKEN=<token> -e SQREEN_APP_NAME="Go DVWA" -p 8080:8080 go-dvwa
```

The Go web application is now running and you can access it
at <http://127.0.0.1:8080/>.

### From sources

Clone the repository and use the Makefile:

```console
$ make
```

Once compiled, you can execute the binary file `dvwa`. Sqreen's agent
configuration can then be passed by file or environment variable.

```console
$ ./dvwa
```

The Go web application is now running and you can access it
at <http://127.0.0.1:8080/>.

Note that the docker image can be also built using the Makefile:

```console
$ make image
```

Cf. the previous docker image instructions to read how to start the container.
