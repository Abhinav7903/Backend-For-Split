package server

import (
	"encoding/json"
	"fmt"
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

func (s *Server) handleRemoveUserFromGroupByCreator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the HTTP method is DELETE
		if r.Method != http.MethodDelete {
			s.respond(w, ResponseMsg{Message: "Method not allowed"}, http.StatusMethodNotAllowed, nil)
			return
		}

		// Get group ID, user ID, and creator ID from URL params or request body
		groupIDStr := r.URL.Query().Get("group_id")
		userIDStr := r.URL.Query().Get("user_id")
		creatorIDStr := r.URL.Query().Get("creator_id")

		if groupIDStr == "" || userIDStr == "" || creatorIDStr == "" {
			s.respond(w, ResponseMsg{Message: "Missing required parameters"}, http.StatusBadRequest, nil)
			return
		}

		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid group ID"}, http.StatusBadRequest, nil)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid user ID"}, http.StatusBadRequest, nil)
			return
		}

		creatorID, err := strconv.Atoi(creatorIDStr)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid creator ID"}, http.StatusBadRequest, nil)
			return
		}

		// Remove user from the group by the group creator
		err = s.group_members.RemoveUserFromGroupByCreator(groupID, userID, creatorID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: fmt.Sprintf("Failed to remove user from group: %v", err)}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "User removed from group successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handleRemoveUserSelfFromGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the HTTP method is DELETE
		if r.Method != http.MethodDelete {
			s.respond(w, ResponseMsg{Message: "Method not allowed"}, http.StatusMethodNotAllowed, nil)
			return
		}

		// Get group ID and user ID from URL params or request body
		groupIDStr := r.URL.Query().Get("group_id")
		userIDStr := r.URL.Query().Get("user_id")

		if groupIDStr == "" || userIDStr == "" {
			s.respond(w, ResponseMsg{Message: "Missing required parameters"}, http.StatusBadRequest, nil)
			return
		}

		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid group ID"}, http.StatusBadRequest, nil)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Invalid user ID"}, http.StatusBadRequest, nil)
			return
		}

		// Remove the user from the group
		err = s.group_members.RemoveUserSelf(groupID, userID)
		if err != nil {
			s.respond(w, ResponseMsg{Message: fmt.Sprintf("Failed to remove user from group: %v", err)}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success message
		s.respond(w, ResponseMsg{Message: "User removed from group successfully"}, http.StatusOK, nil)
	}
}
