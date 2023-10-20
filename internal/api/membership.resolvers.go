package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/google/uuid"
)

// CreateMembership is the resolver for the createMembership field.
func (r *mutationResolver) CreateMembership(ctx context.Context, input generated.CreateMembershipInput) (*MembershipCreatePayload, error) {
	// TODO - add permissions checks
	mem, err := r.client.Membership.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create membership", "error", err)
		return nil, ErrInternalServerError
	}

	return &MembershipCreatePayload{Membership: mem}, nil
}

// UpdateMembership is the resolver for the updateMembership field.
func (r *mutationResolver) UpdateMembership(ctx context.Context, id uuid.UUID, input generated.UpdateMembershipInput) (*MembershipUpdatePayload, error) {
	// TODO - add permissions checks

	mem, err := r.client.Membership.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get membership", "error", err)
		return nil, ErrInternalServerError
	}

	mem, err = mem.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update membership", "error", err)
		return nil, ErrInternalServerError
	}

	return &MembershipUpdatePayload{Membership: mem}, nil
}

// DeleteMembership is the resolver for the deleteMembership field.
func (r *mutationResolver) DeleteMembership(ctx context.Context, id uuid.UUID) (*MembershipDeletePayload, error) {
	// TODO - add permissions checks

	if err := r.client.Membership.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete membership", "error", err)
		return nil, err
	}

	return &MembershipDeletePayload{DeletedID: id}, nil
}

// Membership is the resolver for the membership field.
func (r *queryResolver) Membership(ctx context.Context, id uuid.UUID) (*generated.Membership, error) {
	// TODO - add permissions checks

	mem, err := r.client.Membership.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get membership", "error", err)
		return nil, ErrInternalServerError
	}

	return mem, nil
}
