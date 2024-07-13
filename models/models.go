package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Todo struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Status    string  `json:"status" default:"active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"` // soft delete, initially null
}
