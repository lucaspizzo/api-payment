package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaspizzo/api-payment/domains"
)

type AccountQuerier interface {
	GetById(id uint64, model *domains.Account) error
	Insert(model *domains.Account) error
	FindAll(list *[]*domains.Account) error
	Update(model *domains.Account) error
}

type AccountRepository struct {
	DB *gorm.DB
}

func (a *AccountRepository) FindAll(list *[]*domains.Account) error {
	return a.DB.Find(&list).Error
}

func (a *AccountRepository) GetById(id uint64, model *domains.Account) error {
	return a.DB.Where("account_id = ?", id).Find(&model).Error
}

func (a *AccountRepository) Insert(model *domains.Account) error {
	return a.DB.Create(model).Error
}

func (a *AccountRepository) Update(model *domains.Account) error {
	err := a.DB.Save(model).Error
	return err
}
