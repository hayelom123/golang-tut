package model

type ApiError struct{
	Error string `json:"error"`
	Status int   `json:"status"`
}