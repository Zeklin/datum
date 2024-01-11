package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateOrgMembership is the resolver for the createOrgMembership field.
func (r *mutationResolver) CreateOrgMembership(ctx context.Context, input generated.CreateOrgMembershipInput) (*OrgMembershipCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateOrgMembership - createOrgMembership"))
}

// UpdateOrgMembership is the resolver for the updateOrgMembership field.
func (r *mutationResolver) UpdateOrgMembership(ctx context.Context, id string, input generated.UpdateOrgMembershipInput) (*OrgMembershipUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateOrgMembership - updateOrgMembership"))
}

// DeleteOrgMembership is the resolver for the deleteOrgMembership field.
func (r *mutationResolver) DeleteOrgMembership(ctx context.Context, id string) (*OrgMembershipDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteOrgMembership - deleteOrgMembership"))
}

// OrgMembership is the resolver for the orgMembership field.
func (r *queryResolver) OrgMembership(ctx context.Context, id string) (*generated.OrgMembership, error) {
	panic(fmt.Errorf("not implemented: OrgMembership - orgMembership"))
}