package grpcauth

import (
	"context"

	"github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverAPI) GetTeacherProfileData(ctx context.Context, empty *emptypb.Empty) (res *sso.GetTeacherProfileDataResponse, err error) {
	user, err := s.auth.GetTeacherData(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.GetTeacherProfileDataResponse{
		FirstName:    user.Name.First,
		LastName:     user.Name.Last,
		Patronimic:   user.Name.Patronimic,
		TeacherLogin: user.TeacherCode,
		Email:        user.GetEmail(),
	}, nil
}

func (s *serverAPI) GetStudentProfileData(ctx context.Context, empty *emptypb.Empty) (res *sso.GetStudentProfileDataResponse, err error) {
	user, err := s.auth.GetStudent(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &sso.GetStudentProfileDataResponse{
		FirstName:    user.Name.First,
		LastName:     user.Name.Last,
		Patronimic:   user.Name.Patronimic,
		StudentLogin: user.StudentCode,
		Email:        user.GetEmail(),
	}, nil
}
