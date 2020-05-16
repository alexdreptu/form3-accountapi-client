package accountapi_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	. "github.com/alexdreptu/form3-accountapi-client"
)

// some default values
const (
	accountType = "accounts"
	accountBIC  = "NWBKGB22"
)

// for randomAlphanumeric
const uppercase = true

// for randomStringNumber
const startWithZero = true

// for randomBankIDItaly
const accountNumberPresent = true

func randomAlphanumeric(length int, uppercase ...bool) string {
	const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := make([]byte, length)
	for i := range chars {
		chars[i] = alphanumeric[rnd.Intn(len(alphanumeric))]
	}

	str := string(chars)
	if len(uppercase) != 0 && uppercase[0] {
		str = strings.ToUpper(str)
	}

	return str
}

func randomStringNumber(length int, startWithZero ...bool) string {
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

func randomChoose(x, y int) int {
	numbers := []int{x, y}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return numbers[random.Intn(2)]
}

// TODO: implement
func randomBIC() string {
	return ""
}

func randomBankIDUnitedKingdom() string {
	return randomStringNumber(BankIDLengthUnitedKingdom)
}

func randomBankIDAustralia() string {
	return randomStringNumber(BankIDLengthAustralia)
}

func randomBankIDBelgium() string {
	return randomStringNumber(BankIDLengthBelgium)
}

func randomBankIDCanada() string {
	return randomStringNumber(BankIDLengthCanada)
}

func randomBankIDFrance() string {
	return randomStringNumber(BankIDLengthFrance)
}

func randomBankIDGermany() string {
	return randomStringNumber(BankIDLengthGermany)
}

func randomBankIDGreece() string {
	return randomStringNumber(BankIDLengthGreece)
}

func randomBankIDHongKong() string {
	return randomStringNumber(BankIDLengthHongKong)
}

func randomBankIDItaly(accountNumberPresent ...bool) string {
	var length int
	if len(accountNumberPresent) != 0 && accountNumberPresent[0] {
		length = BankIDLengthItalyAccountNumberPresent
	} else {
		length = BankIDLengthItalyAccountNumberNotPresent
	}

	return randomStringNumber(length)
}

func randomBankIDLuxembourg() string {
	return randomStringNumber(BankIDLengthLuxembourg)
}

// returns empty string because BankID for Netherlands must be blank
func randomBankIDNetherlands() string {
	return ""
}

func randomBankIDPoland() string {
	return randomStringNumber(BankIDLengthPoland)
}

func randomBankIDPortugal() string {
	return randomStringNumber(BankIDLengthPortugal)
}

func randomBankIDSpain() string {
	return randomStringNumber(BankIDLengthSpain)
}

func randomBankIDSwitzerland() string {
	return randomStringNumber(BankIDLengthSwitzerland)
}

func randomBankIDUnitedStates() string {
	return randomStringNumber(BankIDLengthUnitedStates)
}

func randomAccountNumberUnitedKingdom() string {
	return randomStringNumber(AccountNumberLengthUnitedKingdom, startWithZero)
}

func randomAccountNumberAustralia() string {
	length := randomLength(AccountNumberLengthAustraliaStart, AccountNumberLengthAustraliaStop)
	return randomStringNumber(length)
}

func randomAccountNumberBelgium() string {
	return randomStringNumber(AccountNumberLengthBelgium)
}

func randomAccountNumberCanada() string {
	length := randomLength(AccountNumberLengthCanadaStart, AccountNumberLengthCanadaStop)
	return randomStringNumber(length)
}

func randomAccountNumberFrance() string {
	return randomStringNumber(AccountNumberLengthFrance)
}

func randomAccountNumberGermany() string {
	return randomStringNumber(AccountNumberLengthGermany)
}

func randomAccountNumberGreece() string {
	return randomStringNumber(AccountNumberLengthGreece)
}

func randomAccountNumberHongKong() string {
	length := randomLength(AccountNumberLengthHongKongStart, AccountNumberLengthHongKongStop)
	return randomStringNumber(length)
}

func randomAccountNumberItaly() string {
	return randomStringNumber(AccountNumberLengthItaly)
}

func randomAccountNumberLuxembourg() string {
	return randomStringNumber(AccountNumberLengthLuxembourg)
}

func randomAccountNumberNetherlands() string {
	return randomStringNumber(AccountNumberLengthNetherlands)
}

func randomAccountNumberPoland() string {
	return randomStringNumber(AccountNumberLengthPoland)
}

func randomAccountNumberPortugal() string {
	return randomStringNumber(AccountNumberLengthPortugal)
}

func randomAccountNumberSpain() string {
	return randomStringNumber(AccountNumberLengthSpain)
}

func randomAccountNumberSwitzerland() string {
	return randomStringNumber(AccountNumberLengthSwitzerland)
}

func randomAccountNumberUnitedStates() string {
	length := randomLength(AccountNumberLengthUnitedStatesStart, AccountNumberLengthUnitedStatesStop)
	return randomStringNumber(length)
}

// TODO: implement
func TestNewAccount(t *testing.T) {
	//
}
