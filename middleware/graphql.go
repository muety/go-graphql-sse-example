package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"log"
	"net/http"
	"strings"
	"time"
)

type GraphQL struct {
	Schema *graphql.Schema
}

const subscriptionTimeout = 30 * time.Minute

type params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params params

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: string matching is probably expensive
	switch strings.HasPrefix(params.Query, "subscription") {
	case true:
		h.serveSubscription(w, r, &params)
		break
	default:
		h.serveDefault(w, r, &params)
	}
}

func (h *GraphQL) serveDefault(w http.ResponseWriter, r *http.Request, params *params) {
	response := h.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, statusCode := checkError(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

func (h *GraphQL) serveSubscription(w http.ResponseWriter, r *http.Request, params *params) {
	ctx, cancel := context.WithTimeout(r.Context(), subscriptionTimeout)

	// TODO: use context instead, as CloseNotifier is deprecated
	// However, for some reason, context doesn't seem to get cancelled when client aborts request
	callOnClose(w, cancel)

	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Printf("error: not a flusher\n")
	}

	c, err := h.Schema.Subscribe(ctx, params.Query, params.OperationName, params.Variables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/event-stream")
	w.Header().Set("cache-control", "no-cache")

	var eventCount int
	for r := range c {
		response := r.(*graphql.Response)
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if hasError, statusCode := checkError(response); hasError {
			w.WriteHeader(statusCode)
		}

		eventCount++
		// TODO: make json decoder write to stream, don't buffer here
		fmt.Fprintf(w, "data: %s\n\n", responseJSON)
		flusher.Flush()
	}

	if eventCount == 0 {
		w.WriteHeader(http.StatusOK)
	}
}

func checkError(response *graphql.Response) (hasError bool, statusCode int) {
	statusCode = http.StatusOK

	if len(response.Errors) > 0 {
		hasError = true

		for _, e := range response.Errors {
			if e.Message == "unauthorized" {
				statusCode = http.StatusUnauthorized
				break
			}
		}
	}

	return hasError, statusCode
}

func callOnClose(w http.ResponseWriter, cb func()) {
	if cn, ok := w.(http.CloseNotifier); !ok {
		log.Printf("error: not a close notifier\n")
	} else {
		// https://github.com/grpc-ecosystem/grpc-gateway/pull/120
		closeChan := cn.CloseNotify()
		go func(closeChan <-chan bool, cn http.CloseNotifier, cb func()) {
			<-closeChan
			cb()
		}(closeChan, cn, cb)
	}
}
