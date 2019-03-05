package services

import (
	"fmt"
	"github.com/lucaspizzo/api-payment/domains"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/infrastructure/repositories"
	"github.com/pkg/errors"
	"time"
)

type TransactionContract interface {
	List() (*[]*domains.Transaction, error)
	ListByAccountId(accountID uint64) (*[]*domains.Transaction, error)
	Update(transaction *domains.Transaction) (*domains.Transaction, error)
	createPayment(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction
	processPayment(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction
	processPurchase(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction
	process(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) (*domains.Transaction, error)
	RegisterTransaction(form *forms.TransactionForm) (*domains.Transaction, error)
}

type TransactionService struct {
	AccountService        AccountContract                 `inject:""`
	OperationTypeService  OperationTypeContract           `inject:""`
	TransactionRepository repositories.TransactionQuerier `inject:""`
}

func (ts *TransactionService) Save(transaction *domains.Transaction) (*domains.Transaction, error) {
	err := ts.TransactionRepository.Insert(transaction)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to add transaction"))
	}
	return transaction, err
}

func (ts *TransactionService) List() (*[]*domains.Transaction, error) {
	transactions := &[]*domains.Transaction{}
	err := ts.TransactionRepository.FindAll(transactions)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
	}
	return transactions, err
}

func (ts *TransactionService) ListByAccountId(accountID uint64) (*[]*domains.Transaction, error) {
	transactions := &[]*domains.Transaction{}
	err := ts.TransactionRepository.FindAllByAccountId(transactions, accountID)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to list accounts"))
	}
	return transactions, err
}

func (ts *TransactionService) Update(transaction *domains.Transaction) (*domains.Transaction, error) {

	err := ts.TransactionRepository.Update(transaction)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to add transaction"))
	}
	return transaction, err
}

func (ts *TransactionService) createPayment(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction {
	transaction := domains.Transaction{}

	transaction.Account = *account
	transaction.AccountID = account.AccountID

	transaction.OperationType = *operationType
	transaction.OperationTypeID = operationType.OperationTypeID

	transaction.EventDate = time.Now()
	transaction.DueDate = time.Now()

	return &transaction
}

func (ts *TransactionService) processPurchase(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction {
	transaction := ts.createPayment(account, operationType, form)

	transaction.Amount = form.Amount * -1
	transaction.Balance = form.Amount * -1

	return transaction
}

func (ts *TransactionService) processPayment(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) *domains.Transaction {
	transaction := ts.createPayment(account, operationType, form)

	transaction.Amount = form.Amount

	var availableCredit = transaction.Amount
	transactions, _ := ts.ListByAccountId(account.AccountID)

	for _, t := range *transactions {
		if availableCredit > 0 && t.OperationType.Description != domains.OperationTypePayment && t.Balance != 0 {
			if t.Balance*(-1) > availableCredit {
				t.Balance += availableCredit
				availableCredit = 0
				t, _ = ts.Update(t)
				break
			} else if t.Balance*(-1) < availableCredit {
				availableCredit += t.Balance
				t.Balance = 0
				t, _ = ts.Update(t)
			} else if t.Balance*(-1) == availableCredit {
				t.Balance = 0
				availableCredit = 0
				t, _ = ts.Update(t)
				break
			}
		}
	}
	transaction.Balance = availableCredit

	return transaction

}

func (ts *TransactionService) process(account *domains.Account, operationType *domains.OperationType, form *forms.TransactionForm) (*domains.Transaction, error) {

	if account.AvailableCreditLimit <= form.Amount && account.AvailableWithdrawalLimit <= form.Amount {
		return nil, errors.New("Limit not available")
	}

	if operationType.Description == domains.OperationTypePayment {
		return ts.processPayment(account, operationType, form), nil
	} else {
		return ts.processPurchase(account, operationType, form), nil
	}

}

func (ts *TransactionService) RegisterTransaction(form *forms.TransactionForm) (*domains.Transaction, error) {

	// SEARCH ACCOUNT
	account, accountErr := ts.AccountService.Get(form.AccountID)
	if accountErr != nil {
		return nil, errors.New("Account not found")
	}

	// SEARCH OPERATION TYPE
	operationType, operationTypeErr := ts.OperationTypeService.Get(form.OperationTypeId)
	if operationTypeErr != nil {
		return nil, errors.New("Operation not found")
	}

	// PROCESS
	transaction, transactionErr := ts.process(account, operationType, form)
	if transactionErr != nil {
		return nil, transactionErr
	}

	// INSERT PAYMENT INTO DATABASE
	transaction, err := ts.Save(transaction)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to add transaction"))
		return nil, err
	}

	// UPDATE ACCOUNT LIMITS
	_, err = ts.AccountService.Update(transaction.Amount, transaction.Amount, &transaction.Account)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Unable to update account"))
		return nil, err
	}

	return transaction, nil
}
