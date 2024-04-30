package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
	"github.com/orenvadi/auth-grpc/internal/lib/jwt/logger/sl"
	"golang.org/x/crypto/bcrypt"
)

func (a *Auth) RegisterTeacher(ctx context.Context, name models.Name, email, password, teacherCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error) {
	const op = "auth.RegisterTeacher"

	log := a.log.With(
		slog.String("op: ", op),
		// slog.String("email: ", email), // do not do that
	)

	log.Info("registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	user := models.Teacher{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		TeacherCode:  teacherCode,
		Groups:       groups,
		Subjects:     subjects,

		// ID
		// Name
		// Email
		// PasswordHash
		// TeacherCode
		// Groups
		// Subjects
	}
	userID, err = a.usrSaver.SaveTeacher(ctx, user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return userID, accessToken, "", nil
}

func (a *Auth) RegisterStudent(ctx context.Context, name models.Name, email, password, studentCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error) {
	const op = "auth.RegisterStudent"

	log := a.log.With(
		slog.String("op: ", op),
		// slog.String("email: ", email), // do not do that
	)

	log.Info("registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	user := models.Student{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		StudentCode:  studentCode,
		Groups:       groups,
		Subjects:     subjects,

		// ID
		// Name
		// Email
		// PasswordHash
		// StudentCode
		// Groups
		// Subjects
	}
	userID, err = a.usrSaver.SaveStudent(ctx, user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return userID, accessToken, "", nil
}
