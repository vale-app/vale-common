package models

type Pagination struct {
	Page           int `json:"page"`
	ResultsPerPage int `json:"limit"`
}
