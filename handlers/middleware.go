package handlers

import (
	"context"
	"go-microservice/data"
	"net/http"
)

func (t *Teams) MiddlewareTeamValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		team := &data.Team{}
		err := data.FromJSON(team, r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing JSON")
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		errs := t.v.Validate(team)
		if len(errs) != 0 {
			t.l.Println("[ERROR] validating team", errs)
			// Return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// Add the team to the context
		ctx := context.WithValue(r.Context(), KeyTeam{}, team)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
