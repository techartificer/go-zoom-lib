package zoom

import (
	"encoding/json"
	"fmt"
)

// APIError contains the code and message returned by any Zoom errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

type errorContainer struct {
	*APIError
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("Zoom API error %d: \"%s\"", e.Code, e.Message)
}

func checkError(body []byte) error {
	c := errorContainer{}
	if err := json.Unmarshal(body, &c); err != nil {
		return err
	}

	// need to explicitly return nil or it will register as an error
	if c.APIError == nil {
		return nil
	}

	return c.APIError
}
