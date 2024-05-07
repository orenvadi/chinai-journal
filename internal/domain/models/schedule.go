package models

import "time"

type Schedule struct {
	Subject  Subject
	Group    Group
	Teacher  string
	Location Location
	Dateslot time.Time
	Timeslot time.Time
	QrCodes  []string
}
