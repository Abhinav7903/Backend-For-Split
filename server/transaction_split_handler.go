package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhinav7903/split/factory"
	"github.com/gorilla/mux"
)

func (s *Server) CreateTransactionSplitHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request body into a TransactionSplit object
		var request factory.TransactionSplit
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.respond(w, ResponseMsg{
				Message: "failed",
				Data:    "Invalid payload",
			}, http.StatusBadRequest, nil)
			return
		}

		// Call the repository to create the transaction split
		if err := s.transactionsplit.CreateTransactionSplit(&request); err != nil {
			s.respond(w, ResponseMsg{
				Message: "failed",
				Data:    "Error creating transaction split: " + err.Error(),
			}, http.StatusInternalServerError, nil)
			return
		}

		// Return the created object and a 201 Created status
		s.respond(w, ResponseMsg{
			Message: "success",
			Data:    request,
		}, http.StatusCreated, nil)
	}
}

func (s *Server) GetTransactionSplitsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract transaction ID from the URL
		vars := mux.Vars(r)
		transactionID, err := strconv.Atoi(vars["transaction_id"])
		if err != nil {
			s.respond(w, ResponseMsg{
				Message: "failed",
				Data:    "Invalid transaction ID",
			}, http.StatusBadRequest, nil)
			return
		}

		// Call the repository to get transaction splits
		transactionSplits, err := s.transactionsplit.GetTransactionSplits(transactionID)
		if err != nil {
			s.respond(w, ResponseMsg{
				Message: "failed",
				Data:    "Error getting transaction splits: " + err.Error(),
			}, http.StatusInternalServerError, nil)
			return
		}

		// Return the list of transaction splits
		s.respond(w, ResponseMsg{
			Message: "success",
			Data:    transactionSplits,
		}, http.StatusOK, nil)
	}
}

