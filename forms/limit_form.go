package forms

type AmountForm struct{
	Amount float64 `json:"amount"`
}

type LimitForm struct {
	BaseForm
	AccountID uint64
	AvailableCreditLimit AmountForm `json:"available_credit_limit"`
	AvailableWithdrawalLimit AmountForm `json:"available_withdrawal_limit"`
}

func (c *LimitForm) Validate() bool {
	if c.AccountID == 0 {
		c.addError("accountID", "invalid accountId")
	}

	return c.IsValid()
}
