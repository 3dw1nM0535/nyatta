package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// Avatar is the resolver for the avatar field.
func (r *caretakerResolver) Avatar(ctx context.Context, obj *model.Caretaker) (*model.AnyUpload, error) {
	panic(fmt.Errorf("not implemented: Avatar - avatar"))
}

// Properties is the resolver for the properties field.
func (r *caretakerResolver) Properties(ctx context.Context, obj *model.Caretaker) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented: Properties - properties"))
}

// Caretaker returns generated.CaretakerResolver implementation.
func (r *Resolver) Caretaker() generated.CaretakerResolver { return &caretakerResolver{r} }

type caretakerResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *caretakerResolver) Uploads(ctx context.Context, obj *model.Caretaker) ([]*model.AnyUpload, error) {
	panic(fmt.Errorf("not implemented: Uploads - uploads"))
}
