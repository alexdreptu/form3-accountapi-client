package accountapi_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/alexdreptu/form3-accountapi-client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestNewClient(t *testing.T) {
	// test if without passing url to NewClient, client.BaseURL is set to default
	client := NewClient(&http.Client{})
	require.IsType(t, &Client{}, client)
	assert.Equal(t, DefaultBaseURL(), client.BaseURL)

	// test if without passing an url string to NewClient, client.BaseURL is set to default
	client = NewClient(&http.Client{}, "")
	assert.Equal(t, DefaultBaseURL(), client.BaseURL)
}

type CreateAccountSuite struct {
	suite.Suite
	testAccount testOptions
}

func (s *CreateAccountSuite) SetupSuite() {
	s.testAccount = testOptions{
		accType:                        accountType,
		accID:                          uuid.New().String(),
		accOrganisationID:              uuid.New().String(),
		accCountry:                     CountryUnitedKingdom,
		accBIC:                         randomBIC(),
		accBankID:                      randomBankIDUnitedKingdom(),
		accBankIDCode:                  BankIDCodeUnitedKingdom,
		accAccountNumber:               randomAccountNumberUnitedKingdom(),
		accBaseCurrency:                CurrencyUnitedKingdom,
		accJointAccount:                randomBool(),
		accFirstName:                   randomFirstName(),
		accAlternativeBankAccountNames: randomAlternativeBankAccountNames(),
		accAccountMatchingOptOut:       randomBool(),
		accCustomerID:                  randomCustomerID(),
	}
}

func (s *CreateAccountSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := NewClient(&http.Client{})
	err := client.DeleteAccount(ctx, s.testAccount.accID, 0)
	s.Require().NoError(err)
}

func (s *CreateAccountSuite) TestCreateAccount() {
	options := &Options{
		ID:             s.testAccount.accID,
		OrganisationID: s.testAccount.accOrganisationID,
		Type:           s.testAccount.accType,
		Attributes: []Attribute{
			WithAttrCountry(s.testAccount.accCountry),
			WithAttrBIC(s.testAccount.accBIC),
			WithAttrBankID(s.testAccount.accBankID),
			WithAttrBankIDCode(s.testAccount.accBankIDCode),
			WithAttrAccountNumber(s.testAccount.accAccountNumber),
			WithAttrBaseCurrency(s.testAccount.accBaseCurrency),
			WithAttrJointAccount(s.testAccount.accJointAccount),
			WithAttrFirstName(s.testAccount.accFirstName),
			WithAttrAlternativeBankAccountNames(s.testAccount.accAlternativeBankAccountNames...),
			WithAttrAccountMatchingOptOut(s.testAccount.accAccountMatchingOptOut),
			WithAttrCustomerID(s.testAccount.accCustomerID),
		},
	}

	account, err := NewAccount(options)
	s.Require().NoError(err)
	client := NewClient(&http.Client{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	createdAccount, err := client.CreateAccount(ctx, account)
	s.Require().NoError(err)
	s.Require().IsType(Account{}, createdAccount)

	s.Assert().Equal(s.testAccount.accType, createdAccount.Data.Type)
	s.Assert().Equal(s.testAccount.accID, createdAccount.Data.ID)
	s.Assert().Equal(s.testAccount.accOrganisationID, createdAccount.Data.OrganisationID)
	s.Assert().Equal(s.testAccount.accCountry, createdAccount.Data.Attributes.Country)
	s.Assert().Equal(s.testAccount.accBIC, createdAccount.Data.Attributes.BIC)
	s.Assert().Equal(s.testAccount.accBankID, createdAccount.Data.Attributes.BankID)
	s.Assert().Equal(s.testAccount.accBankIDCode, createdAccount.Data.Attributes.BankIDCode)
	s.Assert().Equal(s.testAccount.accAccountNumber, createdAccount.Data.Attributes.AccountNumber)
	s.Assert().Equal(s.testAccount.accBaseCurrency, createdAccount.Data.Attributes.BaseCurrency)
	s.Assert().Equal(s.testAccount.accJointAccount, createdAccount.Data.Attributes.JointAccount)
	s.Assert().Equal(s.testAccount.accJointAccount, createdAccount.Data.Attributes.JointAccount)
	s.Assert().Equal(s.testAccount.accFirstName, createdAccount.Data.Attributes.FirstName)
	s.Assert().Equal(s.testAccount.accAlternativeBankAccountNames, createdAccount.Data.Attributes.AlternativeBankAccountNames)
	s.Assert().Equal(s.testAccount.accAccountMatchingOptOut, createdAccount.Data.Attributes.AccountMatchingOptOut)
	s.Assert().Equal(s.testAccount.accCustomerID, createdAccount.Data.Attributes.CustomerID)
}

func (s *CreateAccountSuite) TestCreateAccount_InvalidURL() {
	options := &Options{
		ID:             uuid.New().String(),
		OrganisationID: s.testAccount.accOrganisationID,
		Type:           s.testAccount.accType,
		Attributes: []Attribute{
			WithAttrCountry(s.testAccount.accCountry),
			WithAttrBIC(s.testAccount.accBIC),
			WithAttrBankID(s.testAccount.accBankID),
			WithAttrBankIDCode(s.testAccount.accBankIDCode),
			WithAttrAccountNumber(s.testAccount.accAccountNumber),
			WithAttrBaseCurrency(s.testAccount.accBaseCurrency),
			WithAttrJointAccount(s.testAccount.accJointAccount),
			WithAttrFirstName(s.testAccount.accFirstName),
			WithAttrAlternativeBankAccountNames(s.testAccount.accAlternativeBankAccountNames...),
			WithAttrAccountMatchingOptOut(s.testAccount.accAccountMatchingOptOut),
			WithAttrCustomerID(s.testAccount.accCustomerID),
		},
	}

	account, err := NewAccount(options)
	s.Require().NoError(err)

	client := NewClient(&http.Client{}, DefaultBaseURL()+"x")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.CreateAccount(ctx, account)
	s.Require().Error(err)
	var e *ResourceNotExistsError
	s.Assert().True(errors.As(err, &e))

	client.BaseURL = DefaultBaseURL() + "x"
	_, err = client.CreateAccount(ctx, account)
	s.Require().Error(err)
	s.Assert().True(errors.As(err, &e))
}

func (s *CreateAccountSuite) TestCreateAccount_DuplicateAccount() {
	options := &Options{
		ID:             s.testAccount.accID,
		OrganisationID: s.testAccount.accOrganisationID,
		Type:           s.testAccount.accType,
		Attributes: []Attribute{
			WithAttrCountry(s.testAccount.accCountry),
			WithAttrBIC(s.testAccount.accBIC),
			WithAttrBankID(s.testAccount.accBankID),
			WithAttrBankIDCode(s.testAccount.accBankIDCode),
			WithAttrAccountNumber(s.testAccount.accAccountNumber),
			WithAttrBaseCurrency(s.testAccount.accBaseCurrency),
			WithAttrJointAccount(s.testAccount.accJointAccount),
			WithAttrFirstName(s.testAccount.accFirstName),
			WithAttrAlternativeBankAccountNames(s.testAccount.accAlternativeBankAccountNames...),
			WithAttrAccountMatchingOptOut(s.testAccount.accAccountMatchingOptOut),
			WithAttrCustomerID(s.testAccount.accCustomerID),
		},
	}

	account, err := NewAccount(options)
	s.Require().NoError(err)
	client := NewClient(&http.Client{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.CreateAccount(ctx, account)
	s.Require().Error(err)
	var e *DuplicateAccountError
	s.Assert().True(errors.As(err, &e))
}

func TestCreateAccountSuite(t *testing.T) {
	suite.Run(t, &CreateAccountSuite{})
}

type FetchAccountSuite struct {
	suite.Suite
	testAccount testOptions
}

func (s *FetchAccountSuite) SetupSuite() {
	testAccount := testOptions{
		accType:                        accountType,
		accID:                          uuid.New().String(),
		accOrganisationID:              uuid.New().String(),
		accCountry:                     CountryUnitedKingdom,
		accBIC:                         randomBIC(),
		accBankID:                      randomBankIDUnitedKingdom(),
		accBankIDCode:                  BankIDCodeUnitedKingdom,
		accAccountNumber:               randomAccountNumberUnitedKingdom(),
		accBaseCurrency:                CurrencyUnitedKingdom,
		accJointAccount:                randomBool(),
		accFirstName:                   randomFirstName(),
		accAlternativeBankAccountNames: randomAlternativeBankAccountNames(),
		accAccountMatchingOptOut:       randomBool(),
		accCustomerID:                  randomCustomerID(),
	}

	s.testAccount = testAccount

	options := &Options{
		ID:             testAccount.accID,
		OrganisationID: testAccount.accOrganisationID,
		Type:           testAccount.accType,
		Attributes: []Attribute{
			WithAttrCountry(testAccount.accCountry),
			WithAttrBIC(testAccount.accBIC),
			WithAttrBankID(testAccount.accBankID),
			WithAttrBankIDCode(testAccount.accBankIDCode),
			WithAttrAccountNumber(testAccount.accAccountNumber),
			WithAttrBaseCurrency(testAccount.accBaseCurrency),
			WithAttrJointAccount(testAccount.accJointAccount),
			WithAttrFirstName(testAccount.accFirstName),
			WithAttrAlternativeBankAccountNames(testAccount.accAlternativeBankAccountNames...),
			WithAttrAccountMatchingOptOut(testAccount.accAccountMatchingOptOut),
			WithAttrCustomerID(testAccount.accCustomerID),
		},
	}

	account, err := NewAccount(options)
	s.Require().NoError(err)
	client := NewClient(&http.Client{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.CreateAccount(ctx, account)
	s.Require().NoError(err)
}

func (s *FetchAccountSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := NewClient(&http.Client{})
	err := client.DeleteAccount(ctx, s.testAccount.accID, 0)
	s.Require().NoError(err)
}

func (s *FetchAccountSuite) TestFetchAccount() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := NewClient(&http.Client{})
	fetchedAccount, err := client.FetchAccount(ctx, s.testAccount.accID)
	s.Require().NoError(err)
	s.Require().IsType(Account{}, fetchedAccount)

	s.Assert().Equal(s.testAccount.accType, fetchedAccount.Data.Type)
	s.Assert().Equal(s.testAccount.accID, fetchedAccount.Data.ID)
	s.Assert().Equal(s.testAccount.accOrganisationID, fetchedAccount.Data.OrganisationID)
	s.Assert().Equal(s.testAccount.accCountry, fetchedAccount.Data.Attributes.Country)
	s.Assert().Equal(s.testAccount.accBIC, fetchedAccount.Data.Attributes.BIC)
	s.Assert().Equal(s.testAccount.accBankID, fetchedAccount.Data.Attributes.BankID)
	s.Assert().Equal(s.testAccount.accBankIDCode, fetchedAccount.Data.Attributes.BankIDCode)
	s.Assert().Equal(s.testAccount.accAccountNumber, fetchedAccount.Data.Attributes.AccountNumber)
	s.Assert().Equal(s.testAccount.accBaseCurrency, fetchedAccount.Data.Attributes.BaseCurrency)
	s.Assert().Equal(s.testAccount.accJointAccount, fetchedAccount.Data.Attributes.JointAccount)
	s.Assert().Equal(s.testAccount.accJointAccount, fetchedAccount.Data.Attributes.JointAccount)
	s.Assert().Equal(s.testAccount.accFirstName, fetchedAccount.Data.Attributes.FirstName)
	s.Assert().Equal(s.testAccount.accAlternativeBankAccountNames, fetchedAccount.Data.Attributes.AlternativeBankAccountNames)
	s.Assert().Equal(s.testAccount.accAccountMatchingOptOut, fetchedAccount.Data.Attributes.AccountMatchingOptOut)
	s.Assert().Equal(s.testAccount.accCustomerID, fetchedAccount.Data.Attributes.CustomerID)
}

func (s *FetchAccountSuite) TestFetchAccount_InvalidURL() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := NewClient(&http.Client{}, DefaultBaseURL()+"x")
	_, err := client.FetchAccount(ctx, uuid.New().String())
	s.Require().Error(err)
	var e *ResourceNotExistsError
	s.Require().True(errors.As(err, &e))
}

func TestFetchAccountSuite(t *testing.T) {
	suite.Run(t, &FetchAccountSuite{})
}

type ListAccountsSuite struct {
	suite.Suite
	ids []string
}

func (s *ListAccountsSuite) SetupSuite() {
	const numberOfAccounts = 30

	s.ids = make([]string, numberOfAccounts)
	for i := range s.ids {
		s.ids[i] = uuid.New().String()
	}

	testAccounts := make([]testOptions, len(s.ids))
	for i := range testAccounts {
		testAccounts[i] = testOptions{
			accType:                        accountType,
			accID:                          s.ids[i],
			accOrganisationID:              uuid.New().String(),
			accCountry:                     CountryUnitedKingdom,
			accBIC:                         randomBIC(),
			accBankID:                      randomBankIDUnitedKingdom(),
			accBankIDCode:                  BankIDCodeUnitedKingdom,
			accAccountNumber:               randomAccountNumberUnitedKingdom(),
			accBaseCurrency:                CurrencyUnitedKingdom,
			accJointAccount:                randomBool(),
			accFirstName:                   randomFirstName(),
			accAlternativeBankAccountNames: randomAlternativeBankAccountNames(),
			accAccountMatchingOptOut:       randomBool(),
			accCustomerID:                  randomCustomerID(),
		}
	}

	for _, ta := range testAccounts {
		options := &Options{
			ID:             ta.accID,
			OrganisationID: ta.accOrganisationID,
			Type:           ta.accType,
			Attributes: []Attribute{
				WithAttrCountry(ta.accCountry),
				WithAttrBIC(ta.accBIC),
				WithAttrBankID(ta.accBankID),
				WithAttrBankIDCode(ta.accBankIDCode),
				WithAttrAccountNumber(ta.accAccountNumber),
				WithAttrBaseCurrency(ta.accBaseCurrency),
				WithAttrJointAccount(ta.accJointAccount),
				WithAttrFirstName(ta.accFirstName),
				WithAttrAlternativeBankAccountNames(ta.accAlternativeBankAccountNames...),
				WithAttrAccountMatchingOptOut(ta.accAccountMatchingOptOut),
				WithAttrCustomerID(ta.accCustomerID),
			},
		}

		account, err := NewAccount(options)
		s.Require().NoError(err)
		client := NewClient(&http.Client{})
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_, err = client.CreateAccount(ctx, account)
		s.Require().NoError(err)
	}
}

func (s *ListAccountsSuite) TearDownSuite() {
	for _, id := range s.ids {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		client := NewClient(&http.Client{})
		err := client.DeleteAccount(ctx, id, 0)
		s.Require().NoError(err)
	}
}

func (s *ListAccountsSuite) TestListAccounts() {
	const selfLink = "/v1/organisation/accounts?page%%5Bnumber%%5D=%d&page%%5Bsize%%5D=%d"

	testCases := []struct {
		name       string
		pageNumber int
		pageSize   int
	}{
		{name: "page number 0 page size 5", pageNumber: 0, pageSize: 5},
		{name: "page number 1 page size 10", pageNumber: 1, pageSize: 10},
		{name: "page number 2 page size 10", pageNumber: 2, pageSize: 10},
		{name: "page number 0 page size 30", pageNumber: 0, pageSize: 30},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			client := NewClient(&http.Client{})
			accountList, err := client.ListAccounts(ctx, tc.pageNumber, tc.pageSize)
			s.Require().NoError(err)
			s.Require().IsType(Accounts{}, accountList)
			s.Require().NotEmpty(accountList)

			self := fmt.Sprintf(selfLink, tc.pageNumber, tc.pageSize)
			s.Assert().Equal(self, accountList.Links.Self)
		})
	}
}

func (s *ListAccountsSuite) TestListAccounts_InvalidURL() {
	// test if it returns an error when an invalid url is passed to NewClient
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := NewClient(&http.Client{}, DefaultBaseURL()+"x")
	_, err := client.ListAccounts(ctx, 1, 5)
	s.Require().Error(err)
	// test if the error returned is ResourceNotExists
	var e *ResourceNotExistsError
	s.Assert().True(errors.As(err, &e))

	client.BaseURL = DefaultBaseURL() + "x"
	_, err = client.ListAccounts(ctx, 1, 5)
	s.Assert().True(errors.As(err, &e))
}

func TestListAccountsSuite(t *testing.T) {
	suite.Run(t, &ListAccountsSuite{})
}

type DeleteAccountSuite struct {
	suite.Suite
	testAccount testOptions
}

func (s *DeleteAccountSuite) SetupSuite() {
	testAccount := testOptions{
		accType:                        accountType,
		accID:                          uuid.New().String(),
		accOrganisationID:              uuid.New().String(),
		accCountry:                     CountryUnitedKingdom,
		accBIC:                         randomBIC(),
		accBankID:                      randomBankIDUnitedKingdom(),
		accBankIDCode:                  BankIDCodeUnitedKingdom,
		accAccountNumber:               randomAccountNumberUnitedKingdom(),
		accBaseCurrency:                CurrencyUnitedKingdom,
		accJointAccount:                randomBool(),
		accFirstName:                   randomFirstName(),
		accAlternativeBankAccountNames: randomAlternativeBankAccountNames(),
		accAccountMatchingOptOut:       randomBool(),
		accCustomerID:                  randomCustomerID(),
	}

	s.testAccount = testAccount

	options := &Options{
		ID:             testAccount.accID,
		OrganisationID: testAccount.accOrganisationID,
		Type:           testAccount.accType,
		Attributes: []Attribute{
			WithAttrCountry(testAccount.accCountry),
			WithAttrBIC(testAccount.accBIC),
			WithAttrBankID(testAccount.accBankID),
			WithAttrBankIDCode(testAccount.accBankIDCode),
			WithAttrAccountNumber(testAccount.accAccountNumber),
			WithAttrBaseCurrency(testAccount.accBaseCurrency),
			WithAttrJointAccount(testAccount.accJointAccount),
			WithAttrFirstName(testAccount.accFirstName),
			WithAttrAlternativeBankAccountNames(testAccount.accAlternativeBankAccountNames...),
			WithAttrAccountMatchingOptOut(testAccount.accAccountMatchingOptOut),
			WithAttrCustomerID(testAccount.accCustomerID),
		},
	}

	account, err := NewAccount(options)
	s.Require().NoError(err)
	client := NewClient(&http.Client{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.CreateAccount(ctx, account)
	s.Require().NoError(err)
}

func (s *DeleteAccountSuite) TestDeleteAccount() {
	const version = 0

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client := NewClient(&http.Client{})

	// invalid account id format
	err := client.DeleteAccount(ctx, randomAlphanumeric(36, alphanumericStyleMix), version)
	s.Require().Error(err)

	// incorrect version
	err = client.DeleteAccount(ctx, s.testAccount.accID, 1)
	s.Require().Error(err)
	var e *ResourceNotExistsError
	s.Require().True(errors.As(err, &e))

	// delete account
	err = client.DeleteAccount(ctx, s.testAccount.accID, version)
	s.Require().NoError(err)
}

func (s *DeleteAccountSuite) TestDeleteAccount_InvalidURL() {
	const version = 0

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client := NewClient(&http.Client{}, DefaultBaseURL()+"x")
	err := client.DeleteAccount(ctx, uuid.New().String(), version)
	s.Require().Error(err)
	var e *ResourceNotExistsError
	s.Require().True(errors.As(err, &e))

	client.BaseURL = DefaultBaseURL() + "x"
	err = client.DeleteAccount(ctx, uuid.New().String(), version)
	s.Require().Error(err)
	s.Require().True(errors.As(err, &e))
}

func TestDeleteAccountSuite(t *testing.T) {
	suite.Run(t, &DeleteAccountSuite{})
}

func TestWithAttrFunctions(t *testing.T) {
	attributes := &Attributes{}

	country := CountryUnitedKingdom
	bankID := randomBankIDUnitedKingdom()
	bankIDCode := BankIDCodeUnitedKingdom
	bic := randomBIC()
	accountNumber := randomAccountNumberUnitedKingdom()
	currency := CurrencyUnitedKingdom
	jointAccount := randomBool()
	firstName := randomFirstName()
	alternativeBankAccountNames := randomAlternativeBankAccountNames()
	accountMatchingOptOut := randomBool()
	customerID := randomCustomerID()

	attrs := []Attribute{
		WithAttrCountry(country),
		WithAttrBankID(bankID),
		WithAttrBankIDCode(bankIDCode),
		WithAttrBIC(bic),
		WithAttrAccountNumber(accountNumber),
		WithAttrBaseCurrency(currency),
		WithAttrJointAccount(jointAccount),
		WithAttrFirstName(firstName),
		WithAttrAlternativeBankAccountNames(alternativeBankAccountNames...),
		WithAttrAccountMatchingOptOut(accountMatchingOptOut),
		WithAttrCustomerID(customerID),
	}

	for _, attr := range attrs {
		attr(attributes)
	}

	assert.Equal(t, country, attributes.Country)
	assert.Equal(t, bankID, attributes.BankID)
	assert.Equal(t, bankIDCode, attributes.BankIDCode)
	assert.Equal(t, bic, attributes.BIC)
	assert.Equal(t, accountNumber, attributes.AccountNumber)
	assert.Equal(t, currency, attributes.BaseCurrency)
	assert.Equal(t, jointAccount, attributes.JointAccount)
	assert.Equal(t, firstName, attributes.FirstName)
	assert.Equal(t, alternativeBankAccountNames, attributes.AlternativeBankAccountNames)
	assert.Equal(t, customerID, attributes.CustomerID)
}
