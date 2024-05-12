package auth

import (
	"context"
	"fmt"
	"log/slog"

	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
)

// students
func (a *Auth) SubmitRoomCode(ctx context.Context, code string) (err error) {
	const op = "auth.SubmitCode"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("submiting room code")

	claims, err := jwtn.ValidateToken(ctx, a.jwtSecret)
	if err != nil {
		log.Info("could not get claims")
		return fmt.Errorf("invalid token claims")
	}

	usrId := claims["uid"].(string)

	if err := a.confirmationProvider.SubmitRoomCode(ctx, usrId, code); err != nil {
		log.Info("could not submit code")
		fmt.Println("error === ", err)
		return fmt.Errorf("invalid code")
	}

	return nil
}

func (a *Auth) SubmitTeacherCode(ctx context.Context, code string) (err error) {
	const op = "auth.SubmitCode"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("submiting teacher code")

	claims, err := jwtn.ValidateToken(ctx, a.jwtSecret)
	if err != nil {
		log.Info("could not get claims")
		return fmt.Errorf("invalid token claims")
	}

	usrId := claims["uid"].(string)

	if err := a.confirmationProvider.SubmitTeacherCode(ctx, usrId, code); err != nil {
		log.Info("could not submit teacher code")
		return fmt.Errorf("invalid code")
	}

	return nil
}
