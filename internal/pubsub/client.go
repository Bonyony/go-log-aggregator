package mypubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID, option.WithServiceAccountFile("service-account.json"))
	if err != nil {
		return nil, err
	}
	return client, nil
}
