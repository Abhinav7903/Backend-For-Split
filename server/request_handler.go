package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"sync"

	"github.com/Abhinav7903/split/factory"
)

func (s *Server) handleAddRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request factory.Request

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			slog.Error("Failed to decode request", "error", err)
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		if request.Amount <= 0 || request.SenderID == 0 || request.ReceiverID == 0 {
			s.logger.Error("Invalid request", "request", request)
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		err := s.request.AddRequest(request)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			//get the sender and receiver details
			sender, err := s.user.GetUserByID(request.SenderID)
			if err != nil {
				slog.Error("Failed to get sender details", "error", err)
				return
			}
			receiver, err := s.user.GetUserByID(request.ReceiverID)
			if err != nil {
				slog.Error("Failed to get receiver details", "error", err)
				return
			}

			//send notification to the receiver
			mailErr := s.mail.SendMail(receiver.Email, "Request Received", fmt.Sprintf("You have received a request from %s for an amount of %.2f", sender.Name, request.Amount))
			if mailErr != nil {
				slog.Error("Failed to send mail", "error", mailErr)
				return
			}

			slog.Info("Notification sent successfully", "receiver", receiver.Email)
		}()
		wg.Wait()

		s.respond(w, ResponseMsg{Message: "Success", Data: "Request added successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleDeleteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request factory.Request

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			slog.Error("Failed to decode request", "error", err)
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		if request.RequestID == 0 {
			s.logger.Error("Invalid request", "request", request)
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		err := s.request.DeleteRequest(request.RequestID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Success", Data: "Request deleted successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleUpdateRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request factory.Request

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			slog.Error("Failed to decode request", "error", err)
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		if request.SenderID == 0 || request.ReceiverID == 0 {
			s.logger.Error("Invalid request", "request", request)
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		if request.RequestID == 0 || request.Amount <= 0 {
			s.logger.Error("Invalid request", "request", request)
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		err := s.request.UpdateRequestStatus(request.RequestID, request.Status)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			//get the sender and receiver details
			sender, err := s.user.GetUserByID(request.SenderID)
			if err != nil {
				slog.Error("Failed to get sender details", "error", err)
				return
			}
			receiver, err := s.user.GetUserByID(request.ReceiverID)
			if err != nil {
				slog.Error("Failed to get receiver details", "error", err)
				return
			}

			//send notification to the sender
			mailErr := s.mail.SendMail(sender.Email, "Request Updated", fmt.Sprintf("Your request to %s for an amount of %.2f has been %s", receiver.Name, request.Amount, request.Status))
			if mailErr != nil {
				slog.Error("Failed to send mail", "error", mailErr)
				return
			}
		}()
		wg.Wait()

		s.respond(w, ResponseMsg{Message: "Success", Data: "Request updated successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetRequestsByReceiverID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		receiverID := r.URL.Query().Get("receiver_id")
		if receiverID == "" {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		//convert receiverID to int
		receiverIDInt, err := strconv.Atoi(receiverID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		requests, err := s.request.GetRequestsByReceiverID(receiverIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Success", Data: requests}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetRequestsBySenderID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		senderID := r.URL.Query().Get("sender_id")
		if senderID == "" {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		//convert senderID to int
		senderIDInt, err := strconv.Atoi(senderID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		requests, err := s.request.GetRequestsBySenderID(senderIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Success", Data: requests}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetRequestsByGroupID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupID := r.URL.Query().Get("group_id")
		if groupID == "" {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		//convert groupID to int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		requests, err := s.request.GetRequestsByGroupID(groupIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Success", Data: requests}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetRequestByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.URL.Query().Get("request_id")
		if requestID == "" {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		//convert requestID to int
		requestIDInt, err := strconv.Atoi(requestID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: "Invalid request"}, http.StatusBadRequest, nil)
			return
		}

		request, err := s.request.GetRequestByID(requestIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed", Data: err.Error()}, http.StatusBadRequest, nil)
			return
		}

		s.respond(w, ResponseMsg{Message: "Success", Data: request}, http.StatusOK, nil)
	}
}
