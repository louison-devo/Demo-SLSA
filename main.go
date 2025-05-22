package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body><h1>SLSA Demo, DevoHack{1L0v3_S3cD3v0ps}</h1></body></html>")
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8084"
	fmt.Println("Serveur démarré sur http://localhost" + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
