package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// Uploads is the resolver for the uploads field.
func (r *userResolver) Uploads(ctx context.Context, obj *model.User) ([]*model.AnyUpload, error) {
	panic(fmt.Errorf("not implemented: Uploads - uploads"))
}

// Properties is the resolver for the properties field.
func (r *userResolver) Properties(ctx context.Context, obj *model.User) ([]*model.Property, error) {
	userProperties, err := ctx.Value("propertyService").(*services.PropertyServices).PropertiesCreatedBy(obj.ID)
	if err != nil {
		return nil, err
	}
	return userProperties, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
