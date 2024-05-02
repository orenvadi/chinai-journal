package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
)

// each func where we need just user id, requires only context, because gRPC headers can be accessed from context
func (a *Auth) GetTeacherData(ctx context.Context) (teacher models.Teacher, err error) {
	const op = "auth.GetTeacherData"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("getting teacher data")
	// Extract user login from token

	claims, err := jwtn.ValidateToken(ctx, a.jwtSecret)
	if err != nil {
		log.Info("could not get claims")
		return models.Teacher{}, fmt.Errorf("invalid token claims")
	}

	usrLogin := claims["login"].(string)

	user, err := a.usrProvider.GetTeacherProfileData(ctx, usrLogin)
	if err != nil {
		log.Info("could not get data from db")
		return models.Teacher{}, fmt.Errorf("error getting data from db")
	}

	return user, nil
}

func (a *Auth) GetStudent(ctx context.Context) (student models.Student, err error) {
	const op = "auth.GetStudent"
	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("getting student data")
	// Extract user login from token

	claims, err := jwtn.ValidateToken(ctx, a.jwtSecret)
	if err != nil {
		log.Info("could not get claims")
		return models.Student{}, fmt.Errorf("invalid token claims")
	}

	usrLogin := claims["login"].(string)

	user, err := a.usrProvider.GetStudentProfileData(ctx, usrLogin)
	if err != nil {
		log.Info("could not get data from db")
		return models.Student{}, fmt.Errorf("error getting data from db")
	}

	return user, nil
}
