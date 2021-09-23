package test

import (
	"net/http"
	"testing"

	"github.com/TemurMannonov/crawler_task/api/models"

	"github.com/stretchr/testify/assert"
)

func TestCrawler(t *testing.T) {
	var testSet = []struct {
		nameOfTest string
		routerURL  string
		methodType string
		code       int
		equal      bool
		request    interface{}
		response   interface{}
		want       interface{}
	}{
		{
			nameOfTest: "crawler",
			routerURL:  "/v1/crawler",
			methodType: http.MethodPost,
			code:       200,
			response:   &models.CrawlerResponse{},
			want:       &models.CrawlerResponse{},
			equal:      false,
			request: models.CrawlerRequest{
				Workers: 3,
			},
		},
		{
			nameOfTest: "crawler",
			routerURL:  "/v1/crawler",
			methodType: http.MethodPost,
			code:       400,
			response:   &models.ResponseError{},
			want: &models.ResponseError{
				Code:    400,
				Message: "Workers count should be between 1 and 4.",
			},
			equal: true,
			request: models.CrawlerRequest{
				Workers: 0,
			},
		},
		{
			nameOfTest: "crawler",
			routerURL:  "/v1/crawler",
			methodType: http.MethodPost,
			code:       400,
			response:   &models.ResponseError{},
			want: &models.ResponseError{
				Code:    400,
				Message: "Workers count should be between 1 and 4.",
			},
			equal: true,
			request: models.CrawlerRequest{
				Workers: 5,
			},
		},
	}

	for i := range testSet {
		test := testSet[i]
		t.Run(test.nameOfTest, func(t *testing.T) {
			resp, err := PerformRequest(test.methodType, test.routerURL, test.request, test.response)
			assert.NoError(t, err)
			assert.NotEmpty(t, resp)
			assert.Equal(t, test.code, resp.Code)

			if test.equal {
				assert.Equal(t, test.want, test.response)
			} else {
				assert.NotEqual(t, test.want, test.response)
			}

		})
	}

}
