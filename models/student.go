package models

type Student struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullName"`
	City     string `json:"city"`
}
