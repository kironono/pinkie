package model

import "time"

type JobID int64

type Job struct {
	ID        JobID     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Slug      string    `json:"slug" db:"slug"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type Jobs []*Job

type JobSessionID int64
type JobSessionStatus string

const (
	JobSessionStatusRunning  JobSessionStatus = "running"
	JobSessionStatusFinished JobSessionStatus = "finished"
	JobSessionStatusFailed   JobSessionStatus = "failed"
)

type JobSession struct {
	ID         JobSessionID     `json:"id" db:"id"`
	JobID      JobID            `json:"job_id" db:"job_id"`
	LastStatus JobSessionStatus `json:"last_status" db:"last_status"`
	StartAt    time.Time        `json:"start_at" db:"start_at"`
	EndAt      time.Time        `json:"end_at" db:"end_at"`
	CreatedAt  time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at" db:"updated_at"`
}

type JobSessionEventID int64

type JobSessionEvent struct {
	ID           JobSessionEventID `json:"id" db:"id"`
	JobSessionID JobSessionID      `json:"job_session_id" db:"job_session_id"`
	Timestamp    time.Time         `json:"timestamp" db:"timestamp"`
	Status       JobSessionStatus  `json:"status" db:"status"`
}
