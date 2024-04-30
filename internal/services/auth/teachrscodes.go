package auth

import (
	"context"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
)

// teachers
func (a *Auth) GetTeachersConfirmCodes(ctx context.Context) (codes []models.QrCode, err error) {
	const op = "auth.GetTeachersConfirmCodes"
	return
}

func (a *Auth) GetAttendanceJournal(ctx context.Context, date time.Time) (journal []models.Attendance, err error) {
	const op = "auth.GetAttendanceJournal"
	return
}
