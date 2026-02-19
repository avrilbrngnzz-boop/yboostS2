package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"yboostS2/routes"
)

func main() {
	db, err := InitDB()

	if err != nil {
		log.Fatal("Impossible de connecter la base :", err)
	}

	db.AutoMigrate(&routes.Quote{})

	routes.SetDB(db)

	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/add", routes.AddHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serveur lancé sur http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
