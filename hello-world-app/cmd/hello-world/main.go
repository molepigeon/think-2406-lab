package main

// Super lazy Go-based HTTP server

import (
	"fmt"
	"log"
	"net/http"
)

const bindAddress = ":8080"
const html = `
<head>
	<title>
		Tutorial Sample App
	</title>
</head>
<body>
	<h1>
		Hi there, I'm an app! Serving %q
	</h1>
</body>
`

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request for %q", r.RequestURI)
	fmt.Fprintf(w, html, r.RequestURI)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://icr.io/static/logo/icon_24px.png")
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", favicon)
	log.Printf("Start listening on %s", bindAddress)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
