package endpoint

import (
	"context"
	"pikabu-signaling/signaling/pkg/grpc/pb"
	service "pikabu-signaling/signaling/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Req *pb.RegisterRequest `json:"req"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	Res *pb.RegisterReply `json:"res"`
	Err error             `json:"err"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.SignalingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		res, err := s.Register(ctx, req.Req)
		return RegisterResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterReply, err error) {
	request := RegisterRequest{Req: req}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).Res, response.(RegisterResponse).Err
}

// WebRTCStartPleaseRequest collects the request parameters for the WebRTCStartPlease method.
type WebRTCStartPleaseRequest struct {
	Req *pb.WebRTCStartPleaseRequest `json:"req"`
}

// WebRTCStartPleaseResponse collects the response parameters for the WebRTCStartPlease method.
type WebRTCStartPleaseResponse struct {
	Res *pb.WebRTCStartPleaseReply `json:"res"`
	Err error                      `json:"err"`
}

// MakeWebRTCStartPleaseEndpoint returns an endpoint that invokes WebRTCStartPlease on the service.
func MakeWebRTCStartPleaseEndpoint(s service.SignalingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WebRTCStartPleaseRequest)
		res, err := s.WebRTCStartPlease(ctx, req.Req)
		return WebRTCStartPleaseResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r WebRTCStartPleaseResponse) Failed() error {
	return r.Err
}

// WebRTCStartPlease implements Service. Primarily useful in a client.
func (e Endpoints) WebRTCStartPlease(ctx context.Context, req *pb.WebRTCStartPleaseRequest) (res *pb.WebRTCStartPleaseReply, err error) {
	request := WebRTCStartPleaseRequest{Req: req}
	response, err := e.WebRTCStartPleaseEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WebRTCStartPleaseResponse).Res, response.(WebRTCStartPleaseResponse).Err
}

// WebRTCAnswerRequest collects the request parameters for the WebRTCAnswer method.
type WebRTCAnswerRequest struct {
	Req *pb.WebRTCAnswerRequest `json:"req"`
}

// WebRTCAnswerResponse collects the response parameters for the WebRTCAnswer method.
type WebRTCAnswerResponse struct {
	Res *pb.WebRTCAnswerReply `json:"res"`
	Err error                 `json:"err"`
}

// MakeWebRTCAnswerEndpoint returns an endpoint that invokes WebRTCAnswer on the service.
func MakeWebRTCAnswerEndpoint(s service.SignalingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WebRTCAnswerRequest)
		res, err := s.WebRTCAnswer(ctx, req.Req)
		return WebRTCAnswerResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r WebRTCAnswerResponse) Failed() error {
	return r.Err
}

// WebRTCAnswer implements Service. Primarily useful in a client.
func (e Endpoints) WebRTCAnswer(ctx context.Context, req *pb.WebRTCAnswerRequest) (res *pb.WebRTCAnswerReply, err error) {
	request := WebRTCAnswerRequest{Req: req}
	response, err := e.WebRTCAnswerEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WebRTCAnswerResponse).Res, response.(WebRTCAnswerResponse).Err
}

// WebRTCAddCandidateRequest collects the request parameters for the WebRTCAddCandidate method.
type WebRTCAddCandidateRequest struct {
	Req *pb.WebRTCAddCandidateRequest `json:"req"`
}

// WebRTCAddCandidateResponse collects the response parameters for the WebRTCAddCandidate method.
type WebRTCAddCandidateResponse struct {
	Res *pb.WebRTCAddCandidateReply `json:"res"`
	Err error                       `json:"err"`
}

// MakeWebRTCAddCandidateEndpoint returns an endpoint that invokes WebRTCAddCandidate on the service.
func MakeWebRTCAddCandidateEndpoint(s service.SignalingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WebRTCAddCandidateRequest)
		res, err := s.WebRTCAddCandidate(ctx, req.Req)
		return WebRTCAddCandidateResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r WebRTCAddCandidateResponse) Failed() error {
	return r.Err
}

// WebRTCAddCandidate implements Service. Primarily useful in a client.
func (e Endpoints) WebRTCAddCandidate(ctx context.Context, req *pb.WebRTCAddCandidateRequest) (res *pb.WebRTCAddCandidateReply, err error) {
	request := WebRTCAddCandidateRequest{Req: req}
	response, err := e.WebRTCAddCandidateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WebRTCAddCandidateResponse).Res, response.(WebRTCAddCandidateResponse).Err
}
