package models

type Attendance struct {
	ID                   string `json:"id,omitempty"`
	Subject              string
	Student              string
	Schedule             string
	IsRoomCodeScanned    bool
	IsConfirmCodeScanned bool
	IsAttended           bool
}
