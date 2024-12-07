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
			"New user signed up with email: "+user.Email+" and name: "+user.Name,
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

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "Email verified successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("get user request")
		email := r.URL.Query().Get("email")
		if email == "" {
			s.logger.Error("missing email in request")
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid request"))
			return
		}

		user, err := s.user.GetUser(email)
		if err != nil {
			s.logger.Error("failed to get user", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to get user"))
			return
		}

		s.respond(w, user, http.StatusOK, nil)
	}
}

func (s *Server) handleUpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("update user request")
		var user factory.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			s.logger.Error("failed to decode request body", "error", err)
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
			return
		}

		err = s.user.UpdateUserDetails(user)
		if err != nil {
			s.logger.Error("failed to update user", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to update user"))
			return
		}

		s.respond(w, ResponseMsg{Message: "User updated successfully", Data: "User updated successfully "}, http.StatusOK, nil)
	}
}

func (s *Server) handleDeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("delete user request")
		email := r.URL.Query().Get("email")
		if email == "" {
			s.logger.Error("missing email in request")
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid request"))
			return
		}

		err := s.user.DeleteUser(email)
		if err != nil {
			s.logger.Error("failed to delete user", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to delete user"))
			return
		}

		s.respond(w, ResponseMsg{Message: "User Deleted"}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("get all users request")

		users, err := s.user.GetAllUsers()
		if err != nil {
			s.logger.Error("failed to get all users", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to get all users"))
			return
		}

		s.respond(w, users, http.StatusOK, nil)
	}
}

func (s *Server) EmailExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("email exists request")
		email := r.URL.Query().Get("email")
		if email == "" {
			s.logger.Error("missing email in request")
			s.respond(w, nil, http.StatusBadRequest, fmt.Errorf("invalid request"))
			return
		}

		exists, err := s.user.EmailExists(email)
		if err != nil {
			s.logger.Error("failed to check email", "error", err)
			s.respond(w, nil, http.StatusInternalServerError, fmt.Errorf("failed to check email"))
			return
		}

		s.respond(w, exists, http.StatusOK, nil)
	}
}
