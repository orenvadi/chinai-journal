package auth

import (
	"fmt"
	"log/slog"
	"time"

	// "time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
)

// teachers
func (a *Auth) GetTeachersConfirmCodes(usrLogin string) (codes []models.QrCode, err error) {
	const op = "auth.GetTeachersConfirmCodes"

	log := a.log.With(
		slog.String("op: ", op),
	)

	log.Info("getting codes for teacher")
	// Extract user login from token

	rawCodes, err := a.confirmationProvider.GetConfirmCode(usrLogin)
	if err != nil {
		log.Info("could not get data from db")
		return []models.QrCode{}, fmt.Errorf("error getting data from db")
	}

	now := time.Now()

	for _, code := range rawCodes {
		if now.Sub(code.Timeslot) >= (75*time.Minute) &&
			now.Sub(code.Timeslot) <= (90*time.Minute) {
			return code.QrCodes, nil
		}
	}

	return []models.QrCode{}, fmt.Errorf("some error occured")
}
