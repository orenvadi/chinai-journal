package auth

import (
	"context"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
)

// each func where we need just user id, requires only context, because gRPC headers can be accessed from context
func (a *Auth) GetTeacherData(ctx context.Context) (teacher models.Teacher, err error) {
	const op = "auth.GetTeacherData"
	return
}

func (a *Auth) GetStudent(ctx context.Context) (student models.Student, err error) {
	const op = "auth.GetStudent"
	return
}
