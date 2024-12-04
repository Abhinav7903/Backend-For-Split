package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhinav7903/split/factory"
)

func (s *Server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("sign up request")
		var user factory.User

		// Decode the request body into a user struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			s.logger.Error("failed to decode request body", "error", err)
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
			return
		}

		// Add user to the database
		err = s.user.AddUser(user)
		if err != nil {
			s.logger.Error("failed to add user", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to create user"))
			return
		}

		// Generate an email hash
		hash, err := s.sessmanager.StoreEmailHash(user.Email)
		if err != nil {
			s.logger.Error("failed to store email hash", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to process email verification"))
			return
		}

		// Send verification email to the user
		err = s.mail.SendMail(
			user.Email,
			"Verify your email",
			"Click the link to verify your email: http://localhost:8080/verify?ehash="+hash,
		)
		if err != nil {
			s.logger.Error("failed to send verification email", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("verification email failed"))
			return
		}

		// Notify admin about the new user
		err = s.mail.SendMail(
			"abhinavashish4@gmail.com",
			"New user signed up",
			"New user signed up with email: "+user.Email,
		)
		if err != nil {
			s.logger.Error("failed to notify admin about the new user", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("admin notification failed"))
			return
		}

		// Respond with success
		s.respond(w, ResponseMsg{Message: "User added successfully. Verification emails sent."}, http.StatusOK, nil)
	}
}

func (s *Server) handleVerify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("verify request")

		// Extract the email hash from the query parameters
		hash := r.URL.Query().Get("ehash")
		if hash == "" {
			s.logger.Error("missing email hash in request")
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid verification link"))
			return
		}

		// Retrieve email associated with the hash
		email, err := s.sessmanager.GetEmailFromHash(hash)
		if err != nil {
			s.logger.Error("failed to retrieve email from hash", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("email verification failed"))
			return
		}

		// Mark the email as verified in the database
		if err := s.user.VerifyEmail(email); err != nil {
			s.logger.Error("failed to verify email", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("email verification update failed"))
			return
		}

		// Redirect to success page
		http.Redirect(w, r, "http://localhost:8080/verify-success", http.StatusSeeOther)
	}
}
