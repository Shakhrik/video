package api

import (
	"github.com/Shakhrik/video_task/pkg/logger"
	"github.com/Shakhrik/video_task/service"

	"github.com/Shakhrik/video_task/api/handler"

	_ "github.com/Shakhrik/video_task/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log logger.Logger
	// Cfg     *config.Config
	Service service.Service
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	h := handler.NewHandler(opt.Service, opt.Log)

	apiV1 := router.Group("/v1")

	{
		apiV1.POST("/destination", h.CreateUser)

	}
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
