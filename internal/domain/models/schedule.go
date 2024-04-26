package models

import "time"

type Schedule struct {
	Subject  string
	Group    string
	Teacher  string
	Location string
	Dateslot time.Time
	Timeslot time.Time
	QrCodes  []string
}
