package factory

import "time"

type Group struct {
	GroupID   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
