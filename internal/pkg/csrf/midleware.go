package csrf

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func CSRFCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		csrfToken := r.URL.Query().Get("csrf_token")
		sessionCSRFToken, ok := session.Values["csrf_token"].(string)

		if !ok || csrfToken == "" || csrfToken != sessionCSRFToken {
			http.Error(w, "Forbidden - invalid CSRF token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupCSRFMiddleware(r *mux.Router) *mux.Router {
	csrfHash := "sdfjkldsjfsdofipoaeiwjfijcsdjfco"
	csrfMiddleware := csrf.Protect([]byte(csrfHash), csrf.Secure(false))
	r.Use(csrfMiddleware)
	return r
}
