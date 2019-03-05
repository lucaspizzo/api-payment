package forms

type PaymentForm struct {
	BaseForm
	AccountID uint64  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

func (p *PaymentForm) Validate() bool {
	if p.AccountID == 0 {
		p.addError("accountID", "invalid accountId")
	}

	if p.Amount <= 0 {
		p.addError("amount", "invalid amount")
	}

	return p.IsValid()
}

func ValidatePaymentFormList(payments *[]*PaymentForm) bool {

	if len(*payments) == 0 {
		return false
	}

	for _, payment := range *payments {
		if !payment.Validate() {
			return false
		}
	}
	return true
}
