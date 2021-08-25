package insightcloudsec

import "fmt"

// Error messages
const (
	errMarshallingParams = "error marshalling params to JSON"
)

// APIRequestError is a type of error raised by API calls made from this library
type APIRequestError struct {
	StatusCode int
	Message    string
}

func (e APIRequestError) Error() string {
	return fmt.Sprintf("\nHTTP Status: %d: %s\n", e.StatusCode, e.Message)
}
