package accountapi

import (
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	accountType           = "accounts"
	countryLength         = 2
	baseCurrencyLength    = 3
	customerIDLengthStart = 5
	customerIDLengthStop  = 15
)

// FirstName attribute length
const (
	firstNameLengthStart = 2
	firstNameLengthStop  = 140
)

// AlternativeBankAccountNames attribute lengths
const (
	alternativeBankAccountNamesArrayLengthStart = 1
	alternativeBankAccountNamesArrayLengthStop  = 3
	alternativeBankAccountNamesElemLengthStart  = 3
	alternativeBankAccountNamesElemLengthStop   = 140
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
				return &InvalidBICLengthError{
					MustLength1: BICLength8,
					MustLength2: BICLength11,
					Length:      length,
				}
			}
			return nil
		},
	)

	validateBaseCurrencyLength = validation.By(
		func(value interface{}) error {
			currency, _ := value.(string)
			length := len(currency)
			if currency != "" && length != baseCurrencyLength {
				return &InvalidBaseCurrencyLengthError{
					MustLength: baseCurrencyLength,
					Length:     length,
				}
			}
			return nil
		},
	)

	validateStringNumber = validation.By(
		func(value interface{}) error {
			number, _ := value.(string)
			if number != "" {
				_, err := strconv.ParseInt(number, 10, 64)
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthUnitedKingdom,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeUnitedKingdom,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthUnitedKingdom {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthUnitedKingdom,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyUnitedKingdom,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthAustralia,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeAustralia,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
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
						MustLengthFrom: AccountNumberLengthAustraliaStart,
						MustLengthTo:   AccountNumberLengthAustraliaStop,
						Length:         length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyAustralia,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthBelgium,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeBelgium,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthBelgium {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthBelgium,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyBelgium,
						Currency:     currency,
					}
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthCanada,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeCanada,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			is.Alpha,
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
						MustLengthFrom: AccountNumberLengthCanadaStart,
						MustLengthTo:   AccountNumberLengthCanadaStop,
						Length:         length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyCanada,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthFrance,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeFrance,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthFrance {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthFrance,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyFrance,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthGermany,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeGermany,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthGermany {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthGermany,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyGermany,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthGreece,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeGreece,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthGreece {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthGreece,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyGreecee,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthHongKong,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeHongKong,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			is.Alpha,
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
						MustLengthFrom: AccountNumberLengthHongKongStart,
						MustLengthTo:   AccountNumberLengthHongKongStop,
						Length:         length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyHongKong,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
						return &InvalidBankIDLengthError{
							MustLength: mustLength,
							Length:     length,
						}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeItaly,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthItaly {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthItaly,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyItaly,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthLuxembourg,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeLuxembourg,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthLuxembourg {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthLuxembourg,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyLuxembourg,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
			validateBICLength,
			validateBICMatch,
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
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthNetherlands {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthNetherlands,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyNetherlands,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthPoland,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodePoland,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthPoland {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthPoland,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyPoland,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthPortugal,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodePortugal,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthPortugal {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthPortugal,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyPortugal,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthSpain,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeSpain,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthSpain {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthSpain,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencySpain,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthSwitzerland,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeSwitzerland,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
			validateBankIDCodeMatch,
		}

		validateAccountNumberLength = validation.By(
			func(value interface{}) error {
				number, _ := value.(string)
				length := len(number)
				if number != "" && length != AccountNumberLengthSwitzerland {
					return &InvalidAccountNumberLengthError{
						MustLength: AccountNumberLengthSwitzerland,
						Length:     length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencySwitzerland,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidBankIDLengthError{
						MustLength: BankIDLengthUnitedStates,
						Length:     length,
					}
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
					return &InvalidBankIDCodeError{
						MustCode: BankIDCodeUnitedStates,
						Code:     code,
					}
				}
				return nil
			},
		)

		validateBankIDCode = []validation.Rule{
			validation.Required,
			is.Alpha,
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
						MustLengthFrom: AccountNumberLengthUnitedStatesStart,
						MustLengthTo:   AccountNumberLengthUnitedStatesStop,
						Length:         length,
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
					return &InvalidBaseCurrencyError{
						MustCurrency: CurrencyUnitedKingdom,
						Currency:     currency,
					}
				}
				return nil
			},
		)

		validateBaseCurrency = []validation.Rule{
			validateBaseCurrencyLength,
			is.Alpha,
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
					return &InvalidAccountTypeError{
						MustType: accountType,
						Type:     match,
					}
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
	const countryLength = 2

	var (
		validateCountryLength = validation.By(
			func(value interface{}) error {
				country, _ := value.(string)
				length := len(country)
				if length != countryLength {
					return &InvalidCountryLengthError{
						MustLength: countryLength,
						Length:     length,
					}
				}
				return nil
			},
		)

		validateCountry = []validation.Rule{
			validation.Required,
			validateCountryLength,
		}

		validateAlternativeBankAccountNamesArrayLength = validation.By(
			func(value interface{}) error {
				array, _ := value.([]string)
				length := len(array)
				if length != 0 &&
					!(length >= alternativeBankAccountNamesArrayLengthStart &&
						length <= alternativeBankAccountNamesArrayLengthStop) {
					return &InvalidAlternativeBankAccountArrayLengthError{
						MustLengthFrom: alternativeBankAccountNamesArrayLengthStart,
						MustLengthTo:   alternativeBankAccountNamesArrayLengthStop,
						Length:         length,
					}
				}
				return nil
			},
		)

		validateAlternativeBankAccountNamesElemLength = validation.By(
			func(value interface{}) error {
				name, _ := value.(string)
				length := len(name)
				if length != 0 &&
					!(length >= alternativeBankAccountNamesElemLengthStart &&
						length <= alternativeBankAccountNamesElemLengthStop) {
					return &InvalidAlternativeBankAccountElemLengthError{
						MustLengthFrom: alternativeBankAccountNamesElemLengthStart,
						MustLengthTo:   alternativeBankAccountNamesElemLengthStop,
						Length:         length,
					}
				}
				return nil
			},
		)

		validateAlternativeBankAccountNames = []validation.Rule{
			validateAlternativeBankAccountNamesArrayLength,
			validation.Each(validateAlternativeBankAccountNamesElemLength),
		}

		validateFirstNameLength = validation.By(
			func(value interface{}) error {
				name, _ := value.(string)
				length := len(name)
				if name != "" &&
					!(length >= firstNameLengthStart &&
						length <= firstNameLengthStop) {
					return &InvalidFirstNameLengthError{
						MustLengthFrom: firstNameLengthStart,
						MustLengthTo:   firstNameLengthStop,
						Length:         length,
					}
				}
				return nil
			},
		)

		validateFirstName = []validation.Rule{
			validateFirstNameLength,
			is.Alpha,
		}

		validateCustomerIDLength = validation.By(
			func(value interface{}) error {
				id, _ := value.(string)
				length := len(id)
				if id != "" &&
					!(length >= customerIDLengthStart &&
						length <= customerIDLengthStop) {
					return &InvalidCustomerIDLengthError{
						MustLengthFrom: customerIDLengthStart,
						MustLengthTo:   customerIDLengthStop,
						Length:         length,
					}
				}
				return nil
			},
		)

		validateCustomerID = []validation.Rule{
			validateCustomerIDLength,
		}
	)

	if err := validation.ValidateStruct(a,
		validation.Field(&a.Country, validateCountry...),
		validation.Field(
			&a.AlternativeBankAccountNames,
			validateAlternativeBankAccountNames...,
		),
		validation.Field(&a.FirstName, validateFirstName...),
		validation.Field(&a.CustomerID, validateCustomerID...),
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
