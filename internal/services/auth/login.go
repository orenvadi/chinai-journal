package auth

import (
	"context"
	"fmt"
	"log/slog"

	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
	"github.com/orenvadi/auth-grpc/internal/lib/jwt/logger/sl"
	"golang.org/x/crypto/bcrypt"
)

func (a *Auth) LoginTeacher(ctx context.Context, teacherCode, password string) (accessToken string, err error) {
	const op = "auth.LoginTeacher"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("attempting to login user")

	user, err := a.usrProvider.GetTeacherProfileData(ctx, teacherCode)
	if err != nil {
		log.Warn("user not found", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Warn("could not create token", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return accessToken, nil
}

func (a *Auth) LoginStudent(ctx context.Context, studentCode, password string) (accessToken string, err error) {
	const op = "auth.LoginStudent"
	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("attempting to login user")

	user, err := a.usrProvider.GetStudentProfileData(ctx, studentCode)
	if err != nil {
		log.Warn("user not found", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Warn("could not create token", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return accessToken, nil
}
