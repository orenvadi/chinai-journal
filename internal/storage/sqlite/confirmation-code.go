// confirmation-code.go
package sqlite

import (
	// "database/sql"
	"chinai-journal/internal/storage"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const (
	qfGetConfirmCodeBySchedId = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_confirmation_code_using_schedule_id.sql"
	qfGetRoomBySchedId        = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_room_using_schedule_id.sql"
	qfGetSheduleByStudentId   = "/home/orenvady/Repos/go/chinai-journal/internal/storage/sqlite/queries/get_schedule_using_student_id.sql"
)

// TODO
func (s *Storage) GetConfirmCodesByStudentId(studentID string) (string, string, error) {
	const op = "storage.sqlite.GetConfirmCodesByStudentId"

	// Step 1: Get Schedule ID for the student
	scheduleID, err := s.getScheduleIDByStudentID(studentID, time.Now())
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}

	// Step 3: Get Confirmation Codes for the specific lesson using the schedule ID and current server time
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}
	firstCode, secondCode, err := s.getConfirmationCodes(scheduleID, time.Now())
	if err != nil {
		return "", "", fmt.Errorf("%s: %v", op, err)
	}

	return firstCode, secondCode, nil
}

// DONE
func (s *Storage) getScheduleIDByStudentID(studentID string, currentTime time.Time) (int, error) {
	const op = "storage.sqlite.getScheduleIDByStudentID"

	query, err := readQueryFromFile(qfGetSheduleByStudentId)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}

	rows, err := s.db.Query(query, studentID)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, err)
	}
	defer rows.Close()

	var scheduleId int = -1

	// По идее это должно быть в конфигах, но мне лень
	location, err := time.LoadLocation("Asia/Bishkek")
	if err != nil {
		log.Println(fmt.Errorf("%s: %v", op, err))
	}

	currentDate := currentTime.In(location).Format("2006-01-02")

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

		fullTimeString := fmt.Sprintf("%s %s", currentDate, timeSlot)

		userTime, err := time.ParseInLocation("2006-01-02 15:04:05", fullTimeString, location)
		if err != nil {
			log.Println(fmt.Errorf("error could not parse time %s: %v", op, err))
		}

		timeDelta := currentTime.Sub(userTime)

		if currentTime.Weekday().String() == dayOfWeek && timeDelta < 91*time.Minute {
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

	query, err := readQueryFromFile(qfGetRoomBySchedId)
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

// DONE
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
		Id     int
		Code   string
		IsUsed int
	}

	confirmationQRCodes := make([]QRCode, 0, 9)

	for rows.Next() {
		qr := QRCode{}
		if err := rows.Scan(
			&qr.Id,
			&qr.Code,
			&qr.IsUsed,
		); err != nil {
			log.Println(fmt.Errorf("%s: %v", op, err))
		}
		confirmationQRCodes = append(confirmationQRCodes, qr)
	}

	// need to check if qrcode is used and return one that not
	// DONE

	// sorting by ascending order
	sort.Slice(confirmationQRCodes, func(i, j int) bool {
		return confirmationQRCodes[i].IsUsed < confirmationQRCodes[j].IsUsed
	})

	for _, qrCode := range confirmationQRCodes {
		if qrCode.IsUsed == 0 {
			secondCode = qrCode.Code
			return firstCode, secondCode, nil
		}
	}

	return "", "", storage.ErrQRExpired
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
