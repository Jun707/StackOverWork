package store

import (

)

type User struct {
	UserName string
	Email string
	Password string
}

var Users []*User