package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	generated1 "gitlab.com/pragmaticreviews/gqlgen_tutorial/graph/generated"
	model1 "gitlab.com/pragmaticreviews/gqlgen_tutorial/graph/model"
	"gitlab.com/pragmaticreviews/graphql-go/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model1.NewVideo) (*model1.Video, error) {
	video := &model.Video{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Title:  input.Title,
		Author: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.videos = append(r.videos, video)
	return video, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model1.Video, error) {
	return r.videos, nil
}

func (r *subscriptionResolver) VideoAdded(ctx context.Context, repoFullName string) (<-chan *model1.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
