package models

import "golang.org/x/crypto/bcrypt"

type Player struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	FirsName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	IsActive bool   `json:"-"`
}

func (u *Player) SetPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash)
}

func (u *Player) ComparePasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return false
	}
	return true
}
