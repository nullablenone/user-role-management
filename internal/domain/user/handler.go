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

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all registered users in the system.
// @Tags Admin - Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /admin/users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.Service.ServiceGetUsers()
	if err != nil {
		log.Printf("Error in GetUsers: %v", err)
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	var responese []ResponseUsers
	for _, user := range users {
		responese = append(responese, ResponseUsers{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	utils.RespondSuccess(c, responese, "Users fetched successfully")
}

// GetUsersByID godoc
// @Summary Get user by ID
// @Description Retrieve a specific user by their ID.
// @Tags Admin - Users
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/users/{id} [get]
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

	response := ResponseUsers{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "User fetched successfully")

}

// CreateUsers godoc
// @Summary Create a new user
// @Description Create a new user account with the given details.
// @Tags Admin - Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateUsersRequest true "User payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /admin/users [post]
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

	response := ResponseUsers{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "User created successfully")
}

// UpdateUsers godoc
// @Summary Update user by ID
// @Description Update an existing user's information based on user ID.
// @Tags Admin - Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body UpdateUsersRequest true "Updated data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/users/{id} [put]
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

	response := ResponseUsers{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "User updated successfully")
}

// DeleteUsers godoc
// @Summary Delete user by ID
// @Description Delete a specific user by their ID.
// @Tags Admin - Users
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/users/{id} [delete]
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

// Profile godoc
// @Summary Get current user profile
// @Description Retrieve the profile information of the currently authenticated user.
// @Tags  User - Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user/profile [get]
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

	response := ResponseUsers{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.RespondSuccess(c, response, "User fetched successfully")
}
