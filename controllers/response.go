package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Responder is the expected behaviour from a http response struct
type Responder interface {
	WriteResponse(w http.ResponseWriter) error
}

// jsonResponse holds our json payload for marshalling
type jsonResponse struct {
	statusCode int
	Data       interface{} `json:"Data,omitempty"`
	Message    string      `json:"Message,omitempty"`
}

// NewJSONErrorResponse is a constructor for an error json Response
func NewJSONSuccessResponse(statusCode int, data interface{}) jsonResponse {
	return jsonResponse{
		statusCode: statusCode,
		Data:       data,
	}
}

// NewJSONErrorResponse is a constructor for an error json Response
func NewJSONErrorResponse(statusCode int, message string) jsonResponse {
	return jsonResponse{
		statusCode: statusCode,
		Message:    message,
	}
}

// WriteResponse writes the structs payload to the Response
func (j *jsonResponse) WriteResponse(w http.ResponseWriter) error {
	// Sets json header
	w.Header().Add("Content-Type", "application/json")

	// writes the jsonResponse struct into json bytes
	b, err := json.Marshal(j)
	if err != nil {
		return fmt.Errorf("could not marshal jsonResponse %w", err)
	}

	// Write status code for response
	w.WriteHeader(j.statusCode)

	// write the payload to the response
	_, err = w.Write(b)
	if err != nil {
		return fmt.Errorf("could not write jsonResponse %w", err)
	}

	return nil
}
