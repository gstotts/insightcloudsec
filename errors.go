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
