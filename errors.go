package chainhooks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HttpError represents an HTTP error response from the Chainhooks API.
type HttpError struct {
	StatusCode int
	URL        string
	Method     string
	Headers    http.Header
	Body       string
	RawBody    []byte
	Err        error
}

// Error implements the error interface.
func (e *HttpError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %s: status %d: %v", e.Method, e.URL, e.StatusCode, e.Err)
	}
	if e.Body != "" {
		return fmt.Sprintf("%s %s: status %d: %s", e.Method, e.URL, e.StatusCode, e.Body)
	}
	return fmt.Sprintf("%s %s: status %d", e.Method, e.URL, e.StatusCode)
}

// Unwrap returns the underlying error.
func (e *HttpError) Unwrap() error {
	return e.Err
}

// MarshalJSON returns a JSON representation of the error.
func (e *HttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"error":       e.Error(),
		"status_code": e.StatusCode,
		"url":         e.URL,
		"method":      e.Method,
		"body":        e.Body,
	})
}

// newHttpError creates a new HttpError from an HTTP response.
func newHttpError(resp *http.Response, req *http.Request) *HttpError {
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	errMsg := ""
	if len(body) > 0 {
		// Try to parse as JSON error response
		var errResp map[string]interface{}
		if err := json.Unmarshal(body, &errResp); err == nil {
			if message, ok := errResp["message"].(string); ok {
				errMsg = message
			} else if message, ok := errResp["error"].(string); ok {
				errMsg = message
			} else {
				errMsg = string(body)
			}
		} else {
			errMsg = string(body)
		}
	}

	return &HttpError{
		StatusCode: resp.StatusCode,
		URL:        req.URL.String(),
		Method:     req.Method,
		Headers:    resp.Header.Clone(),
		Body:       errMsg,
		RawBody:    body,
	}
}

// ValidationError represents a validation error when building requests.
type ValidationError struct {
	Field  string
	Reason string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Reason)
}

// ConfigError represents a configuration error.
type ConfigError struct {
	Message string
}

// Error implements the error interface.
func (e *ConfigError) Error() string {
	return fmt.Sprintf("config error: %s", e.Message)
}
