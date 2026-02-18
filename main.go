package main

import (
	"fmt"
	"net/http"
	"yboostS2/routes"
)

func main() {
	http.HandleFunc("/", routes.HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":80", nil)
}
