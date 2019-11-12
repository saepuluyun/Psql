package main

import "time"

//Users
type Users struct {
	ID         int
	Username   string
	Password   string
	First_name string
	Last_name  string
	Email      string
	Date_users time.Time
}
