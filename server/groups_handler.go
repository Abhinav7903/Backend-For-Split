package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhinav7903/split/factory"
)

func (s *Server) handlerAddGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON payload
		var group factory.Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the input
		if group.GroupName == "" || group.CreatedBy <= 0 {
			http.Error(w, "Invalid group data", http.StatusBadRequest)
			return
		}

		// Add the group
		groupID, err := s.group.AddGroup(group)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to add group"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the group ID
		s.respond(w, ResponseMsg{Message: "Group added successfully", Data: groupID}, http.StatusOK, nil)

	}
}

func (s *Server) handlerGetGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the group ID from the query parameters
		groupID := r.URL.Query().Get("group_id")
		if groupID == "" {
			http.Error(w, "group_id is required", http.StatusBadRequest)
			return
		}

		// Convert the group ID to an integer
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			http.Error(w, "Invalid group_id", http.StatusBadRequest)
			return
		}

		// Retrieve the group
		group, err := s.group.GetGroup(groupIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to get group"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the group details
		s.respond(w, ResponseMsg{Message: "Group fetched successfully", Data: group}, http.StatusOK, nil)
	}
}

func (s *Server) handlerGetAllGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Retrieve all groups
		groups, err := s.group.GetAllGroups()
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to get all groups"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the groups
		s.respond(w, ResponseMsg{Message: "Groups fetched successfully", Data: groups}, http.StatusOK, nil)
	}
}

func (s *Server) handlerUpdateGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON payload
		var group factory.Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the input
		if group.GroupName == "" || group.CreatedBy <= 0 {
			http.Error(w, "Invalid group data", http.StatusBadRequest)
			return
		}

		// Update the group
		err = s.group.UpdateGroup(group)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to update group"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success
		s.respond(w, ResponseMsg{Message: "Group updated successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handlerDeleteGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the group ID from the query parameters
		groupID := r.URL.Query().Get("group_id")
		if groupID == "" {
			http.Error(w, "group_id is required", http.StatusBadRequest)
			return
		}

		// Convert the group ID to an integer
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			http.Error(w, "Invalid group_id", http.StatusBadRequest)
			return
		}

		// Delete the group
		err = s.group.DeleteGroup(groupIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to delete group"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with success
		s.respond(w, ResponseMsg{Message: "Group deleted successfully"}, http.StatusOK, nil)
	}
}

func (s *Server) handlerGroupExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for the correct HTTP method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the group ID from the query parameters
		groupID := r.URL.Query().Get("group_id")
		if groupID == "" {
			http.Error(w, "group_id is required", http.StatusBadRequest)
			return
		}

		// Convert the group ID to an integer
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			http.Error(w, "Invalid group_id", http.StatusBadRequest)
			return
		}

		// Check if the group exists
		exists, err := s.group.GroupExists(groupIDInt)
		if err != nil {
			s.respond(w, ResponseMsg{Message: "Failed to check group existence"}, http.StatusInternalServerError, nil)
			return
		}

		// Respond with the result
		response := map[string]interface{}{
			"exists": exists,
		}
		s.respond(w, ResponseMsg{Message: "Group existence check successful", Data: response}, http.StatusOK, nil)
	}
}
