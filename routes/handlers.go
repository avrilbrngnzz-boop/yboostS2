package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Text     string `gorm:"not null"`
	Author   string
	Category string
}

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("cat")
	var quotes []Quote

	if cat == "" {
		DB.Find(&quotes)
	} else {
		DB.Where("category = ?", cat).Find(&quotes) // Filtre par catégorie
	}

	tmpl := template.Must(template.ParseFiles("templates/start.html"))
	tmpl.Execute(w, quotes)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newQ := Quote{
			Text:     r.FormValue("text"),
			Author:   r.FormValue("author"),
			Category: r.FormValue("category"),
		}
		DB.Create(&newQ)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func DelHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("--- DEBUG SUPPRESSION ---")
	fmt.Println("ID reçu depuis l'URL :", id)

	if id != "" {
		tx := DB.Unscoped().Delete(&Quote{}, id)

		fmt.Println("Erreur éventuelle :", tx.Error)
		fmt.Println("Nombre de lignes impactées :", tx.RowsAffected)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
