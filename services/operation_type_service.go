package services

import (
	"fmt"
	"github.com/lucaspizzo/api-payment/domains"
	"github.com/lucaspizzo/api-payment/infrastructure/repositories"
	"github.com/pkg/errors"
)

type OperationTypeContract interface {
	Get(id uint64) (*domains.OperationType, error)
}

type OperationTypeService struct {
	OperationTypeRepository repositories.OperationTypeQuerier `inject:""`
}

func (os *OperationTypeService) Get(id uint64) (*domains.OperationType, error) {
	operationType := &domains.OperationType{}
	err := os.OperationTypeRepository.GetById(id, operationType)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to get operationType with id: %d", id))
	}
	return operationType, err
}
