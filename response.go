package main

type Data struct {
	Status bool `json:"status"`
	Data []Person `json:"data,omitempty"`
	Message string `json:"message"`
}