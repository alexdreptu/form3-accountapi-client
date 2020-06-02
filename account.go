// Package accountapi provides a client for the fake Form3 Finacial Cloud RESTFul API.
// It implements only create, fetch, list and delete.
package accountapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/google/uuid"
)

const DefaultBaseURL = "http://localhost:8080/v1/organisation/accounts"

// ISO 3166-1 country codes
const (
	CountryUnitedKingdom = "GB"
	CountryAustralia     = "AU"
	CountryBelgium       = "BE"
	CountryCanada        = "CA"
	CountryFrance        = "FR"
	CountryGermany       = "DE"
	CountryGreece        = "GR"
	CountryHongKong      = "HK"
	CountryItaly         = "IT"
	CountryLuxembourg    = "LU"
	CountryNetherlands   = "NL"
	CountryPoland        = "PL"
	CountryPortugal      = "PT"
	CountrySpain         = "ES"
	CountrySwitzerland   = "CH"
	CountryUnitedStates  = "US"
)

// ISO 4217 codes
const (
	CurrencyUnitedKingdom = "GBP"
	CurrencyAustralia     = "AUD"
	CurrencyBelgium       = "EUR"
	CurrencyCanada        = "CAD"
	CurrencyFrance        = "EUR"
	CurrencyGermany       = "EUR"
	CurrencyGreecee       = "EUR"
	CurrencyHongKong      = "HKD"
	CurrencyItaly         = "EUR"
	CurrencyLuxembourg    = "EUR"
	CurrencyNetherlands   = "EUR"
	CurrencyPoland        = "PLN"
	CurrencyPortugal      = "EUR"
	CurrencySpain         = "EUR"
	CurrencySwitzerland   = "CHF"
	CurrencyUnitedStates  = "USD"
)

// Bank ID codes for each country
const (
	BankIDCodeUnitedKingdom = "GBDSC"
	BankIDCodeAustralia     = "AUBSB"
	BankIDCodeBelgium       = "BE"
	BankIDCodeCanada        = "CACPA"
	BankIDCodeFrance        = "FR"
	BankIDCodeGermany       = "DEBLZ"
	BankIDCodeGreece        = "GRBIC"
	BankIDCodeHongKong      = "HKNCC"
	BankIDCodeItaly         = "ITNCC"
	BankIDCodeLuxembourg    = "LULUX"
	BankIDCodeNetherlands   = "" // not supported, must be blank
	BankIDCodePoland        = "PLKNR"
	BankIDCodePortugal      = "PTNCC"
	BankIDCodeSpain         = "ESNCC"
	BankIDCodeSwitzerland   = "CHBCC"
	BankIDCodeUnitedStates  = "USABA"
)

// Bank ID lengths for each country
const (
	BankIDLengthUnitedKingdom = 6
	BankIDLengthAustralia     = 6
	BankIDLengthBelgium       = 3
	BankIDLengthCanada        = 9
	BankIDLengthFrance        = 10
	BankIDLengthGermany       = 8
	BankIDLengthGreece        = 7
	BankIDLengthHongKong      = 3
	//
	BankIDLengthItalyAccountNumberPresent    = 11
	BankIDLengthItalyAccountNumberNotPresent = 10
	//
	BankIDLengthLuxembourg   = 3
	BankIDLengthNetherlands  = 0 // 0 because not supported
	BankIDLengthPoland       = 8
	BankIDLengthPortugal     = 8
	BankIDLengthSpain        = 8
	BankIDLengthSwitzerland  = 5
	BankIDLengthUnitedStates = 9
)

// Account Number lengths for each country
const (
	AccountNumberLengthUnitedKingdom     = 8
	AccountNumberLengthAustraliaStart    = 6
	AccountNumberLengthAustraliaStop     = 10
	AccountNumberLengthBelgium           = 7
	AccountNumberLengthCanadaStart       = 7
	AccountNumberLengthCanadaStop        = 12
	AccountNumberLengthFrance            = 10
	AccountNumberLengthGermany           = 7
	AccountNumberLengthGreece            = 16
	AccountNumberLengthHongKongStart     = 9
	AccountNumberLengthHongKongStop      = 12
	AccountNumberLengthItaly             = 12
	AccountNumberLengthLuxembourg        = 13
	AccountNumberLengthNetherlands       = 10
	AccountNumberLengthPoland            = 16
	AccountNumberLengthPortugal          = 11
	AccountNumberLengthSpain             = 10
	AccountNumberLengthSwitzerland       = 12
	AccountNumberLengthUnitedStatesStart = 6
	AccountNumberLengthUnitedStatesStop  = 17
)

// Options holds the options meant to be passed as an argument to NewAccount
type Options struct {
	Type           string
	ID             string
	OrganisationID string
	Attributes     []Attribute
}

// Attributes holds the account attributes
type Attributes struct {
	// ISO 3166-1 code used to identify the domicile of the account, e.g. 'GB', 'FR'.
	// For more info see https://www.iso.org/iso-3166-country-codes.html
	Country string `json:"country,omitempty"`

	// ISO 4217 code used to identify the base currency of the account, e.g. 'GBP', 'EUR'.
	// For more info see https://www.iso.org/iso-4217-currency-codes.html
	BaseCurrency string `json:"base_currency,omitempty"`

	// Local country bank identifier. Format depends on the country.
	// Required for most countries.
	BankID string `json:"bank_id,omitempty"`

	// Identifies the type of bank ID being used.
	// See https://api-docs.form3.tech/api.html?python#accounts-create-data-table
	// for allowed value for each country. Required value depends on country attribute.
	BankIDCode string `json:"bank_id_code,omitempty"`

	// Account number. A unique account number will automatically be generated if not provided.
	AccountNumber string `json:"account_number,omitempty"`

	// SWIFT BIC in either 8 or 11 character format e.g. 'NWBKGB22'.
	BIC string `json:"bic,omitempty"`

	// A free-format reference that can be used to link this account to an external system.
	CustomerID string `json:"customer_id,omitempty"`

	// First name of the account holder
	FirstName string `json:"first_name,omitempty"`

	// Alternative primary account names, only used for UK Confirmation of Payee
	// CoP: Up to 3 alternative account names, one in each line of the array.
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names,omitempty"`

	// Flag to indicate if the account is a joint account,
	// only used for Confirmation of Payee (CoP).
	// CoP: Set to true is this is a joint account. Defaults to false.
	JointAccount bool `json:"joint_account,omitempty"`

	// Flag to indicate if the account has opted out of account matching,
	// only used for Confirmation of Payee.
	// CoP: Set to true if the account has opted out of account matching. Defaults to false.
	AccountMatchingOptOut bool `json:"account_matching_opt_out,omitempty"`
}

type Attribute func(*Attributes)

type Links struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Self  string `json:"self,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

type Details struct {
	Type           string `json:"type,omitempty"`
	ID             string `json:"id,omitempty"`
	CreatedOn      string `json:"created_on,omitempty"`
	ModifiedOn     string `json:"modified_on,omitempty"`
	OrganisationID string `json:"organisation_id,omitempty"`
	Version        int    `json:"version,omitempty"`
}

type Data struct {
	Attributes *Attributes `json:"attributes,omitempty"`
	Details
}

type Account struct {
	Data         *Data  `json:"data,omitempty"`
	Links        *Links `json:"links,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Accounts struct {
	Data  []Data `json:"data"`
	Links Links  `json:"links"`
}

type Client struct {
	Client  *http.Client
	BaseURL string
}

func (c *Client) CreateAccount(ctx context.Context, account *Account) (Account, error) {
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return Account{}, err
	}

	data, err := json.Marshal(account)
	if err != nil {
		return Account{}, err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return Account{}, err
	}

	req.Header.Set("Accept", "vnd.api+json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	// fmt.Println(resp.StatusCode)
	switch resp.StatusCode {
	case http.StatusNotFound:
		return Account{}, &ResourceNotExistsError{baseURL.String()}
	case http.StatusConflict:
		return Account{}, &DuplicateAccountError{account.Data.ID}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	var a Account
	if err := json.Unmarshal(body, &a); err != nil {
		return Account{}, err
	}

	return a, nil
}

func (c *Client) FetchAccount(ctx context.Context, id string) (Account, error) {
	accountID, err := uuid.Parse(id)
	if err != nil {
		return Account{}, err
	}

	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return Account{}, err
	}

	baseURL.Path = path.Join(baseURL.Path, accountID.String())

	req, err := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return Account{}, err
	}

	req.Header.Set("Accept", "vnd.api+json")

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return Account{}, &RecordNotExistsError{id}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	var a Account
	if err := json.Unmarshal(body, &a); err != nil {
		return Account{}, err
	}

	return a, nil
}

func (c *Client) ListAccounts(ctx context.Context, pageNumber, pageSize int) (Accounts, error) {
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return Accounts{}, err
	}

	params := url.Values{}
	params.Add("page[number]", strconv.Itoa(pageNumber))
	params.Add("page[size]", strconv.Itoa(pageSize))
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return Accounts{}, nil
	}

	req.Header.Set("Accept", "vnd.api+json")

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return Accounts{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return Accounts{}, &ResourceNotExistsError{baseURL.String()}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Accounts{}, err
	}

	var a Accounts
	if err := json.Unmarshal(body, &a); err != nil {
		return Accounts{}, err
	}

	return a, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string, version int) error {
	accountID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}

	baseURL.Path = path.Join(baseURL.Path, accountID.String())
	params := url.Values{}
	params.Add("version", strconv.Itoa(version))
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodDelete, baseURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return &InvalidVersionError{version}
	}

	return nil
}

func NewAccount(opt *Options) (*Account, error) {
	if err := opt.validate(); err != nil {
		return nil, err
	}

	attributes := &Attributes{}

	for _, attr := range opt.Attributes {
		attr(attributes)
	}

	if err := attributes.validate(); err != nil {
		return nil, err
	}

	account := &Account{
		Data: &Data{
			Attributes: attributes,
			Details: Details{
				Type:           opt.Type,
				ID:             opt.ID,
				OrganisationID: opt.OrganisationID,
			},
		},
	}

	return account, nil
}

func NewClient(client *http.Client, baseURL ...string) *Client {
	cl := &Client{}

	if client != nil {
		cl.Client = client
	} else {
		cl.Client = &http.Client{}
	}

	u := os.Getenv("FORM3_BASE_URL")
	if u != "" {
		cl.BaseURL = u
	} else {
		if len(baseURL) != 0 && baseURL[0] != "" {
			cl.BaseURL = baseURL[0]
		} else {
			cl.BaseURL = DefaultBaseURL
		}
	}

	return cl
}

func WithAttrCountry(country string) Attribute {
	return func(a *Attributes) {
		a.Country = country
	}
}

func WithAttrBankID(id string) Attribute {
	return func(a *Attributes) {
		a.BankID = id
	}
}

func WithAttrBankIDCode(code string) Attribute {
	return func(a *Attributes) {
		a.BankIDCode = code
	}
}

func WithAttrBIC(bic string) Attribute {
	return func(a *Attributes) {
		a.BIC = bic
	}
}

func WithAttrAccountNumber(number string) Attribute {
	return func(a *Attributes) {
		a.AccountNumber = number
	}
}

func WithAttrBaseCurrency(currency string) Attribute {
	return func(a *Attributes) {
		a.BaseCurrency = currency
	}
}

func WithAttrJointAccount(joint bool) Attribute {
	return func(a *Attributes) {
		a.JointAccount = joint
	}
}

func WithAttrFirstName(name string) Attribute {
	return func(a *Attributes) {
		a.FirstName = name
	}
}

func WithAttrAlternativeBankAccountNames(names ...string) Attribute {
	return func(a *Attributes) {
		a.AlternativeBankAccountNames = names
	}
}

func WithAttrAccountMatchingOptOut(optOut bool) Attribute {
	return func(a *Attributes) {
		a.AccountMatchingOptOut = optOut
	}
}

func WithAttrCustomerID(id string) Attribute {
	return func(a *Attributes) {
		a.CustomerID = id
	}
}
