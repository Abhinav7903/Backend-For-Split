package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhinav7903/split/factory"
)

func (s *Server) handleAddGroupMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the HTTP method is POST
		if r.Method != http.MethodPost {
			s.respond(w, ResponseMsg{Message: "Method not allowed"}, http.StatusMethodNotAllowed, nil)
			return
		}

		// Ensure the request body is closed
		defer r.Body.Close()

		// Parse the JSON payload
		var groupMember factory.GroupMember
		if err := json.NewDecoder(r.Body).Decode(&groupMember); err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid request payload"}, http.StatusBadRequest, nil)
			return
		}

		// Validate the input data
		if groupMember.GroupID <= 0 || groupMember.UserID <= 0 {
			s.respond(w, ResponseMsg{Message: "Invalid group member data: group_id and user_id must be greater than zero"}, http.StatusBadRequest, nil)
			return
		}

		// Add the group member via repository
		groupMemberID, err := s.group_members.AddGroupMember(groupMember)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to add group member"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success and the new group member ID
		s.respond(w, ResponseMsg{
			Message: "Group member added successfully",
			Data:    groupMemberID,
		}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetGroupMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the HTTP method is GET
		if r.Method != http.MethodGet {
			s.respond(w, ResponseMsg{Message: "Method not allowed"}, http.StatusMethodNotAllowed, nil)
			return
		}

		// Parse the group member ID from the query parameters
		groupMemberID := r.URL.Query().Get("group_member_id")
		if groupMemberID == "" {
			s.respond(w, ResponseMsg{Message: "group_member_id is required"}, http.StatusBadRequest, nil)
			return
		}

		// Convert the group member ID to an integer
		id, err := strconv.Atoi(groupMemberID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid group_member_id"}, http.StatusBadRequest, nil)
			return
		}

		// Retrieve the group member from the repository
		groupMember, err := s.group_members.GetGroupMemberByID(id)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to get group member"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the group member
		s.respond(w, ResponseMsg{
			Message: "Group member retrieved successfully",
			Data:    groupMember,
		}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetGroupMembersByGroupID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the HTTP method is GET
		if r.Method != http.MethodGet {
			s.respond(w, ResponseMsg{Message: "Method not allowed"}, http.StatusMethodNotAllowed, nil)
			return
		}

		// Parse the group ID from the query parameters
		groupID := r.URL.Query().Get("group_id")
		if groupID == "" {
			s.respond(w, ResponseMsg{Message: "group_id is required"}, http.StatusBadRequest, nil)
			return
		}

		// Convert the group ID to an integer
		id, err := strconv.Atoi(groupID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid group_id"}, http.StatusBadRequest, nil)
			return
		}

		// Retrieve the group members from the repository
		groupMembers, err := s.group_members.GetGroupMembersByGroupID(id)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to get group members"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the group members
		s.respond(w, ResponseMsg{
			Message: "Group members retrieved successfully",
			Data:    groupMembers,
		}, http.StatusOK, nil)
	}
}
