package user

import (
	"errors"
	"fmt"
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

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.Service.ServiceGetUsers()
	if err != nil {
		log.Printf("Error in GetUsers: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	utils.RespondSuccess(c, users, "Users fetched successfully")
}

func (h *Handler) GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.ServiceGetUsersByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error in GetUsersByID: %v", err)
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondSuccess(c, user, "User fetched successfully")

}

func (h *Handler) CreateUsers(c *gin.Context) {
	var input CreateUsersRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.Service.ServiceCreateUsers(input)
	if err != nil {
		log.Printf("Error in CreateUsers: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondSuccess(c, user, "User created successfully")
}

func (h *Handler) UpdateUsers(c *gin.Context) {
	id := c.Param("id")
	var input UpdateUsersRequest
	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.Service.ServiceUpdateUsers(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error in UpdateUsers: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.RespondSuccess(c, user, "User updated successfully")
}

func (h *Handler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.ServiceDeleteUsers(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error in DeleteUsers: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.RespondSuccess(c, nil, "User deleted successfully")
}

func (h *Handler) Profile(c *gin.Context) {
	mapClaims, exists := c.Get("claims")
	if !exists {
		utils.RespondError(c, http.StatusUnauthorized, "RequireRole: failed to take claims")
		c.Abort()
		return
	}

	claims, err := utils.AssertTypeClaims(mapClaims)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, fmt.Sprintf("RequireRole: %v", err))
		c.Abort()
		return
	}

	user_id := claims["user_id"].(float64)
	user, err := h.Service.ServiceProfileUsers(user_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondError(c, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error in Profile: %v", err)
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondSuccess(c, user, "User fetched successfully")
}
