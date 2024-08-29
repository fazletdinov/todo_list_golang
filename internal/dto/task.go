package dto

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type DueDate time.Time

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type TaskRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	DueDate     DueDate `json:"due_date"`
}

type TaskResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     DueDate   `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

func (d DueDate) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2 Jan 2006 3:04PM")
	log.Printf("Formatted ============= %v\n", formatted)
	return json.Marshal(formatted)
}

func (d *DueDate) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}
	log.Printf("s ============= %v\n", s)

	t, err := time.Parse("2 Jan 2006 3:04PM", s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	log.Printf("t ============= %v\n", t)

	*d = DueDate(t)

	return nil
}
