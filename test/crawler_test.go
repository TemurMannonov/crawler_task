package test

import (
	"net/http"
	"testing"

	"github.com/TemurMannonov/crawler_task/api/models"

	"github.com/stretchr/testify/assert"
)

func TestCrawler(t *testing.T) {
	var (
		req models.CrawlerRequest
		res models.CrawlerResponse
	)

	req.Workers = 1
	resp, err := PerformRequest(http.MethodPost, "/v1/crawler", req, res)
	assert.NoError(t, err)
	assert.Equal(t, resp.Code, 200)
}
