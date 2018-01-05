package plaid

import (
	"fmt"
)

type plaidError struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"error_message"`
	ErrorType string `json:"error_type"`

	// StatusCode needs to manually set from the http response
	StatusCode int
}

func (e plaidError) Error() string {
	return fmt.Sprintf("Plaid Error - http status: %d, code: %d, message: %s, type: %s",
		e.StatusCode, e.ErrorCode, e.Message, e.ErrorType)
}
