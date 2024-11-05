package middlewares

import (
	"net/http"
	"todo-list/utils"
)

func AdminValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO - Change this to a more secure way of validating the user (this is only temporary of course)
		if r.Header.Get("Authorization") != "Bearer admin" {
			utils.EncodeJSON[any](w, r, http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	})

}
