package grpcauth

import (
	"context"
	"log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/orenvadi/auth-grpc/internal/domain/models"
	"github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) RegisterTeacher(ctx context.Context, req *sso.RegisterTeacherRequest) (*sso.RegisterTeacherResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	// validating
	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	usrName := models.Name{
		First:      req.GetFirstName(),
		Last:       req.GetLastName(),
		Patronimic: req.GetPatronimic(),
	}

	userID, accessToken, refreshToken, err := s.auth.RegisterTeacher(ctx, usrName, req.GetEmail(), req.GetPassword(), req.GetTeacherLogin(), []string{}, []string{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.RegisterTeacherResponse{
		TeacherId:    userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *serverAPI) RegisterStudent(ctx context.Context, req *sso.RegisterStudentRequest) (*sso.RegisterStudentResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	// validating
	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	usrName := models.Name{
		First:      req.GetFirstName(),
		Last:       req.GetLastName(),
		Patronimic: req.GetPatronimic(),
	}

	userID, accessToken, refreshToken, err := s.auth.RegisterStudent(ctx, usrName, req.GetEmail(), req.GetPassword(), req.GetStudentLogin(), []string{}, []string{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.RegisterStudentResponse{
		StudentId:    userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
