package client

import (
	"fmt"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
)

func NewGraphQLClient() (*api.GraphQLClient, error) {
	opts := api.ClientOptions{
		EnableCache: true,
		Timeout:     5 * time.Second,
	}
	client, err := api.NewGraphQLClient(opts)
	if err != nil {
		return nil, fmt.Errorf("while creating GraphQL client: %w", err)
	}
	return client, nil
}
