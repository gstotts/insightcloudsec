package insightcloudsec

import (
	"fmt"
	"net/http"
)

// APIRequestError is a type of error raised by API calls made from this library
type APIRequestError struct {
	Request    http.Request
	StatusCode int
	Message    string
}

func (e APIRequestError) Error() string {
	return fmt.Sprintf("\nRequested URL: %s\nHTTP Status: %d: %s\n", e.Request.URL, e.StatusCode, e.Message)
}

// ValidationError is a type of error raised when validation of a given string is not of an expected/required value
type ValidationError struct {
	ItemToValidate string
	ExpectedValues []string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("\n Validation Error:  %s should be one of %s", e.ItemToValidate, e.ExpectedValues)
}
