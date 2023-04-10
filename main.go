package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return
		}
		w.Write([]byte("Hello guys!"))
	})

	log.Println("run server")
	http.ListenAndServe(GetPort(), nil)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		port = ":" + port
		return port
	}
	return ":3000"
}
