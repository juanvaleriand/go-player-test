package requests

type WalletRequest struct {
	Balance uint `json:"balance" binding:"required"`
}
