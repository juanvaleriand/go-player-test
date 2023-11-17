package models

type Wallet struct {
	ID       uint `gorm:"primaryKey" json:"-"`
	PlayerID uint `json:"-"`
	Balance  uint `json:"balance"`
}
