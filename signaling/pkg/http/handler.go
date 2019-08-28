package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"
	endpoint "pikabu-signaling/signaling/pkg/endpoint"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/register").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)))
}

// decodeRegisterRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeWebRTCStartPleaseHandler creates the handler logic
func makeWebRTCStartPleaseHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/web-rtcstart-please").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.WebRTCStartPleaseEndpoint, decodeWebRTCStartPleaseRequest, encodeWebRTCStartPleaseResponse, options...)))
}

// decodeWebRTCStartPleaseRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeWebRTCStartPleaseRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.WebRTCStartPleaseRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeWebRTCStartPleaseResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeWebRTCStartPleaseResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeWebRTCAnswerHandler creates the handler logic
func makeWebRTCAnswerHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/web-rtcanswer").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.WebRTCAnswerEndpoint, decodeWebRTCAnswerRequest, encodeWebRTCAnswerResponse, options...)))
}

// decodeWebRTCAnswerRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeWebRTCAnswerRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.WebRTCAnswerRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeWebRTCAnswerResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeWebRTCAnswerResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeWebRTCAddCandidateHandler creates the handler logic
func makeWebRTCAddCandidateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/web-rtcadd-candidate").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.WebRTCAddCandidateEndpoint, decodeWebRTCAddCandidateRequest, encodeWebRTCAddCandidateResponse, options...)))
}

// decodeWebRTCAddCandidateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeWebRTCAddCandidateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.WebRTCAddCandidateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeWebRTCAddCandidateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeWebRTCAddCandidateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
