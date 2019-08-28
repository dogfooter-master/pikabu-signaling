package grpc

import (
	"context"
	"pikabu-signaling/signaling/pkg/endpoint"
	"pikabu-signaling/signaling/pkg/grpc/pb"

	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makeRegisterHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)
}

func decodeRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RegisterRequest)
	return endpoint.RegisterRequest{
		Req: req,
	}, nil
}

func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.RegisterResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) Register(ctx context1.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	_, rep, err := g.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterReply), nil
}

func makeWebRTCStartPleaseHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.WebRTCStartPleaseEndpoint, decodeWebRTCStartPleaseRequest, encodeWebRTCStartPleaseResponse, options...)
}

func decodeWebRTCStartPleaseRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.WebRTCStartPleaseRequest)
	return endpoint.WebRTCStartPleaseRequest{
		Req: req,
	}, nil
}

func encodeWebRTCStartPleaseResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.WebRTCStartPleaseResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) WebRTCStartPlease(ctx context1.Context, req *pb.WebRTCStartPleaseRequest) (*pb.WebRTCStartPleaseReply, error) {
	_, rep, err := g.webRTCStartPlease.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebRTCStartPleaseReply), nil
}

func makeWebRTCAnswerHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.WebRTCAnswerEndpoint, decodeWebRTCAnswerRequest, encodeWebRTCAnswerResponse, options...)
}

func decodeWebRTCAnswerRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.WebRTCAnswerRequest)
	return endpoint.WebRTCAnswerRequest{
		Req: req,
	}, nil
}

func encodeWebRTCAnswerResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.WebRTCAnswerResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) WebRTCAnswer(ctx context1.Context, req *pb.WebRTCAnswerRequest) (*pb.WebRTCAnswerReply, error) {
	_, rep, err := g.webRTCAnswer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebRTCAnswerReply), nil
}

func makeWebRTCAddCandidateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.WebRTCAddCandidateEndpoint, decodeWebRTCAddCandidateRequest, encodeWebRTCAddCandidateResponse, options...)
}

func decodeWebRTCAddCandidateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.WebRTCAddCandidateRequest)
	return endpoint.WebRTCAddCandidateRequest{
		Req: req,
	}, nil
}

func encodeWebRTCAddCandidateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.WebRTCAddCandidateResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) WebRTCAddCandidate(ctx context1.Context, req *pb.WebRTCAddCandidateRequest) (*pb.WebRTCAddCandidateReply, error) {
	_, rep, err := g.webRTCAddCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WebRTCAddCandidateReply), nil
}
