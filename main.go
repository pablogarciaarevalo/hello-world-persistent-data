package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// BytesToString Convert []byte to string
func BytesToString(data []byte) string {
	return string(data[:])
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello world!\n")
	fmt.Fprintf(w, "The pod name is Hostname: %s\n", host)

	out, err := exec.Command("sh", "-c", "mount | grep /data").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := BytesToString(out)
	fmt.Fprintf(w, "The persistent volume used has been mounted on: %s\n", output)
}
