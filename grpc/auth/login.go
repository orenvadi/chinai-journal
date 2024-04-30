package grpcauth

import (
	"context"
	"log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) LoginTeacher(
	ctx context.Context,
	req *sso.LoginTeacherRequest,
) (
	res *sso.LoginTeacherResponse,
	err error,
) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	// validating
	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accessToken, err := s.auth.LoginTeacher(ctx, req.GetTeacherLogin(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.LoginTeacherResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *serverAPI) LoginStudent(
	ctx context.Context,
	req *sso.LoginStudentRequest,
) (
	res *sso.LoginStudentResponse,
	err error,
) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	// validating
	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accessToken, err := s.auth.LoginStudent(ctx, req.GetStudentLogin(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.LoginStudentResponse{
		AccessToken: accessToken,
	}, nil
}
