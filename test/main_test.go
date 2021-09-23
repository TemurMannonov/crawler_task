package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"crawler_task/api"
	"crawler_task/config"
	"crawler_task/pkg/logger"

	"github.com/gin-gonic/gin"
)

type header struct {
	Key   string
	Value string
}

var (
	server *gin.Engine
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	logger := logger.New(cfg.LogLevel, "crawler task test")

	server = api.New(api.Config{
		Cfg:    cfg,
		Logger: logger,
	})

	os.Exit(m.Run())
}

func PerformRequest(method, path string, req, res interface{}, headers ...header) (*httptest.ResponseRecorder, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	for _, h := range headers {
		request.Header.Add(h.Key, h.Value)
	}

	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return response, nil
}
