package csrf

import (
	"errors"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("s3cr3t-k3y"))

func SetSession(r *http.Request, w http.ResponseWriter) error {
	session, err := GetSession(r)
	if err != nil {
		return errors.New("problem on setting session")
	}

	csrfToken := csrf.Token(r)
	session.Values["csrf_token"] = csrfToken
	err = session.Save(r, w)
	if err != nil {
		return errors.New("problem on setting session")
	}

	return nil
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session-name")
}
