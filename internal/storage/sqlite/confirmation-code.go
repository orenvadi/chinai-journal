// confirmation-code.go
package sqlite

import (
	// "database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	qfGetConfirmCodeBySchedId = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_confirmation_code_using_schedule_id.sql"
	qfGetRoomByStudentId      = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_room_using_schedule_id.sql"
	qfGetRoomCodeByScheduleId = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_room_code_using_schedule_id.sql"
	qfGetSheduleByStudentId   = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_schedule_using_student_id.sql"
)

// TODO
func (s *Storage) GetConfirmCodesByStudentId(studentID string) (string, string, error) {
	const op = "storage.sqlite.GetConfirmCodesByStudentId"

	// Step 1: Get Schedule ID for the student
	scheduleID, err := s.getScheduleIDByStudentID(studentID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}

	// Step 2: Get Room information for the lesson using the schedule ID
	// room, err := s.getRoomByScheduleID(scheduleID)
	// if err != nil {
	// 	return "", "", fmt.Errorf("%s: %v", op, err)
	// }

	// Step 3: Get Confirmation Codes for the specific lesson using the schedule ID and current server time
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}
	firstCode, secondCode, err := s.getConfirmationCodes(scheduleID, time.Now())

	return fmt.Sprintf("%d", firstCode), fmt.Sprintf("%d", secondCode), nil
}

// DONE
func (s *Storage) getScheduleIDByStudentID(studentID string, currentTime time.Time) (int, error) {
	const op = "storage.sqlite.getScheduleIDByStudentID"

	query, err := readQueryFromFile(qfGetSheduleByStudentId)

	rows, err := s.db.Query(query, studentID)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}
	defer rows.Close()

	var scheduleId int = -1

	// finding needed id
	for rows.Next() {
		var (
			scheduleId   int
			subjName     string
			teacherName  string
			groupName    string
			locationName string
			dayOfWeek    string
			timeSlot     string
		)

		if err := rows.Scan(
			&scheduleId,
			&subjName,
			&teacherName,
			&groupName,
			&locationName,
			&dayOfWeek,
			&timeSlot,
		); err != nil {
			log.Println(fmt.Errorf("%s: %v", op, err))
		}

		if currentTime.Weekday().String() == dayOfWeek {
			return scheduleId, nil
		}
	}

	if scheduleId == -1 {
		return -1, fmt.Errorf("%s: %v", op, errors.New("could not scan to scheduleId or the table is empty"))
	}

	return scheduleId, nil
}

// DONE
func (s *Storage) getRoomByScheduleID(scheduleID int) (int, string, error) {
	const op = "storage.sqlite.getRoomByScheduleID"

	query, err := readQueryFromFile(qfGetRoomCodeByScheduleId)
	if err != nil {
		return -1, "", fmt.Errorf("%s: %v", op, err)
	}

	var (
		auditoriumID     int
		auditoriumName   string
		auditoriumQRCode string
	)

	err = s.db.QueryRow(query, scheduleID).Scan(
		&auditoriumID,
		&auditoriumName,
		&auditoriumQRCode,
	)
	if err != nil {
		return -1, "", fmt.Errorf("%s: %v", op, err)
	}

	return auditoriumID, auditoriumQRCode, nil
}

// DOING
func (s *Storage) getConfirmationCodes(scheduleID int, currentTime time.Time) (string, string, error) {
	op := "storage.sqlite.getConfirmationCodes"

	var firstCode, secondCode string

	// first code
	_, firstCode, err := s.getRoomByScheduleID(scheduleID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}

	// second code
	query2, err := readQueryFromFile(qfGetConfirmCodeBySchedId)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	rows, err := s.db.Query(query2, scheduleID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	type QRCode struct {
		Id   int
		Code string
	}

	confirmationQRCodes := make([]QRCode, 0, 9)

	for rows.Next() {
		qr := QRCode{}
		if err := rows.Scan(
			&qr.Id,
			&qr.Code,
		); err != nil {
			log.Println(fmt.Errorf("%s: %v", op, err))
		}
		confirmationQRCodes = append(confirmationQRCodes, qr)
	}

	// need to check if qrcode is used and return one that not

	return firstCode, secondCode, nil
}

// DONE
func readQueryFromFile(filePath string) (string, error) {
	// returns list of queries to create tables
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return "", err
	}

	query := string(fileContent)

	return query, nil
}
