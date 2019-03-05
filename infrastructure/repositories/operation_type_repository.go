package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaspizzo/api-payment/domains"
)

type OperationTypeQuerier interface {
	GetById(id uint64, model *domains.OperationType) error
}

type OperationTypeRepository struct {
	DB *gorm.DB
}

func (o *OperationTypeRepository) GetById(id uint64, model *domains.OperationType) error {
	return o.DB.Where("operation_type_id = ?", id).Find(&model).Error
}
