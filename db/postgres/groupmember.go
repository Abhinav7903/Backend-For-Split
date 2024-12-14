package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) AddGroupMember(groupMember factory.GroupMember) (int, error) {
	// Validate input
	if groupMember.GroupID <= 0 || groupMember.UserID <= 0 {
		return 0, errors.New("invalid group member data: group_id and user_id must be greater than zero")
	}

	// Use transaction for atomicity
	tx, err := p.dbConn.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Check if group and user exist
	const checkExistsQuery = `
        SELECT 
            EXISTS (SELECT 1 FROM groups WHERE group_id = $1) AS group_exists,
            EXISTS (SELECT 1 FROM users WHERE user_id = $2) AS user_exists
    `
	var groupExists, userExists bool
	err = tx.QueryRow(checkExistsQuery, groupMember.GroupID, groupMember.UserID).Scan(&groupExists, &userExists)
	if err != nil {
		slog.Error("failed to check existence of group and user", "error", err)
		return 0, fmt.Errorf("failed to check existence of group and user: %w", err)
	}
	if !groupExists {
		slog.Error("group does not exist")
		return 0, errors.New("group does not exist")
	}
	if !userExists {
		slog.Error("user does not exist")
		return 0, errors.New("user does not exist")
	}

	// Insert group member
	const insertGroupMemberQuery = `
        INSERT INTO group_members (group_id, user_id)
        VALUES ($1, $2)
        RETURNING group_member_id
    `
	var groupMemberID int
	err = tx.QueryRow(insertGroupMemberQuery, groupMember.GroupID, groupMember.UserID).Scan(&groupMemberID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert group member: %w", err)
	}

	return groupMemberID, nil
}

// GetGroupMemberByID returns the group member with the given ID
// GetGroupMemberByID returns the group member with the given ID
func (p *Postgres) GetGroupMemberByID(groupMemberID int) (*factory.GroupMember, error) {
	// Validate input
	if groupMemberID <= 0 {
		return nil, errors.New("invalid group member ID: must be greater than zero")
	}

	// Query group member
	const query = `
        SELECT group_id, user_id, joined_date
        FROM group_members
        WHERE group_member_id = $1
    `
	groupMember := &factory.GroupMember{}

	err := p.dbConn.QueryRow(query, groupMemberID).Scan(&groupMember.GroupID, &groupMember.UserID, &groupMember.JoinedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("group member not found")
		}
		return nil, fmt.Errorf("failed to query group member: %w", err)
	}

	// Set the group member ID explicitly
	groupMember.GroupMemberID = groupMemberID

	return groupMember, nil
}

// GetGroupMembersByGroupID returns all the group members in the group with the given ID
func (p *Postgres) GetGroupMembersByGroupID(groupID int) ([]factory.GroupMember, error) {
	// Validate input
	if groupID <= 0 {
		return nil, errors.New("invalid group ID: must be greater than zero")
	}

	// Query group members
	const query = `
		SELECT group_member_id, user_id,joined_date
		FROM group_members
		WHERE group_id = $1
	`
	rows, err := p.dbConn.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to query group members: %w", err)
	}
	defer rows.Close()

	groupMembers := []factory.GroupMember{}
	for rows.Next() {
		groupMember := factory.GroupMember{GroupID: groupID}
		err := rows.Scan(&groupMember.GroupMemberID, &groupMember.UserID, &groupMember.JoinedDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group member: %w", err)
		}
		groupMembers = append(groupMembers, groupMember)
	}

	return groupMembers, nil
}
