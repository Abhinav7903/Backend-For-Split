package postgres

import (
	"errors"

	"github.com/Abhinav7903/split/factory"
)

// AddRequest inserts a new request into the "requests" table
func (p *Postgres) AddRequest(data factory.Request) error {
	query := `INSERT INTO requests (sender_id, receiver_id, group_id, amount, status, created_at) 
	          VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)`

	_, err := p.dbConn.Exec(query, data.SenderID, data.ReceiverID, data.GroupID, data.Amount, data.Status)
	if err != nil {
		return errors.New("failed to add request: " + err.Error())
	}
	return nil
}

// GetRequestByID fetches a request by its ID
func (p *Postgres) GetRequestByID(requestID int) (factory.Request, error) {
	query := `SELECT sender_id, receiver_id, group_id, amount, status FROM requests WHERE request_id = $1`
	row := p.dbConn.QueryRow(query, requestID)

	var data factory.Request
	err := row.Scan(&data.SenderID, &data.ReceiverID, &data.GroupID, &data.Amount, &data.Status)
	if err != nil {
		return factory.Request{}, errors.New("failed to get request: " + err.Error())
	}
	return data, nil
}

// UpdateRequestStatus updates the status of a request
func (p *Postgres) UpdateRequestStatus(requestID int, status string) error {
	query := `UPDATE requests SET status = $1 WHERE request_id = $2`
	_, err := p.dbConn.Exec(query, status, requestID)
	if err != nil {
		return errors.New("failed to update request status: " + err.Error())
	}
	return nil
}

// DeleteRequest deletes a request by its ID
func (p *Postgres) DeleteRequest(requestID int) error {
	query := `DELETE FROM requests WHERE request_id = $1`
	_, err := p.dbConn.Exec(query, requestID)
	if err != nil {
		return errors.New("failed to delete request: " + err.Error())
	}
	return nil
}

// GetRequestsByReceiverID fetches all requests received by a user
func (p *Postgres) GetRequestsByReceiverID(receiverID int) ([]factory.Request, error) {
	query := `SELECT request_id, sender_id, group_id, amount, status FROM requests WHERE receiver_id = $1`
	rows, err := p.dbConn.Query(query, receiverID)
	if err != nil {
		return nil, errors.New("failed to get requests: " + err.Error())
	}
	defer rows.Close()

	var requests []factory.Request
	for rows.Next() {
		var data factory.Request
		err := rows.Scan(&data.RequestID, &data.SenderID, &data.GroupID, &data.Amount, &data.Status)
		if err != nil {
			return nil, errors.New("failed to get requests: " + err.Error())
		}
		requests = append(requests, data)
	}
	return requests, nil
}

// GetRequestsBySenderID fetches all requests sent by a user
func (p *Postgres) GetRequestsBySenderID(senderID int) ([]factory.Request, error) {
	query := `SELECT request_id, receiver_id, group_id, amount, status FROM requests WHERE sender_id = $1`
	rows, err := p.dbConn.Query(query, senderID)
	if err != nil {
		return nil, errors.New("failed to get requests: " + err.Error())
	}
	defer rows.Close()

	var requests []factory.Request
	for rows.Next() {
		var data factory.Request
		err := rows.Scan(&data.RequestID, &data.ReceiverID, &data.GroupID, &data.Amount, &data.Status)
		if err != nil {
			return nil, errors.New("failed to get requests: " + err.Error())
		}
		requests = append(requests, data)
	}
	return requests, nil
}

// GetRequestsByGroupID fetches all requests sent to a group
func (p *Postgres) GetRequestsByGroupID(groupID int) ([]factory.Request, error) {
	query := `SELECT request_id, sender_id, receiver_id, amount, status FROM requests WHERE group_id = $1`
	rows, err := p.dbConn.Query(query, groupID)
	if err != nil {
		return nil, errors.New("failed to get requests: " + err.Error())
	}
	defer rows.Close()

	var requests []factory.Request
	for rows.Next() {
		var data factory.Request
		err := rows.Scan(&data.RequestID, &data.SenderID, &data.ReceiverID, &data.Amount, &data.Status)
		if err != nil {
			return nil, errors.New("failed to get requests: " + err.Error())
		}
		requests = append(requests, data)
	}
	return requests, nil
}


