package service

import (
	"context"
	"errors"
	"fmt"
	"pikabu-signaling/signaling/pkg/grpc/pb"

	"gopkg.in/mgo.v2/bson"
)

// SignalingService describes the service.
type SignalingService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterReply, err error)
	WebRTCStartPlease(ctx context.Context, req *pb.WebRTCStartPleaseRequest) (res *pb.WebRTCStartPleaseReply, err error)
	WebRTCAnswer(ctx context.Context, req *pb.WebRTCAnswerRequest) (res *pb.WebRTCAnswerReply, err error)
	WebRTCAddCandidate(ctx context.Context, req *pb.WebRTCAddCandidateRequest) (res *pb.WebRTCAddCandidateReply, err error)
}

type basicSignalingService struct{}

func (b *basicSignalingService) Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterReply, err error) {

	fmt.Println("req:", req)
	// TODO: Register

	res = &pb.RegisterReply{
		Account: req.Account,
	}
	return res, err
}

func (b *basicSignalingService) WebRTCStartPlease(ctx context.Context, req *pb.WebRTCStartPleaseRequest) (res *pb.WebRTCStartPleaseReply, err error) {

	webRTC := WebRTC{}
	webRTC.CreateDataChannel()
	webRTC.CreateOffer()

	if WebRTCMap == nil {
		WebRTCMap = make(map[string]WebRTC)
	}
	connId := bson.NewObjectId().Hex()
	WebRTCMap[connId] = webRTC

	res = &pb.WebRTCStartPleaseReply{
		Offer:  webRTC.Offer,
		ConnId: connId,
	}
	return res, err
}

func (b *basicSignalingService) WebRTCAnswer(ctx context.Context, req *pb.WebRTCAnswerRequest) (res *pb.WebRTCAnswerReply, err error) {

	fmt.Println("req:", req)

	if WebRTCMap == nil {
		err = errors.New("no connection id")
		return
	}
	if _, ok := WebRTCMap[req.ConnId]; !ok {
		err = errors.New("no connection id")
		return
	}
	webRTC := WebRTCMap[req.ConnId]
	webRTC.ReceiveAnswer(req.Answer)

	res = &pb.WebRTCAnswerReply{
		Account: req.Account,
	}

	return res, err
}

func (b *basicSignalingService) WebRTCAddCandidate(ctx context.Context, req *pb.WebRTCAddCandidateRequest) (res *pb.WebRTCAddCandidateReply, err error) {

	if WebRTCMap == nil {
		err = errors.New("no connection id")
		return
	}
	if _, ok := WebRTCMap[req.ConnId]; !ok {
		err = errors.New("no connection id")
		return
	}
	webRTC := WebRTCMap[req.ConnId]
	webRTC.AddCandidate(req.Candidate)

	res = &pb.WebRTCAddCandidateReply{
		Account: req.Account,
	}

	return res, err
}


// NewBasicSignalingService returns a naive, stateless implementation of SignalingService.
func NewBasicSignalingService() SignalingService {
	return &basicSignalingService{}
}

// New returns a SignalingService with all of the expected middleware wired in.
func New(middleware []Middleware) SignalingService {
	var svc SignalingService = NewBasicSignalingService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}