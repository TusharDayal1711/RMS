package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// struct of ClientError
type ClientError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Timestamp  string `json:"timestamp"`
}

func RespondError(w http.ResponseWriter, statusCode int, err error, userMessage string) {
	logrus.Errorf("status: %d, user_message: %s, internal_error: %+v", statusCode, userMessage, err)
	clientError := ClientError{
		Error:      http.StatusText(statusCode),
		Message:    userMessage,
		StatusCode: statusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(clientError); err != nil {
		logrus.Errorf("failed to encode/send error response: %+v", err)
	}
}
