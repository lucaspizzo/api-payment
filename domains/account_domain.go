package domains

type Account struct {
	BaseDomain
	AccountID                uint64 `sql:"primary_key;"`
	AvailableCreditLimit     float64
	AvailableWithdrawalLimit float64
}

func (a *Account) Validate() bool {
	if a.AvailableCreditLimit <= 0 {
		a.addError("availableCreditLimit", "Limite de Crédito disponível não pode ser menor ou igual a 0")
	}

	if a.AvailableWithdrawalLimit <= 0 {
		a.addError("availableWithdrawalLimit", "Limite de Retirada disponível não pode ser menor ou igual a 0")
	}

	return a.IsValid()
}
