package entities

import "time"

type User struct {
	Entity
	Username    string
	DateDeleted *time.Time
	Calls       []Call
}
