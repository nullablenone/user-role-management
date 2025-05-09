package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func (h *Handler) GetRolesByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.Service.ServiceGetRolesByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"role": role,
	})

}

func (h *Handler) CreateRoles(c *gin.Context) {
	var input CreateRolesRequest

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role, err := h.Service.ServiceCreateRoles(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Role created successfully",
		"role":    role,
	})
}

func (h *Handler) UpdateRoles(c *gin.Context) {
	id := c.Param("id")
	var input UpdateRolesRequest
	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role, err := h.Service.ServiceUpdateRoles(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role Update successfully",
		"role":    role,
	})
}

func (h *Handler) DeleteRoles(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.ServiceDeleteRoles(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Delete successfully",
	})
}
