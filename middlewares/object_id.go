package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/horus/res"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidateObjectID returns a middleware that checks if the `id` param is a valid mongo.ObjectID
func ValidateObjectID(id string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("id: %s\n", chi.URLParam(r, id))
			if !primitive.IsValidObjectID(chi.URLParam(r, id)) {
				err := errors.New("invalid object id param")
				render.Render(w, r, res.Err(err, http.StatusBadRequest))
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
