package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrCreditCardTransactionRepositoryHouston = errors.New("Houston, we have unknown error into credit card transaction repository")
	ErrTransactionStatusRepositoryHouston     = errors.New("Houston, we have unknown error into transaction status repository")
	ErrAuthorizationServiceHouston            = errors.New("Houston, we have unknown error into authorization service")
)

type ErrCreditCardTransactionNotFound struct {
	criteria string
	value    string
}

func NewErrCreditCardTransactionNotFound(criteria string, value string) ErrCreditCardTransactionNotFound {
	return ErrCreditCardTransactionNotFound{
		criteria: criteria,
		value:    value,
	}
}

func (e ErrCreditCardTransactionNotFound) Error() string {
	return fmt.Sprintf("Credit card transaction not found by criteria [ %s ] and value [ %s ]", e.criteria, e.value)
}

type ErrInvalidCreditCardNumber struct {
	number string
}

func NewErrInvalidCreditCardNumber(number string) ErrInvalidCreditCardNumber {
	return ErrInvalidCreditCardNumber{
		number: number,
	}
}

func (e ErrInvalidCreditCardNumber) Error() string {
	return fmt.Sprintf("Invalid credit card number [ %s ]", e.number)
}

type ErrCreditCardTransctionInstallments struct {
	installments uint
}

func NewErrCreditCardTransctionInstallments(installments uint) ErrCreditCardTransctionInstallments {
	return ErrCreditCardTransctionInstallments{
		installments: installments,
	}
}

func (e ErrCreditCardTransctionInstallments) Error() string {
	return fmt.Sprintf(
		"Invalid installments, expected between [ %d ] and [ %d ], got [ %d ]",
		INSTALLMENTS_MIN,
		INSTALLMENTS_MAX,
		e.installments,
	)
}

type ErrTransctionStatusInvalid struct {
	status string
}

func NewErrTransctionStatusInvalid(status string) ErrTransctionStatusInvalid {
	return ErrTransctionStatusInvalid{
		status: status,
	}
}

func (e ErrTransctionStatusInvalid) Error() string {
	return fmt.Sprintf(
		"Invalid transction status, expected one of [ %s ], got [ %s ]",
		strings.Join(GetTransactionStatusAvailable(), ", "),
		e.status,
	)
}

type ErrAuthorizationResponseStatusInvalid struct {
	status string
}

func NewErrAuthorizationResponseStatusInvalid(status string) ErrAuthorizationResponseStatusInvalid {
	return ErrAuthorizationResponseStatusInvalid{
		status: status,
	}
}

func (e ErrAuthorizationResponseStatusInvalid) Error() string {
	return fmt.Sprintf(
		"Invalid authorization response status, expected one of [ %s ], got [ %s ]",
		strings.Join(GetAuthorizationStatusAvailable(), ", "),
		e.status,
	)
}
