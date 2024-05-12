package grpcauth

import (
	"context"
	// "log"
	"time"

	// "errors"
	// "log"

	// "github.com/bufbuild/protovalidate-go"
	"github.com/orenvadi/auth-grpc/internal/domain/models"

	// "github.com/orenvadi/auth-grpc/internal/services/auth"
	// "github.com/orenvadi/auth-grpc/internal/storage"
	// "github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	ssov1 "github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type Auth interface {
	RegisterTeacher(ctx context.Context, name models.Name, email, password, teacherCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error)
	RegisterStudent(ctx context.Context, name models.Name, email, password, studentCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error)

	LoginTeacher(ctx context.Context, teacherCode, password string) (accessToken string, err error)
	LoginStudent(ctx context.Context, studentCode, password string) (accessToken string, err error)

	UpdateTeacher(ctx context.Context, email string) error
	UpdateStudent(ctx context.Context, email string) error

	// each func where we need just user id, requires only context, because gRPC headers can be accessed from context
	GetTeacherData(ctx context.Context) (models.Teacher, error)
	GetStudent(ctx context.Context) (models.Student, error)

	SetNewPassword(ctx context.Context, confirmCode, email string, newPassword string) error
}

type ConfCodes interface {
	// teachers
	GetTeachersConfirmCodes(teacherLogin string) (codes []models.QrCode, err error)
	GetAttendanceJournal(ctx context.Context, date time.Time) (journal []models.AttendanceJournal, err error)

	// students
	SubmitRoomCode(ctx context.Context, code string) (err error)
	SubmitTeacherCode(ctx context.Context, code string) (err error)
	GetAttendanceLessons(ctx context.Context, date time.Time) (lessons models.AttendanceLessons, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth      Auth
	confCodes ConfCodes
}

func Register(gRPC *grpc.Server, auth Auth, confCodes ConfCodes) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth, confCodes: confCodes})
}

// const (
// 	emptyValue = 0
// )

// // this took me 8 hours to debug
// func (s *serverAPI) UpdateUser(ctx context.Context, req *ssov1.UpdateUserRequest) (updateUserResponse *ssov1.UpdateUserResponse, err error) {
// 	v, err := protovalidate.New()
// 	if err != nil {
// 		log.Fatalln("error protovalidate", err)
// 	}

// 	// validating
// 	if err := v.Validate(req); err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	if err = s.auth.UpdateUser(ctx, req.GetFirstName(), req.GetLastName(), req.GetPhoneNumber(), req.GetEmail(), req.GetAppId()); err != nil {
// 		return nil, err
// 	}

// 	return &ssov1.UpdateUserResponse{
// 		Success: true,
// 	}, nil
// }

// func (s *serverAPI) GetUserData(ctx context.Context, req *ssov1.GetUserDataRequest) (*ssov1.GetUserDataResponse, error) {
// 	user, err := s.auth.GetUserData(ctx, req.AppId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ssov1.GetUserDataResponse{
// 		Id:          user.ID,
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		PhoneNumber: user.PhoneNumber,
// 		CreatedAt:   timestamppb.New(user.CreatedAt),
// 		UpdatedAt:   timestamppb.New(user.UpdatedAt),
// 		Email:       user.Email,
// 	}, nil
// }
