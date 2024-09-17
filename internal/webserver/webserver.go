package webserver

import (
	"fmt"
	"net/http"
	"time"
)

type WSConfig struct {
}

func StartWebServer() {
	fmt.Println("Initializing webserver...")

	wsc := WSConfig{}

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", wsc.indexHandler)

	port := ":8080"

	// TODO: add a go routine to make a request to the server and print that webserver is live!
	go webserverPulse("http://localhost" + port)

	// Start the server
	fmt.Printf("Starting webserver...\n\n")
	http.ListenAndServe(port, nil)
}

func webserverPulse(url string) {
	c := 0

	for {
		if c > 10 {
			fmt.Printf("Failed to connect after %d attempts, closing connection.", c)
			return
		}

		time.Sleep(1 * time.Second)
		_, err := http.Get(url)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			c++
		} else {
			fmt.Println("*** live connection! ***")
			fmt.Printf("\nURL: %s", url)
			return
		}
	}
}
