package groupmember

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	AddGroupMember(groupMember factory.GroupMember) (int, error)         //AddGroupMember adds a new group member to the database and returns the ID of the new group member
	GetGroupMemberByID(groupMemberID int) (*factory.GroupMember, error)  //GetGroupMemberByID returns the group member with the given ID
	GetGroupMembersByGroupID(groupID int) ([]factory.GroupMember, error) //GetGroupMembersByGroupID returns all the group members in the group with the given ID
	RemoveUserFromGroupByCreator(groupID, userID, creatorID int) error   //RemoveUserFromGroupByCreator removes the user with the given ID from the group with the given ID
	RemoveUserSelf(groupID, userID int) error                            //RemoveUserSelf removes the user with the given ID from the group with the given ID
	// CountMembersInGroup(groupID int) (int, error)                                      //CountMembersInGroup returns the number of members in the group with the given ID
	// ListGroupsForUser(userID int) ([]factory.Group, error)                             //ListGroupsForUser returns all the groups that the user with the given ID is a member of
}
