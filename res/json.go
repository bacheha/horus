package res

import "net/http"

type JSON map[string]interface{}

// Render implements the chi.Render interface for payload responses.
func (j *JSON) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
