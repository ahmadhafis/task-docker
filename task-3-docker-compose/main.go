package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/connect", connectHandler)

	log.Printf("Server starting on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone hit me!")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Wise Words</h1><p>The journey of a thousand miles begins with a single step.</p>")
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Connection parameters
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed connect to db: %v", err)
		fmt.Fprintf(w, "<h1>Failed to connect to database</h1>")
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Printf("Failed connect to db: %v", err)
		fmt.Fprintf(w, "<h1>Failed to connect to database</h1>")
		return
	}

	// Insert current timestamp
	_, err = db.Exec("INSERT INTO ahmadhafis_access_log (timestamp) VALUES ($1)", time.Now())
	if err != nil {
		log.Printf("Failed to insert data: %v", err)
		fmt.Fprintf(w, "<h1>Failed to insert data</h1>")
		return
	}

	log.Println("Success connect to db")
	fmt.Fprintf(w, "<h1>Successfully connected to database and inserted timestamp</h1>")
}
