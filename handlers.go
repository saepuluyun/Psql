package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleSaveUsers(w http.ResponseWriter, r *http.Request) {

	var err error
	// dari form
	// r.ParseForm()
	// params := r.PostForm

	// username := params.Get("username")
	// password := params.Get("password")
	// firts_name := params.Get("first_name")
	// last_name := params.Get("last_name")

	// date_usersStr := params.Get("date_users")
	// var date_users time.Time
	// date_users, err = time.Parse("2006-01-02", date_usersStr)

	// id := r.URL.Query().Get("id")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	date_usersStr := r.URL.Query().Get("date_users")
	var date_users time.Time
	date_users, err = time.Parse("2006-01-02", date_usersStr)

	insertUser(username, password, first_name, last_name, date_users)

	http.Redirect(w, r, "/", 302)

	fmt.Println(err)

}
