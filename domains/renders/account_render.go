package renders

import "github.com/lucaspizzo/api-payment/domains"

type AmountResponse struct{
	Amount float64 `json:"amount"`
}

type AccountResponse struct {

	AvailableCreditLimit AmountResponse `json:"available_credit_limit"`

	AvailableWithdrawalLimit AmountResponse `json:"available_withdrawal_limit"`
}

func NewPaymentTypeResponse(account *domains.Account) *AccountResponse {
	availableCreditLimit := AmountResponse{Amount:account.AvailableCreditLimit}
	availableWithdrawalLimit := AmountResponse{Amount:account.AvailableWithdrawalLimit}

	model := &AccountResponse{
		AvailableCreditLimit:availableCreditLimit,
		AvailableWithdrawalLimit:availableWithdrawalLimit,
	}

	return model
}

func NewPaymentTypeResponseList(accounts *[]*domains.Account) *[]*AccountResponse {
	modelList := []*AccountResponse{}
	for _, paymentType := range *accounts {
		modelList = append(modelList, NewPaymentTypeResponse(paymentType))
	}

	return &modelList
}
