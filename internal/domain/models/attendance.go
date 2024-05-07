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
