package services

import (
	"fmt"
	"github.com/lucaspizzo/api-payment/domains"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/infrastructure/repositories"
	"github.com/pkg/errors"
)

type AccountContract interface {
	Get(id uint64) (*domains.Account, error)
	List() (*[]*domains.Account, error)
	Update(availableCreditLimit float64, availableWithdrawalLimit float64, account *domains.Account) (*domains.Account, error)
	UpdateLimits(form *forms.LimitForm) (*domains.Account, error)
}

type AccountService struct {
	AccountRepository repositories.AccountQuerier `inject:""`
}

func (as *AccountService) Get(id uint64) (*domains.Account, error) {
	account := &domains.Account{}
	err := as.AccountRepository.GetById(id, account)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to get paymentType with id: %d", id))
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

func (as *AccountService) Update(availableCreditLimit float64, availableWithdrawalLimit float64, account *domains.Account) (*domains.Account, error) {

	account.AvailableCreditLimit = account.AvailableCreditLimit + availableCreditLimit
	account.AvailableWithdrawalLimit = account.AvailableWithdrawalLimit + availableWithdrawalLimit

	if !account.Validate() {
		return nil, account.GetError()
	}

	err := as.AccountRepository.Update(account)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
	}

	return account, err
}

func (as *AccountService) UpdateLimits(form *forms.LimitForm) (*domains.Account, error) {
	account, err := as.Get(form.AccountID)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
		return account, err
	}

	return as.Update(form.AvailableCreditLimit.Amount, form.AvailableWithdrawalLimit.Amount, account)
}
