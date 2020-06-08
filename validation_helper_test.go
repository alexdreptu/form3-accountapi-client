package accountapi_test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	. "github.com/alexdreptu/form3-accountapi-client"
)

const accountType = "accounts"

const (
	alpha   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric = "0123456789"
)

// alphanumeric styles
const (
	alphanumericStyleNormal = 0
	alphanumericStyleMix    = 1
	alphanumericStylePure   = 2
)

// for randomAlphanumeric
const uppercase = true

// for randomNumberString
const startWithZero = true

// for randomBankIDItaly
const accountNumberPresent = true

var firstNamesMale = []string{
	"Jacob", "Mason", "Ethan", "Noah", "William",
	"Liam", "Jayden", "Michael", "Alexander", "Aiden",
	"Daniel", "Matthew", "Elijah", "James", "Anthony",
	"Benjamin", "Joshua", "Andrew", "David", "Joseph",
}

var firstNamesFemale = []string{
	"Sophia", "Emma", "Isabella", "Olivia", "Ava",
	"Emily", "Abigail", "Mia", "Madison", "Elizabeth",
	"Chloe", "Ella", "Avery", "Addison", "Aubrey",
	"Lily", "Natalie", "Sofia", "Charlotte", "Zoey",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown",
	"Davis", "Miller", "Wilson", "Moore", "Taylor",
	"Anderson", "Thomas", "Jackson", "White", "Harris",
	"Martin", "Thompson", "Garcia", "Martinez", "Robinson",
}

func randomFirstName() string {
	length := len(firstNamesMale) + len(firstNamesFemale)
	names := make([]string, length)
	names = append(firstNamesMale, firstNamesFemale...)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	name := names[random.Intn(length)]
	return name
}

func randomFullName() string {
	length := len(firstNamesMale) + len(firstNamesFemale)
	firstNames := make([]string, length)
	firstNames = append(firstNamesMale, firstNamesFemale...)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	firstName := firstNames[random.Intn(length)]
	lastName := lastNames[random.Intn(len(lastNames))]
	fullName := firstName + " " + lastName
	return fullName
}

func randomAlphanumeric(length, style int, uppercase ...bool) string {
	chars := make([]byte, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	styleNormal := func() {
		const alphanumeric = alpha + numeric
		for i := range chars {
			chars[i] = alphanumeric[random.Intn(len(alphanumeric))]
		}
	}

	switch style {
	case alphanumericStyleNormal:
		styleNormal()

	case alphanumericStyleMix:
		which := randomBool()
		for i := range chars {
			if which {
				chars[i] = byte(alpha[random.Intn(len(alpha))])
				which = randomBool()
			} else {
				chars[i] = byte(numeric[random.Intn(len(numeric))])
				which = randomBool()
			}
		}

	case alphanumericStylePure:
		which := randomBool()
		for i := range chars {
			if which {
				chars[i] = byte(alpha[random.Intn(len(alpha))])
				which = !which
			} else {
				chars[i] = byte(numeric[random.Intn(len(numeric))])
				which = !which
			}
		}

	default:
		styleNormal()
	}

	str := string(chars)
	if len(uppercase) != 0 && uppercase[0] {
		str = strings.ToUpper(str)
	}

	return str
}

func randomAlpha(length int, uppercase ...bool) string {
	chars := make([]byte, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range chars {
		chars[i] = alpha[random.Intn(len(alpha))]
	}

	str := string(chars)
	if len(uppercase) != 0 && uppercase[0] {
		str = strings.ToUpper(str)
	}

	return str
}

func randomNumberString(length int, startWithZero ...bool) string {
	min := 1
	max := 9
	numbers := make([]string, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range numbers {
		n := random.Intn(max-min) + min
		numbers[i] = fmt.Sprint(n)
	}

	if len(startWithZero) != 0 && startWithZero[0] {
		numbers[0] = "0"
	}

	return strings.Join(numbers, "")
}

func randomLength(min, max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max+1-min) + min
}

func randomChoiceInt(x, y int) int {
	numbers := []int{x, y}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return numbers[random.Intn(len(numbers))]
}

func randomBIC() string {
	length := randomChoiceInt(BICLength8, BICLength11)
	str := randomAlpha(6, uppercase)
	if length == BICLength8 {
		str += randomAlphanumeric(2, alphanumericStyleNormal, uppercase)
	} else if length == BICLength11 {
		str += randomAlphanumeric(5, alphanumericStyleNormal, uppercase)
	}
	return str
}

func randomBICInvalid() string {
	length := randomChoiceInt(BICLength8, BICLength11)
	return randomAlphanumeric(length, alphanumericStylePure, uppercase)
}

func randomBaseCurrencyInvalid() string {
	currencies := []string{
		CurrencyUnitedKingdom,
		CurrencyAustralia,
		CurrencyBelgium,
		CurrencyCanada,
		CurrencyFrance,
		CurrencyGermany,
		CurrencyGreecee,
		CurrencyHongKong,
		CurrencyItaly,
		CurrencyLuxembourg,
		CurrencyNetherlands,
		CurrencyPoland,
		CurrencyPortugal,
		CurrencySpain,
		CurrencySwitzerland,
		CurrencyUnitedStates,
	}

	currency := func() string {
		var currency string
		for {
			currency = randomAlpha(3, uppercase)
			for _, c := range currencies {
				if currency != c {
					return currency
				}
			}
		}
	}()

	return currency
}

func randomBankIDCodeInvalid() string {
	codes := []string{
		BankIDCodeUnitedKingdom,
		BankIDCodeAustralia,
		BankIDCodeBelgium,
		BankIDCodeCanada,
		BankIDCodeFrance,
		BankIDCodeGermany,
		BankIDCodeGreece,
		BankIDCodeHongKong,
		BankIDCodeItaly,
		BankIDCodeLuxembourg,
		BankIDCodeNetherlands,
		BankIDCodePoland,
		BankIDCodePortugal,
		BankIDCodeSpain,
		BankIDCodeSwitzerland,
		BankIDCodeUnitedStates,
	}

	code := func() string {
		var code string
		for {
			code = randomAlpha(randomLength(2, 5), uppercase)
			for _, c := range codes {
				if code != c {
					return code
				}
			}
		}
	}()

	return code
}

func randomCustomerID() string {
	return randomAlphanumeric(randomLength(5, 15), alphanumericStyleNormal, uppercase)
}

func randomAlternativeBankAccountNames(length ...int) []string {
	var mustLength int
	if len(length) != 0 {
		mustLength = length[0]
	} else {
		mustLength = randomLength(1, 3)
	}
	names := make([]string, mustLength)
	for i := range names {
		names[i] = randomFullName()
	}
	return names
}

func randomBool() bool {
	booleans := []bool{true, false}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return booleans[random.Intn(len(booleans))]
}

func randomBankIDUnitedKingdom() string {
	return randomNumberString(BankIDLengthUnitedKingdom)
}

func randomBankIDAustralia() string {
	return randomNumberString(BankIDLengthAustralia)
}

func randomBankIDBelgium() string {
	return randomNumberString(BankIDLengthBelgium)
}

func randomBankIDCanada() string {
	return randomNumberString(BankIDLengthCanada, startWithZero)
}

func randomBankIDFrance() string {
	return randomNumberString(BankIDLengthFrance)
}

func randomBankIDGermany() string {
	return randomNumberString(BankIDLengthGermany)
}

func randomBankIDGreece() string {
	return randomNumberString(BankIDLengthGreece)
}

func randomBankIDHongKong() string {
	return randomNumberString(BankIDLengthHongKong)
}

func randomBankIDItaly(accountNumberPresent ...bool) string {
	var length int
	if len(accountNumberPresent) != 0 && accountNumberPresent[0] {
		length = BankIDLengthItalyAccountNumberPresent
	} else {
		length = BankIDLengthItalyAccountNumberNotPresent
	}
	return randomNumberString(length)
}

func randomBankIDLuxembourg() string {
	return randomNumberString(BankIDLengthLuxembourg)
}

// returns empty string because BankID for Netherlands must be blank
func randomBankIDNetherlands() string {
	return ""
}

func randomBankIDPoland() string {
	return randomNumberString(BankIDLengthPoland)
}

func randomBankIDPortugal() string {
	return randomNumberString(BankIDLengthPortugal)
}

func randomBankIDSpain() string {
	return randomNumberString(BankIDLengthSpain)
}

func randomBankIDSwitzerland() string {
	return randomNumberString(BankIDLengthSwitzerland)
}

func randomBankIDUnitedStates() string {
	return randomNumberString(BankIDLengthUnitedStates)
}

func randomAccountNumberUnitedKingdom() string {
	return randomNumberString(AccountNumberLengthUnitedKingdom)
}

func randomAccountNumberAustralia() string {
	length := randomLength(
		AccountNumberLengthAustraliaStart,
		AccountNumberLengthAustraliaStop,
	)
	return randomNumberString(length)
}

func randomAccountNumberBelgium() string {
	return randomNumberString(AccountNumberLengthBelgium)
}

func randomAccountNumberCanada() string {
	length := randomLength(
		AccountNumberLengthCanadaStart,
		AccountNumberLengthCanadaStop,
	)
	return randomNumberString(length)
}

func randomAccountNumberFrance() string {
	return randomNumberString(AccountNumberLengthFrance)
}

func randomAccountNumberGermany() string {
	return randomNumberString(AccountNumberLengthGermany)
}

func randomAccountNumberGreece() string {
	return randomNumberString(AccountNumberLengthGreece)
}

func randomAccountNumberHongKong() string {
	length := randomLength(
		AccountNumberLengthHongKongStart,
		AccountNumberLengthHongKongStop,
	)
	return randomNumberString(length)
}

func randomAccountNumberItaly() string {
	return randomNumberString(AccountNumberLengthItaly)
}

func randomAccountNumberLuxembourg() string {
	return randomNumberString(AccountNumberLengthLuxembourg)
}

func randomAccountNumberNetherlands() string {
	return randomNumberString(AccountNumberLengthNetherlands)
}

func randomAccountNumberPoland() string {
	return randomNumberString(AccountNumberLengthPoland)
}

func randomAccountNumberPortugal() string {
	return randomNumberString(AccountNumberLengthPortugal)
}

func randomAccountNumberSpain() string {
	return randomNumberString(AccountNumberLengthSpain)
}

func randomAccountNumberSwitzerland() string {
	return randomNumberString(AccountNumberLengthSwitzerland)
}

func randomAccountNumberUnitedStates() string {
	length := randomLength(
		AccountNumberLengthUnitedStatesStart,
		AccountNumberLengthUnitedStatesStop,
	)
	return randomNumberString(length)
}
