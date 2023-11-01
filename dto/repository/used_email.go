package dtorepository

import "time"

type UsedEmailRequest struct {
	ID        int
	AccountID int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UsedEmailResponse struct {
	ID        int
	AccountID int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
