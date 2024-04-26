package surrdb

import (
	"context"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
)

func (s *Storage) GetAttendanceLessons(ctx context.Context) (lessonsJournal []models.Attendance, err error) {
	const op = "storage.surrdb.GetAttendanceLessons"
	return
}

func (s *Storage) GetAttendanceJournal(ctx context.Context, lessonId string) (attendanceJournal []models.Attendance, err error) {
	const op = "storage.surrdb.GetAttendanceJournal"
	return
}

func (s *Storage) GetConfirmCode(ctx context.Context, usrID string, time time.Time) (attendanceJournal []models.QrCode, err error) {
	const op = "storage.surrdb.GetConfirmCode"
	return
}

func (s *Storage) SubmitCode(ctx context.Context) (err error) {
	const op = "storage.surrdb.SubmitCode"
	return
}
