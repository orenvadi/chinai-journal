package grpcauth

import (
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
