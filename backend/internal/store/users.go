package store

type User struct {
	UserName string `binding:"required,min=3,max=10"`
	Email string	`binding:"required,min=5,max=30"`
	Password string	`binding:"required,min=7,max=32"`
}

var Users []*User