package surrdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	"github.com/surrealdb/surrealdb.go"
)

type DbResGetConfirmCode struct {
	Result []models.ScheduleQrCodes `json:"result"`
}

func (s *Storage) GetConfirmCode(
	userLogin string,
) (
	codes []models.ScheduleQrCodes,
	err error,
) {
	const op = "storage.surrdb.GetConfirmCode"

	schema := `
    SELECT QrCodes[*][*], Timeslot FROM Schedule 
    WHERE 
      Teacher.TeacherCode = $teacherLogin;
  `

	data, err := s.db.Query(schema, map[string]string{
		"teacherLogin": userLogin,
	})
	if err != nil {
		return []models.ScheduleQrCodes{}, fmt.Errorf("%s: %w", op, err)
	}

	res := []DbResGetConfirmCode{}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return []models.ScheduleQrCodes{}, fmt.Errorf("%s: %w", op, err)
	}

	if len(res[0].Result) == 0 {
		return []models.ScheduleQrCodes{}, fmt.Errorf("%s: %w", op, errors.New("invalid username"))
	}
	return res[0].Result, nil
}

type DbResSubmitCode struct {
	Result []models.AttendanceCodes `json:"result"`
}

func (s *Storage) SubmitTeacherCode(
	ctx context.Context,
	userId,
	code string,
) (
	err error,
) {
	const op = "storage.surrdb.SubmitTeacherCode"

	// 1 find current scheduleLesson and AttendanceLesson of student

	data, err := s.db.Query(`
  SELECT id, Schedule[*].QrCodes[*][*], Schedule[*].Timeslot FROM Attendance 
    WHERE 
    (SELECT VALUE Groups FROM $userId)[0] 
        CONTAINS Schedule.Group 
    AND Student is $userId 
    AND time::format(Schedule.Dateslot, '%Y-%m-%d') is $date;
    `, map[string]string{
		"userId": userId,
		"date":   time.Now().Format(time.RFC3339)[:10],
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := []DbResSubmitCode{}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// 2 find qrcode that has the same code as the code argument

	var neededAtndCodes models.AttendanceCodes

	now := time.Now()

	if len(res[0].Result) == 0 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid input data"))
	}
	// finding codes for the current lesson
	for _, attndCodes := range res[0].Result {
		codeTimeslot := attndCodes.Schedule.Timeslot
		// this must be it config but i am lazy

		theSub := now.Sub(codeTimeslot)
		if theSub >= (75*time.Minute) &&
			theSub <= (90*time.Minute) {
			neededAtndCodes = attndCodes
		}
	}

	// 3 check if code is outdated

	var neededQrCode models.QrCode

	for _, theCode := range neededAtndCodes.Schedule.QrCodes {
		if theCode.Code == code {
			neededQrCode = theCode
		}
	}

	if now.Sub(neededQrCode.FirstUseTime).Minutes() <= 2 {

		attndId := neededAtndCodes.ID
		data, err := s.db.Select(attndId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		selectedAttnd := new(models.AttendanceNoNested)

		err = surrealdb.Unmarshal(data, &selectedAttnd)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		// set isAttended true if code is correct and room code is scanned
		if selectedAttnd.IsRoomCodeScanned {
			_, err := s.db.Query(`
        UPDATE $attndId SET
          IsConfirmCodeScanned = true,
          IsAttended = true;
        `,
				map[string]string{
					"attndId": attndId,
				})
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		}

		return nil

	} else if neededQrCode.FirstUseTime.Year() == 1 {
		attndId := neededAtndCodes.ID
		data, err := s.db.Select(attndId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		selectedAttnd := new(models.AttendanceNoNested)

		err = surrealdb.Unmarshal(data, &selectedAttnd)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if selectedAttnd.IsRoomCodeScanned {
			_, err := s.db.Query(`
        UPDATE $attndId SET
          IsConfirmCodeScanned = true,
          IsAttended = true;
        `,
				map[string]string{
					"attndId": attndId,
				})
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		}

		qrcodeId := neededQrCode.ID

		_, err = s.db.Query(`
        UPDATE $qrcodeId SET
          FirstUseTime = $time;
        `,
			map[string]string{
				"qrcodeId": qrcodeId,
				"time":     time.Now().Format(time.RFC3339),
			})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil

	}

	return fmt.Errorf("%s: %w", op, errors.New("code is outdated or not exists"))
}

// works
func (s *Storage) SubmitRoomCode(
	ctx context.Context,
	userId,
	code string,
) (
	err error,
) {
	const op = "storage.surrdb.SubmitRoomCode"

	type DbResGetAttendanceRoom struct {
		Result []models.AttendanceRoom `json:"result"`
	}

	schema := `
    SELECT id, Schedule[*], Schedule.Location[*], 
      Schedule.Subject[*], Schedule.Group[*] 
      FROM Attendance
        WHERE 
        (SELECT VALUE Groups FROM $userId)[0] 
            CONTAINS Schedule.Group 
        AND Student is $userId
        AND time::format(Schedule.Dateslot, '%Y-%m-%d') is $date;
  `

	data, err := s.db.Query(schema, map[string]string{
		"userId": userId,
		"date":   time.Now().Format(time.RFC3339)[:10],
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	res := []DbResGetAttendanceRoom{}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// finding current lesson room
	var neededAttndRoom models.AttendanceRoom
	now := time.Now()

	if len(res[0].Result) == 0 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid input data"))
	}
	for _, attndRoom := range res[0].Result {
		codeTimeslot := attndRoom.Schedule.Timeslot
		// this must be it config but i am lazy
		theSub := now.Sub(codeTimeslot)
		if theSub >= (0*time.Minute) &&
			theSub <= (15*time.Minute) {
			neededAttndRoom = attndRoom
		}
	}

	// set isRoomCodeScannedCode true if code is correct
	if neededAttndRoom.Schedule.Location.PrelimConfirmCode == code {
		attndId := neededAttndRoom.ID
		_, err := s.db.Query(`
        UPDATE $attndId SET
          IsRoomCodeScanned = true;
        `,
			map[string]string{
				"attndId": attndId,
			})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil
	}

	return fmt.Errorf("%s: %w", op, errors.New("code is not exists"))
}
