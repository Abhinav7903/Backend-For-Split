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

	// Add a group
	s.router.HandleFunc(
		"/add-group",
		s.handlerAddGroup(),
	).Methods(http.MethodPost, http.MethodOptions)

	// Get a single group by ID
	s.router.HandleFunc(
		"/get-group",
		s.handlerGetGroup(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Get all groups
	s.router.HandleFunc(
		"/get-all-groups",
		s.handlerGetAllGroups(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Update a group
	s.router.HandleFunc(
		"/update-group",
		s.handlerUpdateGroup(),
	).Methods(http.MethodPut, http.MethodOptions)

	// Delete a group by ID
	s.router.HandleFunc(
		"/delete-group",
		s.handlerDeleteGroup(),
	).Methods(http.MethodDelete, http.MethodOptions)

	// Check if a group exists
	s.router.HandleFunc(
		"/group-exists",
		s.handlerGroupExists(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Add a group member
	s.router.HandleFunc(
		"/add-group-member",
		s.handleAddGroupMember(),
	).Methods(http.MethodPost, http.MethodOptions)

	// Get a single group member by ID
	s.router.HandleFunc(
		"/get-group-member",
		s.handleGetGroupMember(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Get all group members in a group
	s.router.HandleFunc(
		"/get-group-members-by-group-id",
		s.handleGetGroupMembersByGroupID(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Update a group member...

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
