package todo

import "time"

// Todo is Todo Model
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status`
	LimitDate time.Time `json:"limit_date"`
}
