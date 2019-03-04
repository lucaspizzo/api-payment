package domains

type Account struct {
	AccountID uint64 `sql:"primary_key;"`
	AvailableCreditLimit float64
	AvailableWithdrawalLimit float64
}
