package httpsrv

import (
	"fmt"
	"goapp/internal/pkg/csrf"
	"net/http"
	"runtime/debug"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HFunc      http.Handler
	Queries    []string
	Middleware func(http.Handler) http.Handler
}

func (s *Server) myRoutes() []Route {
	return []Route{
		{
			Name:       "health",
			Method:     "GET",
			Pattern:    "/goapp/health",
			HFunc:      s.handlerWrapper(s.handlerHealth),
			Middleware: nil,
		},
		{
			Name:       "websocket",
			Method:     "GET",
			Pattern:    "/goapp/ws",
			HFunc:      s.handlerWrapper(s.handlerWebSocket),
			Middleware: csrf.CSRFCheckMiddleware,
		},
		{
			Name:       "home",
			Method:     "GET",
			Pattern:    "/goapp",
			HFunc:      s.handlerWrapper(s.handlerHome),
			Middleware: nil,
		},
	}
}

func (s *Server) handlerWrapper(handlerFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				s.error(w, http.StatusInternalServerError, fmt.Errorf("%v\n%v", r, string(debug.Stack())))
			}
		}()
		handlerFunc(w, r)
	})
}
