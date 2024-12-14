package factory

import "time"

type GroupMember struct {
	GroupMemberID int       `json:"group_member_id"` // Maps to group_member_id (primary key)
	UserID        int       `json:"user_id"`         // Maps to user_id (foreign key to users table)
	GroupID       int       `json:"group_id"`        // Maps to group_id (foreign key to groups table)
	JoinedDate    time.Time `json:"joined_date"`     // Maps to joined_date
}
