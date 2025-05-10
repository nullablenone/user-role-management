package user

import (
	"manajemen-user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	utils.RespondSuccess(c, users, "Users fetched successfully")
}

func (h *Handler) GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.ServiceGetUsersByID(id)
	if err != nil {
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
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.RespondSuccess(c, user, "User updated successfully")
}

func (h *Handler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.ServiceDeleteUsers(id)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.RespondSuccess(c, nil, "User deleted successfully")
}
