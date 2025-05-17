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

// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body RegisterRequest true "Register Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
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

// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
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
