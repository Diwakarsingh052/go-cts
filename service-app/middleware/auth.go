package middleware

import (
	"context"
	"errors"
	"net/http"
	"service-app/auth"
	"service-app/web"
	"strings"
)

func (m *Mid) Authenticate(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		// fetching data from Authorization header
		authHeader := r.Header.Get("Authorization")

		//token format :- bearer <token> // we trying to separate both strings with space
		parts := strings.Split(authHeader, " ")

		// checking if we have got exactly to values in auth header and first word is bearer
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: Bearer <token>")
			return web.NewRequestError(err, http.StatusUnauthorized)
		}

		// sending the token to the auth layer for verification
		claims, err := m.A.ValidateToken(parts[1])
		if err != nil {
			return web.NewRequestError(err, http.StatusUnauthorized)
		}
		// putting the token in the context so we can see the values in the claims struct in the request life cycle //
		// specifically we will look for the subject field as it stores unique user id which will helpful to identify for whom this token was genrated
		ctx = context.WithValue(ctx, auth.Key, claims)

		return next(ctx, w, r)

	}
}

func (m *Mid) HasRole(next web.HandlerFunc, roles ...string) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		claims, ok := ctx.Value(auth.Key).(auth.Claims)
		if !ok {
			return errors.New("claims missing from context: HasRole called without/before Authenticate")
		}

		ok = claims.HasRoles(roles...)
		if !ok {
			return web.NewRequestError(errors.New("you are not authorized for that action"), http.StatusForbidden)
		}

		return next(ctx, w, r)

	}

}
