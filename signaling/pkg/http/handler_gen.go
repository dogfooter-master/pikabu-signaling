// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
	endpoint "pikabu-signaling/signaling/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeRegisterHandler(m, endpoints, options["Register"])
	makeWebRTCStartPleaseHandler(m, endpoints, options["WebRTCStartPlease"])
	makeWebRTCAnswerHandler(m, endpoints, options["WebRTCAnswer"])
	makeWebRTCAddCandidateHandler(m, endpoints, options["WebRTCAddCandidate"])
	return m
}
