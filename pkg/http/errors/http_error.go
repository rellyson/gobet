package errors

type ErrHttpBase struct {
	Error      string `json:"error"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"status_code"`
}
