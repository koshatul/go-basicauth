package basicauth

import (
	"net/http"
)

// NewClient creates an *http.Client from a Context and AuthSource.
//
// As a special case, if src is nil, a non-auth client is returned.
func NewClient(src AuthSource) *http.Client {
	if src == nil {
		return http.DefaultClient
	}
	return &http.Client{
		Transport: &Transport{
			Auth: src,
		},
	}
}
