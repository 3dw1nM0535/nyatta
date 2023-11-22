package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// User is the resolver for the user field.
func (r *tenantResolver) User(ctx context.Context, obj *model.Tenant) (*model.User, error) {
	user, err := ctx.Value("userService").(*services.UserServices).GetUser(obj.UserID)
	if err != nil {
		return nil, err
	}

	return user, err
}

// PropertyUnit is the resolver for the propertyUnit field.
func (r *tenantResolver) PropertyUnit(ctx context.Context, obj *model.Tenant) (*model.PropertyUnit, error) {
	unit, err := ctx.Value("unitService").(*services.UnitServices).GetPropertyUnit(obj.PropertyUnitID)
	if err != nil {
		return nil, err
	}

	return unit, err
}

// Tenant returns generated.TenantResolver implementation.
func (r *Resolver) Tenant() generated.TenantResolver { return &tenantResolver{r} }

type tenantResolver struct{ *Resolver }
