package main

import (
	"API"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fs := http.FileServer(http.Dir("../View"))
	http.Handle("/", fs)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/createuser/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../View/CreateUser.html")
	})
	http.HandleFunc("/users/create", CreateUser)

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	db, err := sql.Open("sqlite3", "../DB/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []API.Users
	for rows.Next() {
		var user API.Users
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_At, &user.Updated_At)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "../DB/test.db")
	if err != nil {
		fmt.Println("Error SQL !")
		log.Fatal(err)
	}
	defer db.Close()

	// body, _ := ioutil.ReadAll(r.Body)
	// log.Println(string(body))

	var user API.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", r.Body)
		fmt.Println("Error Decode !")
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO users (username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Username, user.Password, user.Email, user.Created_At, user.Updated_At)
	if err != nil {
		fmt.Println("Error Insert !")
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user.Id = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
