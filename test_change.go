package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Nightly E2E — intentional security issues for Pullminder to detect

var awsAccessKey = "AKIAIOSFODNN7EXAMPLE"
var dbPassword = "super_secret_password_123"

func unsafeQuery(db *sql.DB, userInput string) error {
	query := "SELECT * FROM users WHERE name = '" + userInput + "'"
	_, err := db.Exec(query)
	return err
}

func unsafeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name)
}
