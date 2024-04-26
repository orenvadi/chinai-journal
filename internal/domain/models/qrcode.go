package models

type QrCode struct {
	ID     string `json:"id,omitempty"`
	Code   string
	IsUsed bool
}
