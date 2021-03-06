package plaid

import (
	"fmt"
)

type PlaidError struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"error_message"`
	ErrorType string `json:"error_type"`

	// StatusCode needs to manually set from the http response
	StatusCode int
}

func (e PlaidError) Error() string {
	return fmt.Sprintf("Plaid Error - http status: %d, code: %s, message: %s, type: %s",
		e.StatusCode, e.ErrorCode, e.Message, e.ErrorType)
}
