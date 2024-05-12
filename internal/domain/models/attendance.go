package models

import "time"

type Attendance struct {
	ID                   string `json:"id,omitempty"`
	Subject              string
	Student              string
	Schedule             Schedule
	IsRoomCodeScanned    bool
	IsConfirmCodeScanned bool
	IsAttended           bool
}
type AttendanceNoNested struct {
	ID                   string `json:"id,omitempty"`
	Subject              string
	Student              string
	Schedule             string
	IsRoomCodeScanned    bool
	IsConfirmCodeScanned bool
	IsAttended           bool
}

type AttendanceWithFullStudent struct {
	ID                   string `json:"id,omitempty"`
	Subject              string
	Student              Student
	Schedule             Schedule
	IsRoomCodeScanned    bool
	IsConfirmCodeScanned bool
	IsAttended           bool
}

// intermediate model to send between logic and transport layers
type AttendanceLessons struct {
	AttendanceLessons []AttendanceLessonLine
	Date              time.Time
}

type AttendanceLessonLine struct {
	TimeSlot   time.Time
	Group      string
	Subject    string
	IsAttended bool
}

// intermediate model to send between logic and transport layers
type AttendanceJournal struct {
	TimeSlot              time.Time
	Lesson                string
	Group                 string
	AttendanceJournalLine []AttendanceJournalLine
}

type AttendanceJournalKey struct {
	TimeSlot time.Time
	Lesson   string
	Group    string
}

type AttendanceJournalLine struct {
	Number      int
	StudentName string
	IsAttended  bool
}

type AttendanceCodes struct {
	ID       string `json:"id,omitempty"`
	Schedule struct {
		QrCodes  []QrCode
		Timeslot time.Time
	}
}

type AttendanceRoom struct {
	ID       string `json:"id,omitempty"`
	Schedule Schedule
}
