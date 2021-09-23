package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TemurMannonov/crawler_task/api/models"
	"github.com/TemurMannonov/crawler_task/config"
	"github.com/TemurMannonov/crawler_task/pkg/logger"
)

type handlerV1 struct {
	log logger.Logger
	cfg config.Config
}

// HandlerV1Config
type HandlerV1Config struct {
	Logger logger.Logger
	Cfg    config.Config
}

// New handler
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log: c.Logger,
		cfg: c.Cfg,
	}
}

func HandleError(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		l.Error(message, logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
		return true
	}
	return false
}
