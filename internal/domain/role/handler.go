package role

import (
	"errors"
	"log"
	"manajemen-user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

// GetRoles godoc
// @Summary Get all roles
// @Description Get list of all roles (admin only)
// @Tags Admin - Roles
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{} "Roles fetched successfully"
// @Failure 500 {object} map[string]interface{} "Failed to fetch roles"
// @Router /admin/roles [get]
func (h *Handler) GetRoles(c *gin.Context) {
	roles, err := h.Service.ServiceGetRoles()
	if err != nil {
		log.Printf("Error in GetRoles: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch roles")
		return
	}

	var response []ResponseRoles
	for _, role := range roles {
		response = append(response, ResponseRoles{
			ID:          role.ID,
			Name:        role.Name,
			Deskription: role.Deskription,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}

	utils.RespondSuccess(c, response, "Roles fetched successfully")
}

// GetRolesByID godoc
// @Summary Get role by ID
// @Description Get a specific role by its ID (admin only)
// @Tags Admin - Roles
// @Security BearerAuth
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]interface{} "Role fetched successfully"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Router /admin/roles/{id} [get]
func (h *Handler) GetRolesByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.Service.ServiceGetRolesByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "Role not found")
			return
		}
		log.Printf("Error in GetRolesByID: %v", err)
		utils.RespondError(c, http.StatusNotFound, "Role not found")
		return
	}

	response := ResponseRoles{
		ID:          role.ID,
		Name:        role.Name,
		Deskription: role.Deskription,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "Role fetched successfully")
}

// CreateRoles godoc
// @Summary Create a new role
// @Description Create a new role with given data (admin only)
// @Tags Admin - Roles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateRolesRequest true "Create Role Request"
// @Success 200 {object} map[string]interface{} "Role created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Failed to create role"
// @Router /admin/roles [post]
func (h *Handler) CreateRoles(c *gin.Context) {
	var input CreateRolesRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	role, err := h.Service.ServiceCreateRoles(input)
	if err != nil {
		log.Printf("Error in CreateRoles: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create role")
		return
	}

	response := ResponseRoles{
		ID:          role.ID,
		Name:        role.Name,
		Deskription: role.Deskription,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "Role created successfully")
}

// UpdateRoles godoc
// @Summary Update a role
// @Description Update a role by ID (admin only)
// @Tags Admin - Roles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param request body UpdateRolesRequest true "Update Role Request"
// @Success 200 {object} map[string]interface{} "Role updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 500 {object} map[string]interface{} "Failed to update role"
// @Router /admin/roles/{id} [put]
func (h *Handler) UpdateRoles(c *gin.Context) {
	id := c.Param("id")
	var input UpdateRolesRequest
	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	role, err := h.Service.ServiceUpdateRoles(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "Role not found")
			return
		}
		log.Printf("Error in UpdateRoles: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update role")
		return
	}

	response := ResponseRoles{
		ID:          role.ID,
		Name:        role.Name,
		Deskription: role.Deskription,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "Role updated successfully")
}

// DeleteRoles godoc
// @Summary Delete a role
// @Description Delete a role by ID (admin only)
// @Tags Admin - Roles
// @Security BearerAuth
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]interface{} "Role deleted successfully"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete role"
// @Router /admin/roles/{id} [delete]
func (h *Handler) DeleteRoles(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.ServiceDeleteRoles(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "Role not found")
			return
		}
		log.Printf("Error in DeleteRoles: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete role")
		return
	}

	utils.RespondSuccess(c, nil, "Role deleted successfully")
}
