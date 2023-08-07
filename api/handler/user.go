package handler

import (
	"net/http"

	"github.com/Shakhrik/video_task/api/models"
	"github.com/Shakhrik/video_task/dto"
	"github.com/Shakhrik/video_task/pkg/logger"
	"github.com/Shakhrik/video_task/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	log logger.Logger
	// cfg         config.Config
	svc service.Service
}

func NewUserHandler(svc service.Service, l logger.Logger) *UserHandler {
	return &UserHandler{svc: svc, log: l}
}

// //@Router /v1/user [post]
// //@Summary Create user
// //@Description API for create user
// //@Tags user
// //@Accept json
// //@Produce json
// //@Param employee body models.BusStopCreate true "user"
// //@Success 200 {object} models.SuccessModel
// //@Failure 400 {object} models.ResponseError
// //@Failure 500 {object} models.ResponseError
func (h UserHandler) CreateUser(c *gin.Context) {
	var u models.CreateUser

	err := c.ShouldBind(&u)
	if err != nil {
		handleErrorResponse(c, h.log, http.StatusBadRequest, "bad request", err)
		return
	}

	req := &dto.CreateUser{Name: u.Name, Email: u.Email}
	res, err := h.svc.UserService().CreateUser(req)
	if err != nil {
		handleErrorResponse(c, h.log, http.StatusInternalServerError, "failed to create user", err)
		return
	}
	c.JSON(http.StatusCreated, models.UserID{ID: res.ID})
}
