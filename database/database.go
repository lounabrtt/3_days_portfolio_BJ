package database

import (
	"database/sql"
	"fmt"
)

func InitTables(db *sql.DB) error {
    if err := createTableContact(db); err != nil {
        return fmt.Errorf("error creating contact table: %v", err)
    }
    if err := createTableExperience(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    if err := createTableCompetence(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    if err := createTableProjet(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    if err := createTableLogin(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    if err := createTableLanguage(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    if err := createTableLogiciel(db); err != nil {
        return fmt.Errorf("error creating post table: %v", err)
    }
    return nil
}

// createTableContact permet de créer la table contact dans la base de données
func createTableContact(db *sql.DB) error {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS contact (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        firstName TEXT NOT NULL,
        phone TEXT NOT NULL,
        email TEXT NOT NULL,
        message TEXT NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }

    func createTableLogin(db *sql.DB) error {
        createTableSQL := `
        CREATE TABLE IF NOT EXISTS contact (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            password TEXT NOT NULL
        );`

    _, err := db.Exec(createTableSQL)
        return err
    }

// createTableLanguage permet de créer la table language dans la base de données
func createTableLanguage(db *sql.DB) error {
    createTableSQL :=`
    CREATE TABLE IF NOT EXISTS language (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        language TEXT NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }

// createTableLogiciel permet de créer la table logiciel dans la base de données
func createTableLogiciel(db *sql.DB) error {
    createTableSQL :=`
    CREATE TABLE IF NOT EXISTS logiciel (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        logiciel TEXT NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }

// createTableEducation permet de créer la table education dans la base de données
func createTableExperience(db *sql.DB) error {
    createTableSQL :=`
    CREATE TABLE IF NOT EXISTS experience (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        institution TEXT NOT NULL,
        degree TEXT NOT NULL,
        year INTEGER NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }

// createTableCompetence permet de créer la table competence dans la base de données
func createTableCompetence(db *sql.DB) error {
    createTableSQL :=`
    CREATE TABLE IF NOT EXISTS competence (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        skill TEXT NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }

// createTableProjet permet de créer la table projet dans la base de données
    func createTableProjet(db *sql.DB) error {
    createTableSQL :=`
    CREATE TABLE IF NOT EXISTS projet (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        skill TEXT NOT NULL
    );`

    _, err := db.Exec(createTableSQL)
        return err
    }
