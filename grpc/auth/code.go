package grpcauth

import (
	"context"
	"time"

	"github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetConfirmCode(req *sso.GetConfirmCodeRequest, stream sso.Auth_GetConfirmCodeServer) (err error) {
	qrcodes, err := s.confCodes.GetTeachersConfirmCodes(req.GetTeacherLogin())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	funcNow := time.Now()

	for _, code := range qrcodes {
		for time.Since(funcNow).Minutes() <= 2 {
			if err = stream.Send(&sso.GetConfirmCodeResponse{Code: code.Code}); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}

func (s *serverAPI) SubmitRoomCode(ctx context.Context, req *sso.SubmitCodeRequest) (res *sso.SubmitCodeResponse, err error) {
	if err := s.confCodes.SubmitRoomCode(ctx, req.Code); err != nil {
		return &sso.SubmitCodeResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}
	return &sso.SubmitCodeResponse{Success: true}, nil
}

func (s *serverAPI) SubmitTeacherCode(ctx context.Context, req *sso.SubmitCodeRequest) (res *sso.SubmitCodeResponse, err error) {
	if err := s.confCodes.SubmitTeacherCode(ctx, req.Code); err != nil {
		return &sso.SubmitCodeResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}
	return &sso.SubmitCodeResponse{Success: true}, nil
}
