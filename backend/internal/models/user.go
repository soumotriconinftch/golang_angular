package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64        `json:"id"`
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Password PasswordData `json:"-"`
	IsAdmin  bool         `json:"-"`
}

type User1 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	User User
}

type PasswordData struct {
	text *string
	Hash []byte
}

func (p *PasswordData) Set(pass string) error {
	log.Println("entering hashing")
	cost := bcrypt.DefaultCost
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}
	p.text = &pass
	p.Hash = bytes
	log.Println("hashed successfully")
	return nil
}

func (u *User) ComparePassword(plaintext string) error {
	return bcrypt.CompareHashAndPassword(u.Password.Hash, []byte(plaintext))
}
