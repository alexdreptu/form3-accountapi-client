package accountapi

import (
	"fmt"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	accountType        = "accounts"
	countryLength      = 2
	baseCurrencyLength = 3
)

// FirstName attribute length
const (
	firstNameLengthStart = 2
	firstNameLengthStop  = 140
)

// AlternativeAccountNames attribute lengths
const (
	alternativeAccountNamesArrayLengthStart = 1
	alternativeAccountNamesArrayLengthStop  = 3
	alternativeAccountNamesElemLengthStart  = 3
	alternativeAccountNamesElemLengthStop   = 140
)

// BIC length range
const (
	BICLength8  = 8
	BICLength11 = 11
)

// general validation rules
var (
	validateBICMatch = validation.Match(
		regexp.MustCompile("^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$"),
	)

	validateBICLength = validation.By(
		func(value interface{}) error {
			bic, _ := value.(string)
			length := len(bic)
			if bic != "" && length != BICLength8 && length != BICLength11 {
				return &InvalidBICLengthError{length}
			}
			return nil
		},
	)

	validateBaseCurrencyLength = validation.By(
		func(value interface{}) error {
			currency, _ := value.(string)
			length := len(currency)
			if currency != "" && length != baseCurrencyLength {
				return &InvalidBaseCurrencyLengthError{length}
			}
			return nil
		},
	)

	validateStringNumber = validation.By(
		func(value interface{}) error {
			number, _ := value.(string)
			if number != "" {
				_, err := strconv.ParseFloat(number, 64)
				if err != nil {
					return &InvalidAccountNumberError{number}
				}
			}
			return nil
		},
	)
)

func (a *Attributes) validateUnitedKingdom() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthUnitedKingdom {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeUnitedKingdom {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthUnitedKingdom {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthUnitedKingdom),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyUnitedKingdom {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateAustralia() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if id != "" && length != BankIDLengthAustralia {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeAustralia {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" &&
					!(length >= AccountNumberLengthAustraliaStart &&
						length <= AccountNumberLengthAustraliaStop) {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be between %d and %d characters long",
							AccountNumberLengthAustraliaStart, AccountNumberLengthAustraliaStop),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccoountNumberFirstCharacter = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				if number != "" && number[0] == '0' {
					return ErrAccountNumberFirstCharZero
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateAccoountNumberFirstCharacter,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyAustralia {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateBelgium() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthBelgium {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeBelgium {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthBelgium {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthBelgium),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyBelgium {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateCanada() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if id != "" && length != BankIDLengthCanada {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankIDFirstCharacter = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				if id != "" && id[0] != '0' {
					return ErrBankIDCodeFirstCharNonZero
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validateBankIDLength,
			validateBankIDFirstCharacter,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != "" && code != BankIDCodeCanada {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" &&
					!(length >= AccountNumberLengthCanadaStart &&
						length <= AccountNumberLengthCanadaStop) {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be between %d and %d characters long",
							AccountNumberLengthCanadaStart, AccountNumberLengthCanadaStop),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyCanada {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateFrance() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthFrance {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeFrance {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthFrance {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthFrance),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyFrance {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateGermany() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthGermany {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeGermany {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthGermany {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthGermany),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyGermany {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateGreece() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthGreece {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeGreece {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthGreece {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthGreece),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyGreecee {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateHongKong() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if id != "" && length != BankIDLengthHongKong {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != "" && code != BankIDCodeHongKong {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" &&
					!(length >= AccountNumberLengthHongKongStart &&
						length <= AccountNumberLengthHongKongStop) {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be between %d and %d characters long",
							AccountNumberLengthHongKongStart, AccountNumberLengthHongKongStop),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyHongKong {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateItaly() error {
	const (
		accountNumberNotPresent = BankIDLengthItalyAccountNumberNotPresent
		accountNumberPresent    = BankIDLengthItalyAccountNumberPresent
	)

	var (
		validateBankIDLength = func(mustLength int) validation.Rule {
			return validation.By(
				func(value interface{}) error {
					id, _ := value.(string)
					length := len(id)
					if length != mustLength {
						return &InvalidBankIDLengthError{length}
					}
					return nil
				},
			)
		}

		validateBankID = []validation.Rule{
			validation.Required,
			validation.When(
				validation.IsEmpty(a.AccountNumber),
				validateBankIDLength(accountNumberNotPresent),
				validation.Skip,
			),
			validation.When(
				!validation.IsEmpty(a.AccountNumber),
				validateBankIDLength(accountNumberPresent),
			),
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeItaly {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthItaly {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthItaly),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyItaly {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateLuxembourg() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthLuxembourg {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeLuxembourg {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthLuxembourg {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthLuxembourg),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyLuxembourg {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateNetherlands() error {
	var (
		validateBankIDMatch = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				if id != "" {
					return ErrBankIDNotBlank
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validateBankIDMatch,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICMatch,
			validateStringNumber,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != "" {
					return ErrBankIDCodeNotBlank
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthNetherlands {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthNetherlands),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyNetherlands {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validatePoland() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthPoland {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodePoland {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthPoland {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthPoland),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyPoland {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validatePortugal() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthPortugal {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodePortugal {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthPortugal {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthPortugal),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyPortugal {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateSpain() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthSpain {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeSpain {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthSpain {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthSpain),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencySpain {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateSwitzerland() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthSwitzerland {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeSwitzerland {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthSwitzerland {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be %d characters long",
							AccountNumberLengthSwitzerland),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencySwitzerland {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (a *Attributes) validateUnitedStates() error {
	var (
		validateBankIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if length != BankIDLengthUnitedStates {
					return &InvalidBankIDLengthError{length}
				}
				return nil
			},
		)

		validateBankID = []validation.Rule{
			validation.Required,
			validateBankIDLength,
			validateStringNumber,
		}

		validateBIC = []validation.Rule{
			validation.Required,
			validateBICLength,
			validateBICMatch,
		}

		validateBankIDCodeMatch = validation.By(
			func(value interface{}) error {
				code, _ := value.(string)
				if code != BankIDCodeUnitedStates {
					return &InvalidBankIDCodeError{code}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" &&
					!(length >= AccountNumberLengthUnitedStatesStart &&
						length <= AccountNumberLengthUnitedStatesStop) {
					return &InvalidAccountNumberLengthError{
						Message: fmt.Sprintf("must be between %d and %d characters long",
							AccountNumberLengthUnitedStatesStart, AccountNumberLengthUnitedStatesStop),
						Length: length,
					}
				}
				return nil
			},
		)

		validateAccountNumber = []validation.Rule{
			validateAccountNumberLength,
			validateStringNumber,
		}

		validateBaseCurrencyMatch = validation.By(
			func(value interface{}) error {
				currency, _ := value.(string)
				if currency != "" && currency != CurrencyUnitedStates {
					return &InvalidBaseCurrencyError{currency}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			validateBaseCurrencyMatch,
		}
	)

	return validation.ValidateStruct(a,
		validation.Field(&a.BankID, validateBankID...),
		validation.Field(&a.BIC, validateBIC...),
		validation.Field(&a.BankIDCode, validateBankIDCode...),
		validation.Field(&a.AccountNumber, validateAccountNumber...),
		validation.Field(&a.BaseCurrency, validateBaseCurrency...),
	)
}

func (o *Options) validate() error {
	var (
		validateTypeMatch = validation.By(
			func(value interface{}) error {
				match, _ := value.(string)
				if match != accountType {
					return &InvalidAccountTypeError{match}
				}
				return nil
			},
		)

		validateType = []validation.Rule{
			validation.Required,
			validateTypeMatch,
		}

		validateID = []validation.Rule{
			validation.Required,
			is.UUID,
		}

		validateOrganisationID = []validation.Rule{
			validation.Required,
			is.UUID,
		}
	)

	return validation.ValidateStruct(o,
		validation.Field(&o.Type, validateType...),
		validation.Field(&o.ID, validateID...),
		validation.Field(&o.OrganisationID, validateOrganisationID...),
	)
}

func (a *Attributes) validate() error {
	var (
		validateCountryLength = validation.By(
			func(value interface{}) error {
				country, _ := value.(string)
				length := len(country)
				if length != countryLength {
					return &InvalidCountryLengthError{length}
				}
				return nil
			},
		)

		validateCountry = []validation.Rule{
			validation.Required,
			validateCountryLength,
		}

		validateAlternativeAccountNames = []validation.Rule{
			validation.Length(
				alternativeAccountNamesArrayLengthStart,
				alternativeAccountNamesArrayLengthStop,
			),
			validation.Each(validation.Length(
				alternativeAccountNamesElemLengthStart,
				alternativeAccountNamesElemLengthStop,
			)),
		}

		validateFirstNameLength = validation.By(
			func(value interface{}) error {
				name, _ := value.(string)
				length := len(name)
				if name != "" &&
					!(length >= firstNameLengthStart &&
						length <= firstNameLengthStop) {
					return &InvalidFirstNameLengthError{length}
				}
				return nil
			},
		)

		validateFirstName = []validation.Rule{
			validateFirstNameLength,
		}
	)

	if err := validation.ValidateStruct(a,
		validation.Field(&a.Country, validateCountry...),
		validation.Field(&a.AlternativeBankAccountNames, validateAlternativeAccountNames...),
		validation.Field(&a.FirstName, validateFirstName...),
	); err != nil {
		return err
	}

	switch a.Country {
	case CountryUnitedKingdom:
		return a.validateUnitedKingdom()

	case CountryAustralia:
		return a.validateAustralia()

	case CountryBelgium:
		return a.validateBelgium()

	case CountryCanada:
		return a.validateCanada()

	case CountryFrance:
		return a.validateFrance()

	case CountryGermany:
		return a.validateGermany()

	case CountryGreece:
		return a.validateGreece()

	case CountryHongKong:
		return a.validateHongKong()

	case CountryItaly:
		return a.validateItaly()

	case CountryLuxembourg:
		return a.validateLuxembourg()

	case CountryNetherlands:
		return a.validateNetherlands()

	case CountryPoland:
		return a.validatePoland()

	case CountryPortugal:
		return a.validatePortugal()

	case CountrySpain:
		return a.validateSpain()

	case CountrySwitzerland:
		return a.validateSwitzerland()

	case CountryUnitedStates:
		return a.validateUnitedStates()

	default:
		return &InvalidCountryError{a.Country}
	}
}
