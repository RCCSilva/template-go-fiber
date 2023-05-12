package domain

type ResultError struct {
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	Fields     []*InvalidFields `json:"fields,omitempty"`
	error
}
