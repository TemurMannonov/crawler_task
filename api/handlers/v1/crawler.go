package v1

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/TemurMannonov/crawler_task/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Data struct {
	Url string `json:"url"`
}

var data2 map[string][]models.Result = make(map[string][]models.Result)

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
		request                       models.CrawlerRequest
		successfullCalls, failedCalls int
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
		return
	}

	id, _ := uuid.NewRandom()
	requestID := id.String()

	jobsCount := len(urls)

	jobs := make(chan Data, jobsCount)
	errors := make(chan error, jobsCount)

	for k := 1; k <= request.Workers; k++ {
		go worker(jobs, errors, requestID)
	}

	for _, v := range urls {
		jobs <- Data{v}
	}
	close(jobs)

	for i := 0; i < jobsCount; i++ {
		err := <-errors
		if err == nil {
			successfullCalls++
		} else {
			failedCalls++
		}
	}
	close(errors)

	c.JSON(http.StatusOK, models.CrawlerResponse{
		SuccessfullCalls: successfullCalls,
		FailedCalls:      failedCalls,
		Results:          data2[requestID],
	})
}

func worker(jobs <-chan Data, errors chan<- error, id string) {
	var mu sync.Mutex

	for v := range jobs {
		title, err := getTitleFromUrl(v.Url)

		mu.Lock()

		data2[id] = append(data2[id], models.Result{
			Title: title,
			Url:   v.Url,
		})

		mu.Unlock()

		errors <- err
	}
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
