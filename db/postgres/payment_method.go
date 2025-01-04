package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Abhinav7903/split/factory"
)

// CreatePaymentMethod creates a new payment method for a user identified by email
func (p *Postgres) CreatePaymentMethod(paymentMethod *factory.PaymentMethod) error {
	// Fetch user_id based on email
	userID, err := p.GetUserIDByEmail(paymentMethod.Email)
	if err != nil {
		return fmt.Errorf("failed to retrieve user_id for email %s: %w", paymentMethod.Email, err)
	}

	// Log the fetched user_id
	if userID <= 0 {
		return fmt.Errorf("invalid user_id for email %s: %d", paymentMethod.Email, userID)
	}
	paymentMethod.UserID = userID

	// Log the user_id to ensure it’s set
	slog.Info("UserID fetched and assigned for email %s: %d", paymentMethod.Email, paymentMethod.UserID)

	// Insert payment method into the database
	query := `
        INSERT INTO payment_methods (
            user_id, payment_type, upi_id, account_number, ifsc_code, wallet_provider, is_primary
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING payment_id`

	var paymentID int
	err = p.dbConn.QueryRow(query,
		paymentMethod.UserID,
		paymentMethod.PaymentType,
		paymentMethod.UPIID,
		paymentMethod.AccountNumber,
		paymentMethod.IFSCCode,
		paymentMethod.WalletProvider,
		paymentMethod.IsPrimary,
	).Scan(&paymentID)
	if err != nil {
		return fmt.Errorf("failed to create payment method: %w", err)
	}
	paymentMethod.PaymentID = paymentID

	// Return success
	return nil
}

// GetPaymentMethods retrieves all payment methods for a user identified by email
func (p *Postgres) GetPaymentMethods(email string) ([]factory.PaymentMethod, error) {
	// Fetch user_id based on email
	userID, err := p.GetUserIDByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user_id for email %s: %w", email, err)
	}

	// Query to get payment methods for the user
	query := `
        SELECT 
            payment_id, user_id, payment_type, upi_id, account_number, ifsc_code, wallet_provider, is_primary, created_at 
        FROM payment_methods 
        WHERE user_id = $1`

	rows, err := p.dbConn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}
	defer rows.Close()

	var methods []factory.PaymentMethod
	for rows.Next() {
		var method factory.PaymentMethod
		err := rows.Scan(
			&method.PaymentID,
			&method.UserID,
			&method.PaymentType,
			&method.UPIID,
			&method.AccountNumber,
			&method.IFSCCode,
			&method.WalletProvider,
			&method.IsPrimary,
			&method.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan payment method: %w", err)
		}
		methods = append(methods, method)
	}
	return methods, nil
}

// UpdatePaymentMethod updates an existing payment method for a user identified by email
func (p *Postgres) UpdatePaymentMethod(paymentMethod *factory.PaymentMethod) error {
	// Ensure email and payment_type are provided
	if paymentMethod.Email == "" {
		return fmt.Errorf("email is required")
	}
	if paymentMethod.PaymentType == "" {
		return fmt.Errorf("payment_type is required")
	}

	// Fetch user_id based on email
	userID, err := p.GetUserIDByEmail(paymentMethod.Email)
	if err != nil {
		return fmt.Errorf("failed to retrieve user_id for email %s: %w", paymentMethod.Email, err)
	}

	// Ensure user_id is positive
	if userID <= 0 {
		return fmt.Errorf("invalid user_id for email %s: %d", paymentMethod.Email, userID)
	}

	// Find the existing payment method for the user based on payment_type
	var existingPaymentMethod factory.PaymentMethod
	query := `
        SELECT payment_id, payment_type, upi_id, account_number, ifsc_code, wallet_provider, is_primary
        FROM payment_methods 
        WHERE user_id = $1 AND payment_type = $2
        LIMIT 1`

	err = p.dbConn.QueryRow(query, userID, paymentMethod.PaymentType).Scan(
		&existingPaymentMethod.PaymentID,
		&existingPaymentMethod.PaymentType,
		&existingPaymentMethod.UPIID,
		&existingPaymentMethod.AccountNumber,
		&existingPaymentMethod.IFSCCode,
		&existingPaymentMethod.WalletProvider,
		&existingPaymentMethod.IsPrimary,
	)

	// If no payment method is found, return an error
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no payment method found for user %s with payment type %s", paymentMethod.Email, paymentMethod.PaymentType)
		}
		return fmt.Errorf("failed to find payment method: %w", err)
	}

	// If the payment method exists, update the fields that were provided
	updateQuery := `
        UPDATE payment_methods
        SET
            upi_id = COALESCE($1, upi_id),
            account_number = COALESCE($2, account_number),
            ifsc_code = COALESCE($3, ifsc_code),
            wallet_provider = COALESCE($4, wallet_provider),
            is_primary = COALESCE($5, is_primary)
        WHERE payment_id = $6 AND user_id = $7`

	_, err = p.dbConn.Exec(updateQuery,
		paymentMethod.UPIID,
		paymentMethod.AccountNumber,
		paymentMethod.IFSCCode,
		paymentMethod.WalletProvider,
		paymentMethod.IsPrimary,
		existingPaymentMethod.PaymentID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("failed to update payment method: %w", err)
	}

	return nil
}

func (p *Postgres) DeletePaymentMethod(paymentType, email string) error {
	// Ensure email and payment_type are provided
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if paymentType == "" {
		return fmt.Errorf("payment_type is required")
	}

	// Fetch user_id based on email
	userID, err := p.GetUserIDByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to retrieve user_id for email %s: %w", email, err)
	}

	// Ensure user_id is positive
	if userID <= 0 {
		return fmt.Errorf("invalid user_id for email %s: %d", email, userID)
	}

	// Check if the payment method exists for the user by payment_type
	var existingPaymentMethod factory.PaymentMethod
	query := `
        SELECT payment_id
        FROM payment_methods
        WHERE payment_type = $1 AND user_id = $2
        LIMIT 1`

	err = p.dbConn.QueryRow(query, paymentType, userID).Scan(&existingPaymentMethod.PaymentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("payment method with payment_type %s not found for user %s", paymentType, email)
		}
		return fmt.Errorf("failed to find payment method: %w", err)
	}

	// Proceed with deletion of the payment method
	deleteQuery := `
        DELETE FROM payment_methods
        WHERE payment_type = $1 AND user_id = $2`

	_, err = p.dbConn.Exec(deleteQuery, paymentType, userID)
	if err != nil {
		return fmt.Errorf("failed to delete payment method: %w", err)
	}

	return nil
}
