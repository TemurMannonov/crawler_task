package models

type CrawlerRequest struct {
	Workers int `json:"workers"`
}

type Result struct {
	Url   string
	Title string
}

type Result3 struct {
	Title        string
	FailCount    uint64
	SuccessCount uint64
}

type CrawlerResponse struct {
	SuccessfullCalls int      `json:"successfull_calls"`
	FailedCalls      int      `json:"failed_calls"`
	Results          []Result `json:"results"`
}
