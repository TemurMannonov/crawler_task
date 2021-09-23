package models

type ResponseOK struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
