package service

import (
	"context"
	"pikabu-signaling/signaling/pkg/grpc/pb"

	log "github.com/go-kit/kit/log"
)

type Middleware func(SignalingService) SignalingService

type loggingMiddleware struct {
	logger log.Logger
	next   SignalingService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SignalingService) SignalingService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterReply, err error) {
	defer func() {
		l.logger.Log("method", "Register", "req", req, "res", res, "err", err)
	}()
	return l.next.Register(ctx, req)
}

func (l loggingMiddleware) WebRTCStartPlease(ctx context.Context, req *pb.WebRTCStartPleaseRequest) (res *pb.WebRTCStartPleaseReply, err error) {
	defer func() {
		l.logger.Log("method", "WebRTCStartPlease", "req", req, "res", res, "err", err)
	}()
	return l.next.WebRTCStartPlease(ctx, req)
}

func (l loggingMiddleware) WebRTCAnswer(ctx context.Context, req *pb.WebRTCAnswerRequest) (res *pb.WebRTCAnswerReply, err error) {
	defer func() {
		l.logger.Log("method", "WebRTCAnswer", "req", req, "res", res, "err", err)
	}()
	return l.next.WebRTCAnswer(ctx, req)
}

func (l loggingMiddleware) WebRTCAddCandidate(ctx context.Context, req *pb.WebRTCAddCandidateRequest) (res *pb.WebRTCAddCandidateReply, err error) {
	defer func() {
		l.logger.Log("method", "WebRTCAddCandidate", "req", req, "res", res, "err", err)
	}()
	return l.next.WebRTCAddCandidate(ctx, req)
}
