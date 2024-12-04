package postgres

import (
	"fmt"
	"time"

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

func (p *Postgres) LoginUser(factory.User) (string, bool, string, time.Time, error) {
	return "", false, "", time.Time{}, nil
}
