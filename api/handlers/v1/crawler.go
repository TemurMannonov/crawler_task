package v1

import (
	"errors"
	"fmt"
	"net/http"

	"crawler_task/api/models"
	workerpool "crawler_task/pkg/worker_pool"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Url   string
	Title string
	Error error
}

//@Security ApiKeyAuth
//@Router /v1/crawler [post]
//@Summary Crawler
//@Description API for requesting crawler
//@Tags crawler
//@Accept json
//@Produce json
//@Param input body models.CrawlerRequest true "input"
//@Success 201 {object} models.CrawlerResponse
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h *handlerV1) Crawler(c *gin.Context) {
	var (
		urls []string = []string{
			"https://www.result.si/projekti/",
			"https://www.result.si/o-nas/",
			"https://www.result.si/kariera/",
			"https://www.result.si/blog/",
		}
		request  models.CrawlerRequest
		response models.CrawlerResponse
	)

	h.log.Info("Crawler request")

	err := c.ShouldBind(&request)
	if HandleError(c, h.log, err, "Error while binding request to model") {
		return
	}

	if request.Workers < 1 || request.Workers > 4 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Workers count should be between 1 and 4.",
		})
	}

	wp := workerpool.NewWorkerPool(request.Workers)
	wp.Run()

	resultC := make(chan Result, len(urls))
	for _, val := range urls {
		url := val
		wp.AddTask(func() {
			title, err := getTitleFromUrl(url)
			resultC <- Result{url, title, err}
		})
	}

	for i := 0; i < len(urls); i++ {
		res := <-resultC
		if res.Error == nil {
			response.SuccessfullCalls++
		} else {
			response.FailedCalls++
		}

		response.Results = append(response.Results, models.Result{
			Title: res.Title,
			Url:   res.Url,
		})
	}

	c.JSON(http.StatusOK, response)
}

func getTitleFromUrl(url string) (string, error) {
	var elementClass = ".et_pb_module_header"

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		errStr := fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)
		return "", errors.New(errStr)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	title := doc.Find(elementClass).First().Text()
	return title, nil
}
