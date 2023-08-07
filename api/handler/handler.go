package handler

import (
	"github.com/Shakhrik/video_task/pkg/logger"
	"github.com/Shakhrik/video_task/service"
)

type Handler struct {
	UserHandler
}

func NewHandler(svc service.Service, log logger.Logger) *Handler {
	return &Handler{
		UserHandler: *NewUserHandler(svc, log),
	}
}
