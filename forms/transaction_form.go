package forms

type TransactionForm struct {
	BaseForm
	AccountID       uint64  `json:"account_id"`
	OperationTypeId uint64  `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (t *TransactionForm) Validate() bool {
	if t.AccountID == 0 {
		t.addError("accountID", "invalid accountId")
	}

	if t.Amount <= 0 {
		t.addError("amount", "invalid amount")
	}

	if t.OperationTypeId == 0 {
		t.addError("operationTypeId", "invalid operationTypeId")
	}

	return t.IsValid()
}
