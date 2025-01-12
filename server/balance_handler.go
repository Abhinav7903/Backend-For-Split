package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhinav7903/split/factory"
)

func (s *Server) handlerAddBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON payload
		var balance factory.Balance
		err := json.NewDecoder(r.Body).Decode(&balance)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the input
		if balance.UserID <= 0 || balance.OwedAmount < 0 || balance.LentAmount < 0 {
			http.Error(w, "Invalid balance data", http.StatusBadRequest)
			return
		}

		// Add the balance
		balanceID, err := s.balance.CreateBalance(&balance)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to add balance"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the balance ID
		s.respond(w, ResponseMsg{Message: "Balance added successfully", Data: balanceID}, http.StatusOK, nil)
	}
}

func (s *Server) handlerGetBalanceByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract balance ID from the URL
		balanceID := r.URL.Query().Get("balance_id")
		if balanceID == "" {
			http.Error(w, "Balance ID is required", http.StatusBadRequest)
			return
		}

		// Convert balance ID to int
		id, err := strconv.Atoi(balanceID)
		if err != nil {
			http.Error(w, "Invalid balance ID", http.StatusBadRequest)
			return
		}

		// Fetch the balance by ID
		balance, err := s.balance.GetBalanceByID(id)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to fetch balance"}, http.StatusInternalServerError, nil)
			return
		}

		// If no balance found
		if balance == nil {
			http.Error(w, "Balance not found", http.StatusNotFound)
			return
		}

		// Respond with the balance data
		s.respond(w, ResponseMsg{Message: "Balance fetched successfully", Data: balance}, http.StatusOK, nil)
	}
}

func (s *Server) handlerUpdateBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract balance ID from the URL
		balanceID := r.URL.Query().Get("balance_id")
		if balanceID == "" {
			http.Error(w, "Balance ID is required", http.StatusBadRequest)
			return
		}

		// Convert balance ID to int
		id, err := strconv.Atoi(balanceID)
		if err != nil {
			http.Error(w, "Invalid balance ID", http.StatusBadRequest)
			return
		}

		// Parse the JSON payload
		var balance factory.Balance
		err = json.NewDecoder(r.Body).Decode(&balance)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the input
		if balance.OwedAmount < 0 || balance.LentAmount < 0 {
			http.Error(w, "Invalid balance data", http.StatusBadRequest)
			return
		}

		// Update the balance
		err = s.balance.UpdateBalanceAmounts(id, &balance.OwedAmount, &balance.LentAmount)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to update balance"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "Balance updated successfully"}, http.StatusOK, nil)
	}
}
func (s *Server) handlerDeleteBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract balance ID from the URL
		balanceID := r.URL.Query().Get("balance_id")
		if balanceID == "" {
			http.Error(w, "Balance ID is required", http.StatusBadRequest)
			return
		}

		// Convert balance ID to int
		id, err := strconv.Atoi(balanceID)
		if err != nil {
			http.Error(w, "Invalid balance ID", http.StatusBadRequest)
			return
		}

		// Delete the balance
		err = s.balance.DeleteBalance(id)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to delete balance"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "Balance deleted successfully"}, http.StatusOK, nil)
	}
}
