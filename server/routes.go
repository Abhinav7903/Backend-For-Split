package server

import "net/http"

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/ping", s.HandlePong()).Methods(http.MethodGet)

	s.router.HandleFunc(
		"/signup",
		(s.handleSignUp()),
	).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc(
		"/verify",
		(s.handleVerify()),
	).Methods(http.MethodGet, http.MethodOptions)

}
func (s *Server) HandlePong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(
			w,
			"pong",
			http.StatusOK,
			nil,
		)
	}
}
