package basicauth

import "net/http"

// AuthSource is any struct that can modify headers for auth.
type AuthSource interface {
	// SetAuthHeader modifies the supplied header to include basic auth.
	SetAuthHeader(r *http.Request)
}

// Auth contains the username and password used for basic auth
type Auth struct {
	user string
	pw   string
}

// NewAuth returns a new AuthSource compatible struct (Auth)
func NewAuth(user, pass string) *Auth {
	return &Auth{
		user: user,
		pw:   pass,
	}
}

// User returns the username used for the authentication
func (a *Auth) User() string {
	return a.user
}

// SetAuthHeader sets the Authorization header to r using the auth
// details in a
//
// This method is unnecessary when using Transport or an HTTP Client
// returned by this package.
func (a *Auth) SetAuthHeader(r *http.Request) {
	r.SetBasicAuth(a.user, a.pw)
}
