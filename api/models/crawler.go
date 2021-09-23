package models

type CrawlerRequest struct {
	Workers int `json:"workers"`
}

type Result struct {
	Url   string
	Title string
}

type CrawlerResponse struct {
	SuccessfullCalls int      `json:"successfull_calls"`
	FailedCalls      int      `json:"failed_calls"`
	Results          []Result `json:"results"`
}
