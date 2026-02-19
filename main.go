package main

import (
	"fmt"
	"net/http"
	"os"
	"yboostS2/routes"
)

func main() {
	db, err := InitDB()

	if err != nil {
		fmt.Println("ATTENTION : La base de données ne répond pas :", err)
		fmt.Println("Vérifie ton mot de passe Postgres ou si le service est lancé.")
	} else {
		db.AutoMigrate(&routes.Quote{})
		routes.SetDB(db)
	}

	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/add", routes.AddHandler)
	http.HandleFunc("/delete", routes.DelHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serveur lancé sur http://localhost:" + port)
	errServer := http.ListenAndServe(":"+port, nil)
	if errServer != nil {
		os.Exit(1)
	}
}
