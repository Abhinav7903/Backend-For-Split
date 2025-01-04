package server

import (
	"encoding/json"
	"net/http"

	"github.com/Abhinav7903/split/factory"
)

// handleCreatePaymentMethod handles the creation of a new payment method
func (s *Server) handleCreatePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body
		var paymentMethod factory.PaymentMethod
		if err := json.NewDecoder(r.Body).Decode(&paymentMethod); err != nil {
			// Log error and respond with a bad request message
			s.logger.Error("Error decoding request body: %v", err)
			s.respond(w, ResponseMsg{Message: "invalid request body", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.logger.Info("Attempting to create payment method for email: %s", paymentMethod.Email)

		// Create the payment method in the database
		if err := s.payment.CreatePaymentMethod(&paymentMethod); err != nil {
			// Log error if creation fails
			s.logger.Error("Error creating payment method for email %s: %v", paymentMethod.Email, err)
			s.respond(w, ResponseMsg{Message: "failed", Data: err.Error()}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success
		s.logger.Info("Payment method created successfully for email: %s", paymentMethod.Email)
		s.respond(w, ResponseMsg{Message: "success", Data: "Payment method added successfully"}, http.StatusCreated, nil)
	}
}

// handleGetPaymentMethods handles retrieving payment methods for a user
func (s *Server) handleGetPaymentMethods() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract email from query parameters
		email := r.URL.Query().Get("email")
		if email == "" {
			s.logger.Error("email parameter is missing")
			s.respond(w, ResponseMsg{Message: "failed", Data: "email parameter is required"}, http.StatusBadRequest, nil)
			return
		}

		// Retrieve payment methods from the database
		paymentMethods, err := s.payment.GetPaymentMethods(email)
		if err != nil {
			s.logger.Error("error fetching payment methods", "email:", email, err)
			s.respond(w, ResponseMsg{Message: "failed", Data: "could not retrieve payment methods"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the retrieved payment methods
		s.logger.Info("retrieved payment methods for email:", email)
		s.respond(w, ResponseMsg{Message: "success", Data: paymentMethods}, http.StatusOK, nil)
	}
}

// handleUpdatePaymentMethod updates an existing payment method
func (s *Server) handleUpdatePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paymentMethod factory.PaymentMethod

		// Decode the request body
		if err := json.NewDecoder(r.Body).Decode(&paymentMethod); err != nil {
			s.logger.Error("error decoding request body: %v", err)
			s.respond(w, ResponseMsg{Message: "failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		// Validate the payment method data (only email and payment_type are required)
		if paymentMethod.Email == "" {
			s.logger.Error("validation error: email is required")
			s.respond(w, ResponseMsg{Message: "failed", Data: "email is required"}, http.StatusBadRequest, nil)
			return
		}
		if paymentMethod.PaymentType == "" {
			s.logger.Error("validation error: payment_type is required")
			s.respond(w, ResponseMsg{Message: "failed", Data: "payment_type is required"}, http.StatusBadRequest, nil)
			return
		}

		// Log the email and payment type being used for update
		s.logger.Info("attempting to update payment method for email: %s with payment_type: %s", paymentMethod.Email, paymentMethod.PaymentType)

		// Update the payment method in the database
		if err := s.payment.UpdatePaymentMethod(&paymentMethod); err != nil {
			s.logger.Error("error updating payment method: %v", err)
			s.respond(w, ResponseMsg{Message: "failed", Data: err.Error()}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success
		s.logger.Info("payment method updated successfully for payment_id:", "id=%d", paymentMethod.PaymentID)
		s.respond(w, ResponseMsg{Message: "success", Data: "Payment method updated successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleDeletePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Struct to capture incoming request data
		var request struct {
			Email       string `json:"email"`
			PaymentType string `json:"payment_type"`
		}

		// Decode the request body into the struct
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			s.logger.Error("error decoding request body: %v", err)
			s.respond(w, ResponseMsg{Message: "failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		// Validate the inputs
		if request.Email == "" || request.PaymentType == "" {
			s.logger.Error("validation error: email or payment_type is invalid")
			s.respond(w, ResponseMsg{Message: "failed", Data: "Invalid email or payment_type"}, http.StatusBadRequest, nil)
			return
		}

		// Call the method to delete the payment method
		if err := s.payment.DeletePaymentMethod(request.PaymentType, request.Email); err != nil {
			s.logger.Error("error deleting payment method: %v", err)
			s.respond(w, ResponseMsg{Message: "failed", Data: err.Error()}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success
		s.logger.Info("payment method deleted successfully for email:", "email=%s", request.Email)
		s.respond(w, ResponseMsg{Message: "success", Data: "Payment method deleted successfully"}, http.StatusOK, nil)
	}
}
