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

func (p *Postgres) RemoveUserFromGroupByCreator(groupID, userID, creatorID int) error {
	// Validate input
	if groupID <= 0 || userID <= 0 || creatorID <= 0 {
		return errors.New("invalid input: groupID, userID, and creatorID must be greater than zero")
	}

	// Check if the user is part of the group
	var isMember bool
	const checkUserInGroupQuery = `
		SELECT EXISTS (
			SELECT 1 FROM group_members
			WHERE group_id = $1 AND user_id = $2
		)
	`
	err := p.dbConn.QueryRow(checkUserInGroupQuery, groupID, userID).Scan(&isMember)
	if err != nil {
		return fmt.Errorf("failed to check if user is part of the group: %w", err)
	}
	if !isMember {
		return errors.New("user is not a member of the group")
	}

	// Check if the user is the creator of the group
	var createdBy int
	const checkCreatorQuery = `
        SELECT created_by 
        FROM groups
        WHERE group_id = $1
    `
	err = p.dbConn.QueryRow(checkCreatorQuery, groupID).Scan(&createdBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("group not found")
		}
		return fmt.Errorf("failed to check group creator: %w", err)
	}

	// Ensure the creator is the one making the request
	if createdBy != creatorID {
		return errors.New("only the group creator can remove users from the group")
	}

	// Remove the user from the group
	const removeUserQuery = `
        DELETE FROM group_members
        WHERE group_id = $1 AND user_id = $2
    `
	result, err := p.dbConn.Exec(removeUserQuery, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %w", err)
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user was not removed, possibly not a member of the group")
	}

	return nil
}

// RemoveUserSelf removes the user from the group they belong to
func (p *Postgres) RemoveUserSelf(groupID, userID int) error {
	// Validate input
	if groupID <= 0 || userID <= 0 {
		return errors.New("invalid input")
	}

	// Check if the user is part of the group
	var exists bool
	const checkUserInGroupQuery = `
        SELECT EXISTS (
            SELECT 1 FROM group_members
            WHERE group_id = $1 AND user_id = $2
        )
    `
	err := p.dbConn.QueryRow(checkUserInGroupQuery, groupID, userID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if user is part of the group: %w", err)
	}

	if !exists {
		return errors.New("user is not a member of the group")
	}

	// Remove the user from the group
	const removeUserQuery = `
        DELETE FROM group_members
        WHERE group_id = $1 AND user_id = $2
    `
	_, err = p.dbConn.Exec(removeUserQuery, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %w", err)
	}

	return nil
}
