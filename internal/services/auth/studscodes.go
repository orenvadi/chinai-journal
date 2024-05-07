package auth

import (
	"context"
	// "time"
	// "github.com/orenvadi/auth-grpc/internal/domain/models"
)

// students
func (a *Auth) SubmitCode(ctx context.Context, code string) (err error) {
	const op = "auth.SubmitCode"
	return
}

// func (a *Auth) GetAttendanceLessons(ctx context.Context, date time.Time) (lessons []models.Attendance, err error) {
// 	const op = "auth.GetAttendanceLessons"
// 	return
// }
