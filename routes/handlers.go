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
	search := r.URL.Query().Get("search")
	var quotes []Quote

	query := DB

	if cat != "" {
		query = query.Where("category = ?", cat)
	}

	if search != "" {
		query = query.Where("author ILIKE ?", "%"+search+"%")
	}

	query.Find(&quotes)

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
