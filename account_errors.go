package accountapi

import "fmt"

// ResourceNotExistsError is returned if URL is invalid and the resource it points to does not exist.
type ResourceNotExistsError struct {
	Resource string
}

func (e *ResourceNotExistsError) Error() string {
	return fmt.Sprintf("resource '%s' does not exist", e.Resource)
}

// InvalidVersionError is returned if the version arrgument that's passsed to DeleteAccount is not 0.
type InvalidVersionError struct {
	Version int
}

func (e *InvalidVersionError) Error() string {
	return fmt.Sprintf("invalid version %d", e.Version)
}

// RecordNotExistsError is returned if trying to fetch and account that does not exist.
type RecordNotExistsError struct {
	Record string
}

func (e *RecordNotExistsError) Error() string {
	return fmt.Sprintf("record '%s' does not exist", e.Record)
}

// DuplicateAccountError is returned if an account with that id already exists.
type DuplicateAccountError struct {
	ID string
}

func (e *DuplicateAccountError) Error() string {
	return fmt.Sprintf("duplicate account '%s'", e.ID)
}
