package handler

import (
	"github.com/Shakhrik/video_task/api/models"
	"github.com/Shakhrik/video_task/pkg/logger"
	"github.com/gin-gonic/gin"
)

func handleErrorResponse(c *gin.Context, l logger.Logger, code int, message string, err error) {
	l.Error(message, logger.Int("code", code), logger.Error(err))
	c.JSON(code, models.ErrorModel{Error: message})
}

func handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.SuccessModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
