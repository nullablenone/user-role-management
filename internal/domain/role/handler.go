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

func (h *Handler) GetRoles(c *gin.Context) {
	roles, err := h.Service.ServiceGetRoles()
	if err != nil {
		log.Printf("Error in GetRoles: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch roles")
		return
	}

	utils.RespondSuccess(c, roles, "Roles fetched successfully")
}

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

	utils.RespondSuccess(c, role, "Role fetched successfully")
}

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

	utils.RespondSuccess(c, role, "Role created successfully")
}

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

	utils.RespondSuccess(c, role, "Role updated successfully")
}

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
