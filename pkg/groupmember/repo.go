package groupmember

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	AddGroupMember(groupMember factory.GroupMember) (int, error)         //AddGroupMember adds a new group member to the database and returns the ID of the new group member
	GetGroupMemberByID(groupMemberID int) (*factory.GroupMember, error)  //GetGroupMemberByID returns the group member with the given ID
	GetGroupMembersByGroupID(groupID int) ([]factory.GroupMember, error) //GetGroupMembersByGroupID returns all the group members in the group with the given ID
	// UpdateGroupMember(groupMemberID int, updatedGroupMember factory.GroupMember) error //UpdateGroupMember updates the group member with the given ID
	// DeleteGroupMember(groupMemberID int) error                                         //DeleteGroupMember deletes the group member with the given ID
	// CountMembersInGroup(groupID int) (int, error)                                      //CountMembersInGroup returns the number of members in the group with the given ID
	// ListGroupsForUser(userID int) ([]factory.Group, error)                             //ListGroupsForUser returns all the groups that the user with the given ID is a member of
}
