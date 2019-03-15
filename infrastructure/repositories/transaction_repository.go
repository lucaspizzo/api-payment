package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaspizzo/api-payment/domains"
)

type TransactionQuerier interface {
	GetById(id uint64, model *domains.Transaction) error
	Insert(model *domains.Transaction) error
	FindAll(list *[]*domains.Transaction) error
	FindAllByAccountId(list *[]*domains.Transaction, accountID uint64) error
	Update(model *domains.Transaction) error
}

type TransactionRepository struct {
	DB *gorm.DB
}

func (t *TransactionRepository) FindAll(list *[]*domains.Transaction) error {
	return t.DB.Find(&list).Error
}

func (t *TransactionRepository) FindAllByAccountId(list *[]*domains.Transaction, accountID uint64) error {
	return t.DB.Where("account_id = ? ", accountID).Preload("OperationType").Find(&list).Error
}

func (t *TransactionRepository) GetById(id uint64, model *domains.Transaction) error {
	return t.DB.Where("transaction_id = ?", id).Find(&model).Error
}

func (t *TransactionRepository) Insert(model *domains.Transaction) error {
	return t.DB.Create(model).Error
}

func (t *TransactionRepository) Update(model *domains.Transaction) error {
	err := t.DB.Save(model).Error
	return err
}
