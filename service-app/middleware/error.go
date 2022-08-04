package middleware

import (
	"context"
	"errors"
	"net/http"
	"service-app/web"
)

func (m *Mid) Error(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		v, ok := ctx.Value(web.KeyValue).(*web.Values)
		if !ok {
			return errors.New("value not found in the context")
		}
		err := next(ctx, w, r)
		if err != nil {
			// Log the error.
			m.Log.Printf("%s : ERROR     : %v", v.TraceID, err)

			err = web.RespondError(ctx, w, err)
			if err != nil {
				return err
			}

		}
		return nil
	}

}
