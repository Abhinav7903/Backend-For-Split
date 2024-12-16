package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhinav7903/split/factory"
)

// handleCreateTransaction handles the creation of a new transaction.
func (s *Server) handleCreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction factory.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid request body"}, http.StatusBadRequest, nil)
			return
		}

		transactionID, err := s.transaction.CreateTransaction(&transaction)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to create transaction"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transaction created", Data: transactionID}, http.StatusCreated, nil)
	}
}

// handleGetTransactionByID retrieves a transaction by its ID.
func (s *Server) handleGetTransactionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || transactionID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid transaction ID"}, http.StatusBadRequest, nil)
			return
		}

		transaction, err := s.transaction.GetTransactionByID(transactionID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to retrieve transaction"}, http.StatusInternalServerError, nil)
			return
		}
		if transaction == nil {
			s.respond(w, ResponseMsg{Message: "Transaction not found"}, http.StatusNotFound, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transaction retrieved", Data: transaction}, http.StatusOK, nil)
	}
}

// handleGetTransactionsByLenderID retrieves transactions by lender ID.
func (s *Server) handleGetTransactionsByLenderID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lenderID, err := strconv.Atoi(r.URL.Query().Get("lender_id"))
		if err != nil || lenderID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid lender ID"}, http.StatusBadRequest, nil)
			return
		}

		transactions, err := s.transaction.GetTransactionsByLenderID(lenderID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to retrieve transactions"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transactions retrieved", Data: transactions}, http.StatusOK, nil)
	}
}

// handleGetTransactionsByBorrowerID retrieves transactions by borrower ID.
func (s *Server) handleGetTransactionsByBorrowerID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		borrowerID, err := strconv.Atoi(r.URL.Query().Get("borrower_id"))
		if err != nil || borrowerID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid borrower ID"}, http.StatusBadRequest, nil)
			return
		}

		transactions, err := s.transaction.GetTransactionsByBorrowerID(borrowerID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to retrieve transactions"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transactions retrieved", Data: transactions}, http.StatusOK, nil)
	}
}

// handleUpdateTransactionStatus updates the status of a transaction.
func (s *Server) handleUpdateTransactionStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || transactionID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid transaction ID"}, http.StatusBadRequest, nil)
			return
		}

		var payload struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid request body"}, http.StatusBadRequest, nil)
			return
		}

		err = s.transaction.UpdateTransactionStatus(transactionID, payload.Status)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to update transaction status"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transaction status updated"}, http.StatusOK, nil)
	}
}

// handleDeleteTransaction deletes a transaction by its ID.
func (s *Server) handleDeleteTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || transactionID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid transaction ID"}, http.StatusBadRequest, nil)
			return
		}

		err = s.transaction.DeleteTransaction(transactionID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to delete transaction"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transaction deleted"}, http.StatusOK, nil)
	}
}

// handleSearchTransactions searches for transactions based on filters.
func (s *Server) handleSearchTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filters factory.TransactionFilters
		if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid request body"}, http.StatusBadRequest, nil)
			return
		}

		transactions, err := s.transaction.SearchTransactions(filters)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to search transactions"}, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Transactions retrieved", Data: transactions}, http.StatusOK, nil)
	}
}
