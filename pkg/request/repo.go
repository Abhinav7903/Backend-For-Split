package request

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	AddRequest(data factory.Request) error
	GetRequestByID(requestID int) (factory.Request, error)
	UpdateRequestStatus(requestID int, status string) error
	DeleteRequest(requestID int) error
	GetRequestsByReceiverID(receiverID int) ([]factory.Request, error)
	GetRequestsBySenderID(senderID int) ([]factory.Request, error)
	GetRequestsByGroupID(groupID int) ([]factory.Request, error)
}
