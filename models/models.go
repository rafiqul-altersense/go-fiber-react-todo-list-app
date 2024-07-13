package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Todo struct {
	ID        *string `json:"id" default:"1"` // string data type id auto generated
	Title     string  `json:"title"`
	Status    *string `json:"status" default:"active"`
	CreatedAt *string `json:"created_at" default:"2021-09-01T00:00:00Z"`
	UpdatedAt *string `json:"updated_at" default:"2021-09-01T00:00:00Z"`
	DeletedAt *string `json:"deleted_at"` // soft delete, initially null
}
