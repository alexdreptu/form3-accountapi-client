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

// InvalidAccountTypeError is returned if Account Type is not 'accounts'.
type InvalidAccountTypeError struct {
	MustType string
	Type     string
}

func (e *InvalidAccountTypeError) Error() string {
	return fmt.Sprintf("must be '%s' but it's '%s'", e.MustType, e.Type)
}

// InvalidCountryError is returned if Country Code for a country is incorrect.
type InvalidCountryError struct {
	Country string
}

func (e *InvalidCountryError) Error() string {
	return fmt.Sprintf("invalid country '%s'", e.Country)
}

// InvalidCountryLengthError is returned if Country Code length is not 2 characters long.
type InvalidCountryLengthError struct {
	MustLength int
	Length     int
}

func (e *InvalidCountryLengthError) Error() string {
	return fmt.Sprintf("must be %d characters long but its length is %d",
		e.MustLength, e.Length)
}

// InvalidBankIDLengthError is returned if Bank ID length for a country is incorrect.
type InvalidBankIDLengthError struct {
	MustLength int
	Length     int
}

func (e *InvalidBankIDLengthError) Error() string {
	return fmt.Sprintf("must be %d characters long but its length is %d",
		e.MustLength, e.Length)
}

// InvalidBankIDCodeError is returned if the Bank ID Code for a country is incorrect.
type InvalidBankIDCodeError struct {
	MustCode string
	Code     string
}

func (e *InvalidBankIDCodeError) Error() string {
	return fmt.Sprintf("must be '%s' but it's '%s'", e.MustCode, e.Code)
}

// InvalidBICLengthError is returned if BIC length is incorrect.
type InvalidBICLengthError struct {
	MustLength1 int
	MustLength2 int
	Length      int
}

func (e *InvalidBICLengthError) Error() string {
	return fmt.Sprintf("must be either %d or %d characters long but its length is %d",
		e.MustLength1, e.MustLength2, e.Length)
}

// InvalidAccountNumberLengthError is returned if Account Number length for a country is incorrect.
type InvalidAccountNumberLengthError struct {
	MustLength     int
	MustLengthFrom int
	MustLengthTo   int
	Length         int
}

func (e *InvalidAccountNumberLengthError) Error() string {
	var message string
	if e.MustLength == 0 || e.MustLengthFrom != 0 && e.MustLengthTo != 0 {
		message = fmt.Sprintf("must be between %d and %d characters long but its length is %d",
			e.MustLengthFrom, e.MustLengthTo, e.Length)
	} else {
		message = fmt.Sprintf("must be %d characters long but its length is %d",
			e.MustLength, e.Length)
	}
	return message
}

// InvalidAccountNumberError is returned if Account Number is not a number.
type InvalidAccountNumberError struct {
	Number string
}

func (e *InvalidAccountNumberError) Error() string {
	return fmt.Sprintf("must be a number but '%s' is not", e.Number)
}

// InvalidBaseCurrencyLengthError is returned if Base Currency length is not 3 characters long.
// See https://www.iso.org/iso-4217-currency-codes.html
type InvalidBaseCurrencyLengthError struct {
	MustLength int
	Length     int
}

func (e *InvalidBaseCurrencyLengthError) Error() string {
	return fmt.Sprintf("must be %d characters long but its length is %d",
		e.MustLength, e.Length)
}

// InvalidBaseCurrencyError is returned if Base Currency for a country is incorrect.
type InvalidBaseCurrencyError struct {
	MustCurrency string
	Currency     string
}

func (e *InvalidBaseCurrencyError) Error() string {
	return fmt.Sprintf("must be '%s' but it's '%s'", e.MustCurrency, e.Currency)
}

// InvalidFirstNameLengthError is returned if Firstname is not between between 2 and 140 characters long.
type InvalidFirstNameLengthError struct {
	MustLengthFrom int
	MustLengthTo   int
	Length         int
}

func (e *InvalidFirstNameLengthError) Error() string {
	return fmt.Sprintf("must be between %d and %d characters long but its length is %d",
		e.MustLengthFrom, e.MustLengthTo, e.Length)
}

// InvalidCustomerIDLengthError is returned if CustomerID Length is not between 5 and 15 characters long.
type InvalidCustomerIDLengthError struct {
	MustLengthFrom int
	MustLengthTo   int
	Length         int
}

func (e *InvalidCustomerIDLengthError) Error() string {
	return fmt.Sprintf("must be between %d and %d characters long but its length is %d",
		e.MustLengthFrom, e.MustLengthTo, e.Length)
}

// InvalidAlternativeBankAccountArrayLengthError is returned if Alternative Bank Account Array's length is bigger than 3.
type InvalidAlternativeBankAccountArrayLengthError struct {
	MustLengthFrom int
	MustLengthTo   int
	Length         int
}

func (e *InvalidAlternativeBankAccountArrayLengthError) Error() string {
	return fmt.Sprintf("must be between %d and %d in length but its length is %d",
		e.MustLengthFrom, e.MustLengthTo, e.Length)
}

// InvalidAlternativeBankAccountElemLengthError is returned if Alternative Bank Account element is not between 3 and 140 characters long.
type InvalidAlternativeBankAccountElemLengthError struct {
	MustLengthFrom int
	MustLengthTo   int
	Length         int
}

func (e *InvalidAlternativeBankAccountElemLengthError) Error() string {
	return fmt.Sprintf("must between %d and %d characters long but its length is %d",
		e.MustLengthFrom, e.MustLengthTo, e.Length)
}
