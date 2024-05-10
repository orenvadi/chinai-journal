package models

import "time"

type QrCode struct {
	ID           string `json:"id,omitempty"`
	Code         string
	FirstUseTime time.Time
}

type ScheduleQrCodes struct {
	ID       string `json:"id,omitempty"`
	QrCodes  []QrCode
	Timeslot time.Time
}
