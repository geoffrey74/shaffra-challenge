package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("postgres", "user=postgres dbname=test sslmode=disable")

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	getChannel := make(chan []string)

	go func() {
		rows, err := db.Query("SELECT name FROM users")
		if err != nil {
			fmt.Fprintf(w, "Failed to get user: %v", err)
			return
		}
		defer rows.Close()

		var names []string
		for rows.Next() {
			var name string
			rows.Scan(&name)
			names = append(names, name)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(names)

		getChannel <- names
	}()

	retrievedUsers := <-getChannel
	fmt.Fprintf(w, "User: %s\n", retrievedUsers)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	createChannel := make(chan string)

	go func() {
		time.Sleep(5 * time.Second) // Simulate a long database operation

		username := r.URL.Query().Get("name")
		_, err := db.Exec("INSERT INTO users (name) VALUES ('" + username + "')")

		if err != nil {
			fmt.Fprintf(w, "Failed to create user: %v", err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("User created")

		createChannel <- username
	}()

	newUser := <-createChannel
	fmt.Fprintf(w, "User %s created successfully", newUser)
}
