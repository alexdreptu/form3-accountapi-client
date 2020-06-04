package accountapi

import "fmt"

// ResourceNotExistsError is returned if URL is invalid and the resource it points to does not exist.
type ResourceNotExistsError struct {
	Resource string
}

func (e *ResourceNotExistsError) Error() string {
	return fmt.Sprintf("resource '%s' does not exist", e.Resource)
}

// DuplicateAccountError is returned if an account with that id already exists.
type DuplicateAccountError struct {
	ID string
}

func (e *DuplicateAccountError) Error() string {
	return fmt.Sprintf("duplicate account '%s'", e.ID)
}
