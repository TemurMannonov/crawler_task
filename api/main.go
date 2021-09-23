package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/TemurMannonov/crawler_task/api/docs"
	v1 "github.com/TemurMannonov/crawler_task/api/handlers/v1"
	"github.com/TemurMannonov/crawler_task/config"

	"github.com/TemurMannonov/crawler_task/pkg/logger"
)

// Config ...
type Config struct {
	Logger logger.Logger
	Cfg    config.Config
}

func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger: cnf.Logger,
		Cfg:    cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "OK"})
	})

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/v1/crawler", handlerV1.Crawler)

	return r
}
