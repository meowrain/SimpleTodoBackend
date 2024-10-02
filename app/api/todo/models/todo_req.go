package models

type TodoRequest struct {
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	Tag       string `json:"tag"`
}
