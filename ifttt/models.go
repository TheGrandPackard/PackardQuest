package ifttt

type responseObject struct {
	Errors []errorObject `json:"errors"`
}

type errorObject struct {
	Message string `json:"message"`
}
