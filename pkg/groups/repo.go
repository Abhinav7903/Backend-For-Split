package groups

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	AddGroup(group factory.Group) (int, error)   // Return the ID of the new group
	GetGroup(groupID int) (factory.Group, error) // Use int for groupID
	GetAllGroups() ([]factory.Group, error)
	UpdateGroup(group factory.Group) error
	DeleteGroup(groupID int) error // Use int for groupID
	GroupExists(groupID int) (bool, error)
}
