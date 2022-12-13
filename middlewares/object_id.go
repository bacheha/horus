package middlewares

import (
	e "errors"
	"net/http"

	"github.com/bacheha/horus/res"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectID returns a middleware that checks if the `id` param  is a valid mongo.ObjectID
func ObjectID(id string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !primitive.IsValidObjectID(chi.URLParam(r, id)) {
				err := e.New("invalid object id")
				render.Render(w, r, res.Err(err, http.StatusBadRequest))
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
