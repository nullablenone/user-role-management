package role

type CreateRolesRequest struct {
	Name        string `json:"name" binding:"required"`
	Deskription string `json:"deskription" binding:"required"`
}

type UpdateRolesRequest struct {
	Name        string `json:"name" binding:"required"`
	Deskription string `json:"deskription" binding:"required"`
}
