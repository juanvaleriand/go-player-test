package requests

type BankAccountRequest struct {
	AccountName   string `json:"account_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
}
