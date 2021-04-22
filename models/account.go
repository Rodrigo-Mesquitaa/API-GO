package models

import (
	"time"
)

type Account struct {
	AccountID uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Ballance  uint      `json:"ballance"`
	CreatedAt time.Time `json:"created_at"`
}
