package domain_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCardNumber(t *testing.T) {
	tests := []struct {
		name           string
		cardNumber     string
		returnExpected bool
	}{
		{
			name:           "Test is valid card number with correct number",
			cardNumber:     "5130731304267489",
			returnExpected: true,
		},
		{
			name:           "Test is valid card number with invalid number",
			cardNumber:     "513073130426",
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidCardNumber(test.cardNumber)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}

func TestSanitizeCreditCardNumber(t *testing.T) {
	tests := []struct {
		name               string
		cardNumber         string
		cardNumberExpected string
		errExpected        error
	}{
		{
			name:               "Test sanitize credit card number with correct number",
			cardNumber:         "5130731304267489",
			cardNumberExpected: "513073XXXXXX7489",
			errExpected:        nil,
		},
		{
			name:               "Test sanitize credit card number with invalid number",
			cardNumber:         "513073130426",
			cardNumberExpected: "",
			errExpected:        domain.NewErrInvalidCreditCardNumber("513073130426"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cardNumberGot, errGot := domain.SanitizeCreditCardNumber(test.cardNumber)
			assert.Equal(t, test.cardNumberExpected, cardNumberGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
