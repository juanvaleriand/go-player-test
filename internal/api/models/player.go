package models

import "time"

type Player struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	BankAccounts BankAccount `gorm:"foreignKey:PlayerID" json:"bank_accounts"`
	Wallet       Wallet      `gorm:"foreignKey:PlayerID" json:"wallet"`
}
