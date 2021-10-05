package helper

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Response is general api response structure
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// SuccessResponse is here for present success responses such: 200, 201, ...
func SuccessResponse(w http.ResponseWriter, message string, data interface{}, status bool, code int) {
	jsonResp, _ := json.Marshal(Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	})

	// Set headers and send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}

// FailureResponse is here for present failure responses such: 500, 400, 404, ...
func FailureResponse(w http.ResponseWriter, message string, err string, data interface{}, code int) {
	var resp Response

	// Hide 500 error messages from client
	if code == http.StatusInternalServerError {
		resp = Response{
			Status:  false,
			Message: message,
			Errors:  "an error has occurred",
			Data:    data,
		}
	} else {
		// Split error message to []string
		se := strings.Split(err, "\n")
		resp = Response{
			Status:  false,
			Message: message,
			Errors:  se,
			Data:    data,
		}
	}

	jsonResp, _ := json.Marshal(resp)

	// Set headers and send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}
