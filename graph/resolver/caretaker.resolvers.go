package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// Avatar is the resolver for the avatar field.
func (r *caretakerResolver) Avatar(ctx context.Context, obj *model.Caretaker) (*model.AnyUpload, error) {
	upload, err := ctx.Value("propertyService").(services.PropertyService).GetCaretakerAvatar(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return upload, nil
}

// Properties is the resolver for the properties field.
func (r *caretakerResolver) Properties(ctx context.Context, obj *model.Caretaker) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented: Properties - properties"))
}

// Caretaker returns generated.CaretakerResolver implementation.
func (r *Resolver) Caretaker() generated.CaretakerResolver { return &caretakerResolver{r} }

type caretakerResolver struct{ *Resolver }
