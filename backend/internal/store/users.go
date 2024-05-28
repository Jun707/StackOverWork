package store

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserName string    `gorm:"size:10;not null" binding:"required,min=3,max=10"`
	Email    string    `gorm:"size:30;not null" binding:"required,min=5,max=30"`
	Password string    `gorm:"size:32;not null" binding:"required,min=7,max=32"`
}

func GetUsers() ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func AddUser(user *User) error {
	if err :=db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func Authenticate(email, password string) (*User, error) {
	user := new(User)
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	if password !=user.Password {
		return nil, errors.New("password not valid")
	}
	return user, nil
}