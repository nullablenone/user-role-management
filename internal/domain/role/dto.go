package role

import "time"

type CreateRolesRequest struct {
	Name        string `json:"name" binding:"required"`
	Deskription string `json:"deskription" binding:"required"`
}

type UpdateRolesRequest struct {
	Name        string `json:"name" binding:"required"`
	Deskription string `json:"deskription" binding:"required"`
}

type ResponseRoles struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Deskription string    `json:"deskription"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
