package basicauth

import (
	"context"
	"net/http"
)

// NewClient creates an *http.Client from a Context and AuthSource.
//
// As a special case, if src is nil, a non-auth client is returned.
func NewClient(ctx context.Context, src AuthSource) *http.Client {
	if src == nil {
		return http.DefaultClient
	}
	return &http.Client{
		Transport: &Transport{
			Base: http.DefaultTransport,
			Auth: src,
		},
	}
}
