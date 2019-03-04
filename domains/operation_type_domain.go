package domains

type OperationType struct {
	OperationTypeID uint64
	Description     OperationTypeDescription `sql:"primary_key;"`
	ChargeOrder     uint64
}
