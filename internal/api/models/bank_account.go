package models

type BankAccount struct {
	ID            uint   `gorm:"primaryKey" json:"-"`
	PlayerID      uint   `json:"-"`
	AccountName   string `json:"account_name"`
	AccountNumber string `gorm:"unique;not null" json:"account_number"`
	BankName      string `json:"bank_name"`
}
