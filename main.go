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

// Structure pour le contact
type Contact struct {
	Name        string
	FirstName   string
	Email       string
	Phone	    string
	Description string
}

type Login struct {
	id    string
	password string
}

type Langage struct {
	Langage string
	url       string

}

type Logiciel struct {
	Logiciel string
	url1       string
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

func addlangageHandler(w http.ResponseWriter, r *http.Request) {
    langage := r.FormValue("langage")
    url := r.FormValue("url") // Assurez-vous d'avoir une entrée "url" dans le formulaire HTML

    _, err := db.Exec("INSERT INTO langage (langage, url) VALUES (?, ?)", langage, url)
    if err != nil {
        log.Printf("Erreur lors de l'insertion des données : %v", err)
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addlogicielHandler(w http.ResponseWriter, r *http.Request) {
    logiciel := r.FormValue("logiciel")
    url1 := r.FormValue("url1") // Assurez-vous d'avoir une entrée "url" dans le formulaire HTML

    _, err := db.Exec("INSERT INTO logiciel (logiciel, url1) VALUES (?, ?)", logiciel, url1)
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

// func adminHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseFiles("./templates/admin.html")
// 	if err != nil {
// 		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, nil)
// }

func checkTables(db *sql.DB) {
    rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
    if err != nil {
        log.Fatalf("Erreur lors de la vérification des tables : %v", err)
    }
    defer rows.Close()

    log.Println("Tables présentes dans la base de données :")
    for rows.Next() {
        var tableName string
        err := rows.Scan(&tableName)
        if err != nil {
            log.Fatalf("Erreur lors de la lecture des noms de tables : %v", err)
        }
        log.Println(tableName)
    }
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
	checkTables(db)
	// Routage des fichiers statiques (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes des pages HTML
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/login", loginHandler)


	http.HandleFunc("/addlangage", addlangageHandler)
	http.HandleFunc("/addlogiciel", addlogicielHandler)
	http.HandleFunc("/addprojet", addprojetHandler)

	// http.HandleFunc("./templates/index.html", contactHandler)
	http.HandleFunc("./templates/login.html", loginHandler)
	// http.HandleFunc("./templates/admin.html", adminHandler)

	// Lancement du serveur sur le port 8080
	log.Println("Serveur démarré sur http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
