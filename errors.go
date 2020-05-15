package accountapi

import (
	"errors"
	"fmt"
)

// error strings
const (
	errMsgBlank            = "cannot be blank"
	errMsgNotBlank         = "must be blank"
	errMsgFirstCharZero    = "first character cannot be '0'"
	errMsgFirstCharNotZero = "first character must be '0'"
)

// custom errors
var (
	ErrAccountTypeBlank           = errors.New(errMsgBlank)
	ErrBankIDBlank                = errors.New(errMsgBlank)
	ErrBankIDNotBlank             = errors.New(errMsgNotBlank)
	ErrBankIDCodeFirstCharNonZero = errors.New(errMsgFirstCharNotZero)
	ErrBankIDCodeNotBlank         = errors.New(errMsgNotBlank)
	ErrBankIDCodeBlank            = errors.New(errMsgBlank)
	ErrAccountNumberBlank         = errors.New(errMsgBlank)
	ErrAccountNumberFirstCharZero = errors.New(errMsgFirstCharZero)
)

// InvalidAccountTypeError is returned when Account Type is not 'accounts'.
type InvalidAccountTypeError struct {
	Type string
}

func (e *InvalidAccountTypeError) Error() string {
	return fmt.Sprintf("must be 'accounts' but it's '%s'", e.Type)
}

// InvalidCountryError is returned when Country Code for a country is incorrect.
type InvalidCountryError struct {
	Country string
}

func (e *InvalidCountryError) Error() string {
	return fmt.Sprintf("invalid country '%s'", e.Country)
}

// InvalidCountryLengthError is returned when Country Code length is not 2 characters long.
type InvalidCountryLengthError struct {
	Length int
}

func (e *InvalidCountryLengthError) Error() string {
	return fmt.Sprintf("must be 2 characters long but it's %d", e.Length)
}

// InvalidBankIDLengthError is returned when Bank ID length for a country is incorrect.
type InvalidBankIDLengthError struct {
	Length int
}

func (e *InvalidBankIDLengthError) Error() string {
	return fmt.Sprintf("must be 6 characters long but it's %d", e.Length)
}

// InvalidBankIDCodeError is returned when the Bank ID Code for a country is incorrect.
type InvalidBankIDCodeError struct {
	Code string
}

func (e *InvalidBankIDCodeError) Error() string {
	return fmt.Sprintf("invalid bank id code '%s'", e.Code)
}

// InvalidBICLengthError is returned when BIC length is incorrect.
type InvalidBICLengthError struct {
	Length int
}

func (e *InvalidBICLengthError) Error() string {
	return fmt.Sprintf("must be either 8 or 11 characters long but it's %d", e.Length)
}

// InvalidAccountNumberLengthError is returned when Account Number length for a country is incorrect.
type InvalidAccountNumberLengthError struct {
	Message string
	Length  int
}

func (e *InvalidAccountNumberLengthError) Error() string {
	return fmt.Sprintf(e.Message+" "+"but it's %d", e.Length)
}

// InvalidAccountNumberError is returned when Account Number is not a number
type InvalidAccountNumberError struct {
	Number string
}

func (e *InvalidAccountNumberError) Error() string {
	return fmt.Sprintf("must be a number but '%s' is not", e.Number)
}

// InvalidBaseCurrencyLengthError is returned when Base Currency length is not 3 characters long.
// See https://www.iso.org/iso-4217-currency-codes.html
type InvalidBaseCurrencyLengthError struct {
	Length int
}

func (e *InvalidBaseCurrencyLengthError) Error() string {
	return fmt.Sprintf("must be 3 characters long but it's %d", e.Length)
}

// InvalidBaseCurrencyError is returned when Base Currency for a country is incorrect.
type InvalidBaseCurrencyError struct {
	Currency string
}

func (e *InvalidBaseCurrencyError) Error() string {
	return fmt.Sprintf("invalid base currency '%s'", e.Currency)
}

// InvalidFirstNameLengthError is returned if Firstname is not between between 2 and 140 characters long.
type InvalidFirstNameLengthError struct {
	Length int
}

func (e *InvalidFirstNameLengthError) Error() string {
	return fmt.Sprintf("must be between 2 and 140 characters long but it's %d", e.Length)
}
