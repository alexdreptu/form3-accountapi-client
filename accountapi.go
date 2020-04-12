// Package accountapi provides a client for the fake Form3 Finacial Cloud RESTFul API.
// It implements only create, fetch, list and delete.
package accountapi

import (
	"github.com/google/uuid"
)

type AccountAttributes struct{}

type AccountAttribute func(*AccountAttributes)

type Account struct{}

type AccountClient struct{}

func (ac *AccountClient) Create(a Account) (Account, error) {
	return Account{}, nil
}

func (ac *AccountClient) Fetch(id uuid.UUID) (Account, error) {
	return Account{}, nil
}

func (ac *AccountClient) List() ([]Account, error) {
	return []Account{}, nil
}

func (ac *AccountClient) Delete(id uuid.UUID) error {
	return nil
}

func NewAccountClient() *AccountClient {
	return &AccountClient{}
}
