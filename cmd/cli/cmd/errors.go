package datum

import "errors"

var (
	// ErrOrgNameRequired is returned when no organization name is provided when creating a new organization
	ErrOrgNameRequired = errors.New("organization name is required")

	// ErrOrgIDRequired is returned when no organization id is provided when deleting or updating an organization
	ErrOrgIDRequired = errors.New("organization id is required")
)