package factory

import "time"

type User struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	FirebaseUID string `json:"firebase_uid"`
	Verified    bool   `json:"verified"`
}

type LoginAttempt struct {
	FailedAttempts int
	LockoutTime    time.Time
}
