package auth

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

func (h *Handler) Register(c *gin.Context) {
	var input RegisterRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.Service.ServiceRegister(input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to register user")
		return
	}

	response := ResponseRegister{
		Name:  user.Name,
		Email: user.Email,
	}

	utils.RespondSuccess(c, response, "User registered successfully")
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid login payload")
		return
	}

	token, err := h.Service.ServiceLogin(input)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	response := ResponseLogin{
		Token: token,
	}

	utils.RespondSuccess(c, response, "Login successful")
}
