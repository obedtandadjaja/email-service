package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type SendRequest struct {
	Sender        string   `json:"sender"`
	Recipients    []string `json:"recipient"`
	Subject       string   `json:"subject"`
	Body          string   `json:"body"`
	BccRecipients []string `json:"body"`
}

type SendResponse struct {
	Message string `json:"message"`
	EmailId string `json:"emailId"`
}

func parseSendRequest(r *http.Request) (*SendRequest, error) {
	var request SendRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	return &request, err
}

func (server Server) Send(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request, err := parseSendRequest(r)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	message := server.mailgun.NewMessage(request.Sender, request.Subject, request.Body, request.Recipients...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	msg, id, err := server.mailgun.Send(ctx, message)
	if err != nil {
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	}

	response := SendResponse{
		Message: msg,
		EmailId: id,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func mailgunRoutes(server Server) {
}
