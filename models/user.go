package models

type User struct {
	Id      string `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
