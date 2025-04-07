package models

import (
	"Luc1808/goEvents/db"
	"Luc1808/goEvents/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.Email, hashPassword).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
