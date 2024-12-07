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

	s.router.HandleFunc(
		"/getuser",
		(s.handleGetUser()),
	).Methods(http.MethodGet, http.MethodOptions)

	s.router.HandleFunc(
		"/updateuser",
		(s.handleUpdateUser()),
	).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc(
		"/deleteuser",
		(s.handleDeleteUser()),
	).Methods(http.MethodDelete, http.MethodOptions)

	s.router.HandleFunc(
		"/getallusers",
		(s.handleGetAllUsers()),
	).Methods(http.MethodGet, http.MethodOptions)

	s.router.HandleFunc(
		"/email-exists",
		(s.EmailExists()),
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
