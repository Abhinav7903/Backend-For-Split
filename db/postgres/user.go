package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) AddUser(user factory.User) error {
	query := `
		INSERT INTO users (email, name, firebase_id) 
		VALUES ($1, $2, $3)
	`
	_, err := p.dbConn.Exec(query, user.Email, user.Name, user.FirebaseUID)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	return nil
}

func (p *Postgres) VerifyEmail(email string) error {
	_, err := p.dbConn.Exec(
		"UPDATE users SET verified=TRUE WHERE email=$1",
		email,
	)
	if err != nil {
		return fmt.Errorf("failed to verify email %s: %w", email, err)
	}
	return nil
}

// Getuser Details from the database
func (p *Postgres) GetUser(email string) (factory.User, error) {
	var user factory.User
	err := p.dbConn.QueryRow("SELECT email, name, firebase_uid, verified FROM users WHERE email=$1", email).Scan(&user.Email, &user.Name, &user.FirebaseUID, &user.Verified)
	if err != nil {
		return user, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// UpdateUserDetails updates the user details in the database
func (p *Postgres) UpdateUserDetails(user factory.User) error {
	_, err := p.dbConn.Exec("UPDATE users SET name=$1, firebase_uid=$2 WHERE email=$3", user.Name, user.FirebaseUID, user.Email)
	if err != nil {
		return fmt.Errorf("failed to update user details: %w", err)
	}
	return nil
}

// DeleteUser deletes the user from the database
func (p *Postgres) DeleteUser(email string) error {
	_, err := p.dbConn.Exec("DELETE FROM users WHERE email=$1", email)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// GetAllUsers gets all the users from the database
func (p *Postgres) GetAllUsers() ([]factory.User, error) {
	var users []factory.User
	rows, err := p.dbConn.Query("SELECT email, name, firebase_uid, verified FROM users")
	if err != nil {
		return users, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user factory.User
		err := rows.Scan(&user.Email, &user.Name, &user.FirebaseUID, &user.Verified)
		if err != nil {
			return users, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUserVerification updates the user verification status in the database
func (p *Postgres) UpdateUserVerification(email string, status bool) error {
	_, err := p.dbConn.Exec("UPDATE users SET verified=$1 WHERE email=$2", status, email)
	if err != nil {
		return fmt.Errorf("failed to update user verification status: %w", err)
	}
	return nil
}

// EmailExists checks if the email exists in the database
func (p *Postgres) EmailExists(email string) (bool, error) {
	var exists bool
	err := p.dbConn.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("failed to check if email exists: %w", err)
	}
	return exists, nil
}

func (p *Postgres) GetUserIDByEmail(email string) (int, error) {
	const query = `
        SELECT user_id 
        FROM users 
        WHERE email = $1
    `
	var userID int
	err := p.dbConn.QueryRow(query, email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user not found: %w", err)
		}
		return 0, fmt.Errorf("failed to get user ID: %w", err)
	}

	return userID, nil
}

func (p *Postgres) GetUserByID(id int) (factory.User, error) {
	const query = `
		SELECT email, name, verified
		FROM users
		WHERE user_id = $1
	`
	var user factory.User
	err := p.dbConn.QueryRow(query, id).Scan(&user.Email, &user.Name, &user.Verified)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found: %w", err)
		}
		return user, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
