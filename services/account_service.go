package services

import (
	"fmt"
	"github.com/lucaspizzo/api-payment/domains"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/infrastructure/repositories"

	"github.com/pkg/errors"
)

type AccountContract interface {
	Register(name string) (*domains.Account, error)
	Get(id uint64) (*domains.Account, error)
	List() (*[]*domains.Account, error)
	UpdateLimits(form *forms.LimitForm) (*domains.Account, error)
}

type AccountService struct {
	AccountRepository repositories.AccountQuerier `inject:""`
}


func (as *AccountService) Register(name string) (*domains.Account, error) {

	account := &domains.Account{}

	if err := as.AccountRepository.Insert(account); err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to save new account"))
		return nil, err
	}
	return account, nil
}

func (as *AccountService) Get(id uint64) (*domains.Account, error) {
	account := &domains.Account{}
	err := as.AccountRepository.GetById(id, account)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to get paymentType with id: %s", id))
	}
	return account, err
}

func (as *AccountService) List() (*[]*domains.Account, error) {
	accounts := &[]*domains.Account{}
	err := as.AccountRepository.FindAll(accounts)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
	}
	return accounts, err
}

func (as *AccountService) UpdateLimits(form *forms.LimitForm) (*domains.Account, error) {
	account, err := as.Get(form.AccountID)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
		return account, err
	}

	account.AvailableCreditLimit = form.AvailableCreditLimit.Amount
	account.AvailableWithdrawalLimit = form.AvailableWithdrawalLimit.Amount


	err = as.AccountRepository.Update(account)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
	}
	return account, err
}

