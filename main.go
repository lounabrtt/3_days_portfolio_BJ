package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"portfolio/database"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Structure pour le profil
type Contact struct {
	Name        string
	FirstName   string
	Email       string
	Phone	    string
	Description string
}

type Login struct {
	Id    string
	password string
}

type Language struct {
	Language string
}

type Logiciel struct {
	Logiciel string
}

type Projet struct {
	Projet string
}

// Page d'accueil
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	firstName := r.FormValue("firstname")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	message := r.FormValue("message")

	_, err := db.Exec("INSERT INTO contact (name, firstName, phone, email, message) VALUES (?, ?, ?, ?, ?)", name, firstName, phone, email, message)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données : %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")

	_, err := db.Exec("INSERT INTO login (id, password) VALUES (?, ?)", id, password)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données : %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addlanguageHandler(w http.ResponseWriter, r *http.Request) {
	language := r.FormValue("language")

	_, err := db.Exec("INSERT INTO language (language) VALUES (?)", language)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données : %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addlogicielHandler(w http.ResponseWriter, r *http.Request) {
	logiciel := r.FormValue("logiciel")

	_, err := db.Exec("INSERT INTO logiciel (logiciel) VALUES (?)", logiciel)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données : %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addprojetHandler(w http.ResponseWriter, r *http.Request) {
	projet := r.FormValue("projet")

	_, err := db.Exec("INSERT INTO projet (projet) VALUES (?)", projet)
	if err != nil {
		log.Printf("Erreur lors de l'insertion des données : %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func main() {
	var err error
	// Connexion à la base de données SQLite
	db, err = sql.Open("sqlite3", "./database/portfolio.db")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données : %v", err)
	}
	defer db.Close()

	// Création de la table si elle n'existe pas
	database.InitTables(db)

	// Routage des fichiers statiques (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes des pages HTML
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/login", loginHandler)


	// http.HandleFunc("/addlanguage", addlanguageHandler)
	// http.HandleFunc("/addlogiciel", addlogicielHandler)
	// http.HandleFunc("/addprojet", addprojetHandler)

	// http.HandleFunc("./templates/index.html", contactHandler)
	http.HandleFunc("./templates/login.html", loginHandler)

	// Lancement du serveur sur le port 8080
	log.Println("Serveur démarré sur http://localhost:8080")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
