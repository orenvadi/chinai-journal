package surrdb

import (
	"context"
	"fmt"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	"github.com/surrealdb/surrealdb.go"
)

type DbResGetAttendanceLessons struct {
	Result []models.Attendance `json:"result"`
}

func (s *Storage) GetAttendanceLessons(
	ctx context.Context,
	date time.Time,
	userLogin string,
) (
	lessonsJournal []models.Attendance,
	err error,
) {
	const op = "storage.surrdb.GetAttendanceLessons"

	schema := `
    SELECT 
        IsAttended, IsConfirmCodeScanned, IsRoomCodeScanned, Schedule[*], 
        Schedule.Subject[*], Schedule.Group[*], Schedule.Location[*], Student, id 
    FROM 
        Attendance 
    WHERE 
        (SELECT 
            VALUE Groups 
        FROM 
            (SELECT 
                VALUE id 
            FROM 
                Student 
            WHERE 
                StudentCode = $studentLogin
            )[0]
        )[0] CONTAINS Schedule.Group 
        AND Student.StudentCode is $studentLogin 
        AND time::format(Schedule.Dateslot, '%d-%m-%Y') IS $date;
  `

	data, err := s.db.Query(schema, map[string]string{
		"studentLogin": userLogin,
		"date":         date.Format("02-01-2006"),
	})
	if err != nil {
		return []models.Attendance{}, fmt.Errorf("%s: %w", op, err)
	}

	res := []DbResGetAttendanceLessons{}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return []models.Attendance{}, fmt.Errorf("%s: %w", op, err)
	}

	attendances := res[0].Result

	return attendances, nil
}

type DbResGetAttendanceJournal struct {
	Result []models.AttendanceWithFullStudent `json:"result"`
}

func (s *Storage) GetAttendanceJournal(
	ctx context.Context,
	date time.Time,
	userLogin string,
) (
	attendanceJournal []models.AttendanceWithFullStudent,
	err error,
) {
	const op = "storage.surrdb.GetAttendanceJournal"

	schema := `
    SELECT 
      IsAttended, IsConfirmCodeScanned, IsRoomCodeScanned, Schedule[*], 
      Schedule.Subject[*], Schedule.Group[*], Schedule.Location[*], Student[*], id 
    FROM 
      Attendance 
    WHERE 
      (SELECT 
          VALUE Groups 
      FROM 
          (SELECT 
              VALUE id 
          FROM 
              Teacher 
          WHERE 
              TeacherCode is $teacherLogin
          )[0]
      )[0] CONTAINS Schedule.Group
      AND Schedule.Teacher.TeacherCode is $teacherLogin
      AND time::format(Schedule.Dateslot, '%d-%m-%Y') IS $date;
  `
	data, err := s.db.Query(schema, map[string]string{
		"teacherLogin": userLogin,
		"date":         date.Format("02-01-2006"),
	})
	if err != nil {
		return []models.AttendanceWithFullStudent{}, fmt.Errorf("%s: %w", op, err)
	}

	res := []DbResGetAttendanceJournal{}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return []models.AttendanceWithFullStudent{}, fmt.Errorf("%s: %w", op, err)
	}

	attendances := res[0].Result

	return attendances, nil
}

func (s *Storage) GetConfirmCode(ctx context.Context, userID string, time time.Time) (attendanceJournal []models.QrCode, err error) {
	const op = "storage.surrdb.GetConfirmCode"
	return
}

func (s *Storage) SubmitCode(ctx context.Context, userId, code string) (err error) {
	const op = "storage.surrdb.SubmitCode"
	return
}
