package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
)

func (a *Auth) GetAttendanceLessons(
	ctx context.Context,
	date time.Time,
) (
	lessons models.AttendanceLessons,
	err error,
) {
	const op = "auth.GetAttendanceLessons"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("getting attendance lessons")

	claims, err := jwtn.ValidateToken(ctx, a.jwtSecret)
	if err != nil {
		log.Info("could not get claims")
		return models.AttendanceLessons{}, fmt.Errorf("invalid token claims")
	}

	usrLogin := claims["login"].(string)

	rawLessons, err := a.attendanceProvider.GetAttendanceLessons(ctx, date, usrLogin)
	if err != nil {
		log.Info("could not get data from db")
		return models.AttendanceLessons{}, fmt.Errorf("error getting data from db")
	}

	lessonLines := []models.AttendanceLessonLine{}

	for _, rawLesson := range rawLessons {
		lessonLines = append(lessonLines, models.AttendanceLessonLine{
			TimeSlot:   rawLesson.Schedule.Timeslot,
			Group:      rawLesson.Schedule.Group.Name,
			Subject:    rawLesson.Schedule.Subject.Name,
			IsAttended: rawLesson.IsAttended,
		})
	}

	lessons = models.AttendanceLessons{
		Date:              date,
		AttendanceLessons: lessonLines,
	}

	return lessons, nil
}

func (a *Auth) GetAttendanceJournal(
	ctx context.Context,
	date time.Time,
) (
	journal []models.AttendanceLessons,
	err error,
) {
	const op = "auth.GetAttendanceJournal"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("getting attendance journal")

	return
}
