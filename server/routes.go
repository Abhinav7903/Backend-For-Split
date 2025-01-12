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

	//remove group member by creator
	s.router.HandleFunc(
		"/remove-group-member",
		s.handleRemoveUserFromGroupByCreator(),
	).Methods(http.MethodDelete, http.MethodOptions)

	//self remove
	s.router.HandleFunc(
		"/remove-group-member-self",
		s.handleRemoveUserSelfFromGroup(),
	).Methods(http.MethodDelete, http.MethodOptions)

	// Create transaction
	s.router.HandleFunc(
		"/create-transaction",
		s.handleCreateTransaction(),
	).Methods(http.MethodPost, http.MethodOptions)

	// Get transaction by ID
	s.router.HandleFunc(
		"/get-transaction",
		s.handleGetTransactionByID(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Get transactions by lender ID
	s.router.HandleFunc(
		"/get-transactions-by-lender",
		s.handleGetTransactionsByLenderID(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Get transactions by borrower ID
	s.router.HandleFunc(
		"/get-transactions-by-borrower",
		s.handleGetTransactionsByBorrowerID(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Update transaction status
	s.router.HandleFunc(
		"/update-transaction-status",
		s.handleUpdateTransactionStatus(),
	).Methods(http.MethodPut, http.MethodOptions)

	// Delete transaction
	s.router.HandleFunc(
		"/delete-transaction",
		s.handleDeleteTransaction(),
	).Methods(http.MethodDelete, http.MethodOptions)

	// Search transactions
	s.router.HandleFunc(
		"/search-transactions",
		s.handleSearchTransactions(),
	).Methods(http.MethodPost, http.MethodOptions)

	// Transaction Split
	s.router.HandleFunc(
		"/create-transaction-split",
		s.CreateTransactionSplitHandler(),
	).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc(
		"/get-transaction-splits/{transaction_id}",
		s.GetTransactionSplitsHandler(),
	).Methods(http.MethodGet, http.MethodOptions)

	// Payment Methods
	s.router.HandleFunc(
		"/create-payment-method",
		s.handleCreatePaymentMethod(),
	).Methods(http.MethodPost, http.MethodOptions)

	s.router.HandleFunc(
		"/get-payment-methods",
		s.handleGetPaymentMethods(),
	).Methods(http.MethodGet, http.MethodOptions)

	s.router.HandleFunc(
		"/update-payment-method",
		s.handleUpdatePaymentMethod(),
	).Methods(http.MethodPut, http.MethodOptions)

	s.router.HandleFunc(
		"/delete-payment-method",
		s.handleDeletePaymentMethod(),
	).Methods(http.MethodDelete, http.MethodOptions)

	s.router.HandleFunc("/balance", s.handlerAddBalance()).Methods(http.MethodPost)
	s.router.HandleFunc("/balance", s.handlerGetBalanceByID()).Methods(http.MethodGet)
	s.router.HandleFunc("/balance", s.handlerUpdateBalance()).Methods(http.MethodPut)
	s.router.HandleFunc("/balance", s.handlerDeleteBalance()).Methods(http.MethodDelete)

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
