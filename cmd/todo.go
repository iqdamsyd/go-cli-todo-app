package cmd

import "time"

type Todo struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Completed   bool       `json:"completed"`
	Deadline    string     `json:"deadline"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
