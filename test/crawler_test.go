package test

import (
	"crawler_task/api/models"
	"net/http"
	"testing"

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
