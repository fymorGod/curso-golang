package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User
type User struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userToId(w, r, id)
	case r.Method == "GET":
		userAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry...")
	}
}

func userToId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.ID, &u.NAME)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(json))
}

func userAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.NAME)
		users = append(users, u)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(json))
}
