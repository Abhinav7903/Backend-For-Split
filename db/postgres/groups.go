package postgres

import (
	"database/sql"
	"errors"

	"github.com/Abhinav7903/split/factory"
)

// AddGroup adds a new group to the database.
func (p *Postgres) AddGroup(group factory.Group) (int, error) {
	if group.GroupName == "" || group.CreatedBy <= 0 {
		return 0, errors.New("invalid group data")
	}

	var exists bool
	checkUserQuery := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE user_id = $1
		)
	`
	err := p.dbConn.QueryRow(checkUserQuery, group.CreatedBy).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errors.New("creator does not exist")
	}

	insertGroupQuery := `
		INSERT INTO groups (group_name, created_by)
		VALUES ($1, $2)
		RETURNING group_id
	`
	var groupID int
	err = p.dbConn.QueryRow(insertGroupQuery, group.GroupName, group.CreatedBy).Scan(&groupID)
	if err != nil {
		return 0, err
	}

	return groupID, nil
}

// GetGroup retrieves a group by its ID.
func (p *Postgres) GetGroup(groupID int) (factory.Group, error) {
	var group factory.Group
	query := `
		SELECT group_id, group_name, created_by, created_at
		FROM groups
		WHERE group_id = $1
	`
	err := p.dbConn.QueryRow(query, groupID).Scan(
		&group.GroupID,
		&group.GroupName,
		&group.CreatedBy,
		&group.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return factory.Group{}, errors.New("group not found")
		}
		return factory.Group{}, err
	}

	return group, nil
}

// GetAllGroups retrieves all groups from the database.
func (p *Postgres) GetAllGroups() ([]factory.Group, error) {
	query := `
		SELECT group_id, group_name, created_by, created_at
		FROM groups
	`
	rows, err := p.dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []factory.Group
	for rows.Next() {
		var group factory.Group
		err := rows.Scan(&group.GroupID, &group.GroupName, &group.CreatedBy, &group.CreatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

// UpdateGroup updates an existing group's details.
func (p *Postgres) UpdateGroup(group factory.Group) error {
	if group.GroupID <= 0 || group.GroupName == "" {
		return errors.New("invalid group data")
	}

	query := `
		UPDATE groups
		SET group_name = $1
		WHERE group_id = $2
	`
	result, err := p.dbConn.Exec(query, group.GroupName, group.GroupID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("group not found")
	}

	return nil
}

// DeleteGroup deletes a group by its ID.
func (p *Postgres) DeleteGroup(groupID int) error {
	if groupID <= 0 {
		return errors.New("invalid group ID")
	}

	query := `
		DELETE FROM groups
		WHERE group_id = $1
	`
	result, err := p.dbConn.Exec(query, groupID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("group not found")
	}

	return nil
}

// GroupExists checks if a group with the given ID exists.
func (p *Postgres) GroupExists(groupID int) (bool, error) {
	if groupID <= 0 {
		return false, errors.New("invalid group ID")
	}

	query := `
		SELECT EXISTS (
			SELECT 1 FROM groups WHERE group_id = $1
		)
	`
	var exists bool
	err := p.dbConn.QueryRow(query, groupID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
