package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"e-commerce/domain"
	"e-commerce/entities"
)

type Handler struct {
	userUseCase domain.UserUsecase
}

func NewHandler(r *gin.RouterGroup, userUsecase domain.UserUsecase) *Handler {
	h := &Handler{userUseCase: userUsecase}

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)

	return h
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req entities.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warn(errors.Wrap(err, "[Handler.CreateUser]: failed to bind request body"))
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if err := h.userUseCase.CreateUser(req); err != nil {
		log.Error(errors.Wrap(err, "[Handler.CreateUser]: failed to create user"))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Warn(errors.Wrap(err, "[Handler.GetUser]: failed to parse id"))
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
		return
	}

	user, err := h.userUseCase.GetUser(uint32(id))
	if err != nil {
		log.Error(errors.Wrap(err, "[Handler.GetUser]: failed to get user"))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
