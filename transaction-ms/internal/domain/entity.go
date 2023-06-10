package domain

import (
	"time"

	"github.com/Rhymond/go-money"
)

const (
	INSTALLMENTS_MIN                       = 1
	INSTALLMENTS_MAX                       = 12
	TRANSACTION_STATUS_CREATED             = "created"
	TRANSACTION_STATUS_AUTHORIZED          = "authorized"
	TRANSACTION_STATUS_REFUSED             = "refused"
	TRANSACTION_STATUS_FINISHED            = "finished"
	TRANSACTION_STATUS_ERROR_CREATED       = "error_created"
	TRANSACTION_STATUS_ERROR_AUTHORIZATION = "error_authorization"
	AUTHORIZATION_STATUS_AUTHORIZED        = "authorized"
	AUTHORIZATION_STATUS_DECLINED          = "declined"
)

type Amount struct {
	Value uint
}

func (a Amount) BRL() string {
	return money.New(int64(a.Value), money.BRL).Display()
}

type CreditCardTransction struct {
	ID               *uint
	TransactionID    string
	CardNumber       string
	Amount           Amount
	Installments     uint
	Description      string
	TransctionStatus []TransctionStatus
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}

func NewCreditCardTransction(
	transactionID, cardNumber string,
	amount uint,
	installments uint,
	description string,
	transctionStatus []TransctionStatus,
) (CreditCardTransction, error) {
	if !IsValidInstallments(installments) {
		return CreditCardTransction{}, NewErrCreditCardTransctionInstallments(installments)
	}
	cardNumberSanitized, err := SanitizeCreditCardNumber(cardNumber)
	if err != nil {
		return CreditCardTransction{}, err
	}

	return CreditCardTransction{
		TransactionID: transactionID,
		CardNumber:    cardNumberSanitized,
		Amount: Amount{
			Value: amount,
		},
		Installments:     installments,
		Description:      description,
		TransctionStatus: transctionStatus,
	}, nil
}

type TransctionStatus struct {
	ID                     *uint
	CreditCardTransctionID uint
	Status                 string
	CreatedAt              *time.Time
}

func NewTransctionStatus(creditCardTransctionID uint, status string) (TransctionStatus, error) {
	if !IsValidTransactionStatus(status) {
		return TransctionStatus{}, NewErrTransctionStatusInvalid(status)
	}

	return TransctionStatus{
		CreditCardTransctionID: creditCardTransctionID,
		Status:                 status,
	}, nil
}

type AuthorizationRequest struct {
	HolderName   string
	CardNumber   string
	CVV          string
	ExpireDate   time.Time
	Amount       uint
	Installments uint
}

func NewAuthorizationRequest(
	holderName, cardNumer, cvv string,
	expireDate time.Time,
	amount, installments uint,
) (AuthorizationRequest, error) {
	if !IsValidCardNumber(cardNumer) {
		return AuthorizationRequest{}, NewErrInvalidCreditCardNumber(cardNumer)
	}

	return AuthorizationRequest{
		HolderName:   holderName,
		CardNumber:   cardNumer,
		CVV:          cvv,
		ExpireDate:   expireDate,
		Amount:       amount,
		Installments: installments,
	}, nil
}

type AuthorizationResponse struct {
	Status string
}

func NewAuthorizationResponse(status string) (AuthorizationResponse, error) {
	if !IsValidAuthorizationStatus(status) {
		return AuthorizationResponse{}, NewErrAuthorizationResponseStatusInvalid(status)
	}

	return AuthorizationResponse{
		Status: status,
	}, nil
}

func IsValidInstallments(i uint) bool {
	return i >= INSTALLMENTS_MIN && i <= INSTALLMENTS_MAX
}

func GetTransactionStatusAvailable() []string {
	return []string{
		TRANSACTION_STATUS_CREATED,
		TRANSACTION_STATUS_AUTHORIZED,
		TRANSACTION_STATUS_REFUSED,
		TRANSACTION_STATUS_FINISHED,
		TRANSACTION_STATUS_ERROR_CREATED,
		TRANSACTION_STATUS_ERROR_AUTHORIZATION,
	}
}

func IsValidTransactionStatus(v string) bool {
	return contains(v, GetTransactionStatusAvailable())
}

func GetAuthorizationStatusAvailable() []string {
	return []string{
		AUTHORIZATION_STATUS_AUTHORIZED,
		AUTHORIZATION_STATUS_DECLINED,
	}
}

func IsValidAuthorizationStatus(v string) bool {
	return contains(v, GetAuthorizationStatusAvailable())
}

func contains(v string, e []string) bool {
	for _, s := range e {
		if v == s {
			return true
		}
	}
	return false
}
