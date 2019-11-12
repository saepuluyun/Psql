package main

import (
	"fmt"
	"time"

	"github.com/lib/pq"
)

func getUsers(usersID int) (Users, error) {
	//List
	res := Users{}

	var id int
	var username string
	var password string
	var first_name string
	var last_name string
	var date_users pq.NullTime

	err := db.QueryRow(`SELECT id, username, password, first_name, last_name, date_users from users where id =$1`, usersID).Scan(&id, &username, &password, &first_name, &last_name, &date_users)
	if err == nil {
		res = Users{ID: id, Username: username, Password: password, First_name: first_name, Last_name: last_name, Date_users: date_users.Time}
	}

	return res, err
}

func insertUser(username string, password string, first_name string, last_name string, date_users time.Time) (int, error) {
	//Create
	var usersID int
	err := db.QueryRow(`INSERT INTO users(id, username, password, first_name, last_name, date_users) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, 1, username, password, first_name, last_name, date_users).Scan(&usersID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", usersID)
	return usersID, err
}

func updateUser(id int, username string, password string, first_name string, last_name string, date_users time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE users set username=$1, password=$2, first_name=$3, last_name=$4, date_users=&5 where id=$5 RETURNING id`, username, password, first_name, last_name, date_users, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}
