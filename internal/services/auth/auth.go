package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	jwtn "github.com/orenvadi/auth-grpc/internal/lib/jwt"
	"github.com/orenvadi/auth-grpc/internal/lib/jwt/logger/sl"
	"golang.org/x/crypto/bcrypt"
	// "github.com/orenvadi/auth-grpc/internal/lib/jwt"
	// "github.com/orenvadi/auth-grpc/internal/lib/rnd"
	// "github.com/orenvadi/auth-grpc/internal/storage"
)

type Auth struct {
	log                  *slog.Logger
	usrSaver             UserSaver
	usrUpdater           UserUpdater
	usrProvider          UserProvider
	attendanceProvider   AttendanceProvider
	confirmationProvider ConfirmationProvider
	tokenTTL             time.Duration
	jwtSecret            string
}

type UserSaver interface {
	SaveTeacher(ctx context.Context, teacher models.Teacher) (uid string, err error)
	SaveStudent(ctx context.Context, student models.Student) (uid string, err error)
}

type UserUpdater interface {
	UpdateTeacher(ctx context.Context, usr models.Teacher, teacherLogin, email string) (err error)
	UpdateStudent(ctx context.Context, usr models.Student, email string) (err error)
}

type UserProvider interface {
	GetTeacherProfileData(ctx context.Context, email string) (models.Teacher, error)
	GetStudentProfileData(ctx context.Context, email string) (models.Student, error)
}

type AttendanceProvider interface {
	GetAttendanceLessons(ctx context.Context) ([]models.Attendance, error)
	GetAttendanceJournal(ctx context.Context, lessonId string) ([]models.Attendance, error)
}

type ConfirmationProvider interface {
	GetConfirmCode(ctx context.Context, usrID string, time time.Time) ([]models.QrCode, error)
	SubmitCode(ctx context.Context) error
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

// New return a new instance of the Auth service.
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	userUpdater UserUpdater,
	attendanceProvider AttendanceProvider,
	confirmationProvider ConfirmationProvider,
	tokenTTL time.Duration,
	jwtSecret string,
) *Auth {
	// Can cause panic because of nil pointer
	return &Auth{
		log:                  log,
		usrSaver:             userSaver,
		usrProvider:          userProvider,
		usrUpdater:           userUpdater, // из-за этой херни я потерял 3 часа
		attendanceProvider:   attendanceProvider,
		confirmationProvider: confirmationProvider,
		tokenTTL:             tokenTTL,
		jwtSecret:            jwtSecret,
	}
}

func (a *Auth) RegisterTeacher(ctx context.Context, name models.Name, email, password, teacherCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error) {
	const op = "auth.RegisterTeacher"

	log := a.log.With(
		slog.String("op: ", op),
		// slog.String("email: ", email), // do not do that
	)

	log.Info("registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	user := models.Teacher{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		TeacherCode:  teacherCode,
		Groups:       groups,
		Subjects:     subjects,

		// ID
		// Name
		// Email
		// PasswordHash
		// TeacherCode
		// Groups
		// Subjects
	}
	userID, err = a.usrSaver.SaveTeacher(ctx, user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return userID, accessToken, "", nil
}

func (a *Auth) RegisterStudent(ctx context.Context, name models.Name, email, password, studentCode string, groups, subjects []string) (userID string, accessToken, refreshToken string, err error) {
	const op = "auth.RegisterStudent"

	log := a.log.With(
		slog.String("op: ", op),
		// slog.String("email: ", email), // do not do that
	)

	log.Info("registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	user := models.Student{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		StudentCode:  studentCode,
		Groups:       groups,
		Subjects:     subjects,

		// ID
		// Name
		// Email
		// PasswordHash
		// StudentCode
		// Groups
		// Subjects
	}
	userID, err = a.usrSaver.SaveStudent(ctx, user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err = jwtn.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return userID, accessToken, "", nil
}

func (a *Auth) LoginTeacher(ctx context.Context, teacherCode, password string) (accessToken string, err error) {
	const op = "auth.LoginTeacher"
	return
}

func (a *Auth) LoginStudent(ctx context.Context, studentCode, password string) (accessToken string, err error) {
	const op = "auth.LoginStudent"
	return
}

func (a *Auth) UpdateTeacher(ctx context.Context, email string) (err error) {
	const op = "auth.UpdateTeacher"
	return
}

func (a *Auth) UpdateStudent(ctx context.Context, email string) (err error) {
	const op = "auth.UpdateStudent"
	return
}

// each func where we need just user id, requires only context, because gRPC headers can be accessed from context
func (a *Auth) GetTeacherData(ctx context.Context) (teacher models.Teacher, err error) {
	const op = "auth.GetTeacherData"
	return
}

func (a *Auth) GetStudent(ctx context.Context) (student models.Student, err error) {
	const op = "auth.GetStudent"
	return
}

func (a *Auth) SetNewPassword(ctx context.Context, confirmCode, email string, newPassword string) (err error) {
	const op = "auth.SetNewPassword"
	return
}

// teachers
func (a *Auth) GetTeachersConfirmCodes(ctx context.Context) (codes []models.QrCode, err error) {
	const op = "auth.GetTeachersConfirmCodes"
	return
}

func (a *Auth) GetAttendanceJournal(ctx context.Context, date time.Time) (journal []models.Attendance, err error) {
	const op = "auth.GetAttendanceJournal"
	return
}

// students
func (a *Auth) SubmitCode(ctx context.Context, code string) (err error) {
	const op = "auth.SubmitCode"
	return
}

func (a *Auth) GetAttendanceLessons(ctx context.Context, date time.Time) (lessons []models.Attendance, err error) {
	const op = "auth.GetAttendanceLessons"
	return
}

// // Login checks if user with given credentials exists in the system and returns access token.
// //
// // If user exists, but password is incorrect, returns error.
// // If user doesn't exist, returns error.
// func (a *Auth) Login(ctx context.Context, email, password string, appID int64) (accessToken string, err error) {
// 	const op = "auth.Login"

// 	log := a.log.With(
// 		slog.String("op: ", op),
// 		// slog.String("email: ", email), // do not do that
// 	)

// 	log.Info("attempting to login user")

// 	user, err := a.usrProvider.User(ctx, email)
// 	if err != nil {
// 		if errors.Is(err, storage.ErrUserNotFound) {
// 			a.log.Warn("user not found", sl.Err(err))

// 			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
// 		}

// 		log.Warn("user not found", sl.Err(err))
// 		return "", fmt.Errorf("%s: %w", op, err)
// 	}

// 	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
// 		log.Info("invalid credentials", sl.Err(err))

// 		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
// 	}

// 	app, err := a.appProvider.App(ctx, appID)
// 	if err != nil {
// 		return "", fmt.Errorf("%s: %w", op, err)
// 	}
// 	log.Info("user logged in successfully")

// 	accessToken, err = jwtn.NewToken(user, app, a.tokenTTL)
// 	if err != nil {
// 		log.Error("failed to generate token", sl.Err(err))

// 		return "", fmt.Errorf("%s: %w", op, err)
// 	}

// 	return accessToken, nil
// }

// // UpdateUser updates user information.
// func (a *Auth) UpdateUser(ctx context.Context, firstName, lastName, phoneNumber, email string, appID int64) error {
// 	const op = "auth.UpdateUser"

// 	log := a.log.With(
// 		slog.String("op: ", op),
// 		// slog.String("user_email", email),
// 	)

// 	// Extract username from token

// 	app, err := a.appProvider.App(ctx, appID)
// 	if err != nil {
// 		if errors.Is(err, storage.ErrUserNotFound) {
// 			a.log.Warn("user not found", sl.Err(err))

// 			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
// 		}

// 		log.Warn("user not found", sl.Err(err))
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	claims, err := jwtn.ValidateToken(ctx, app)
// 	if err != nil {
// 		return fmt.Errorf("invalid token claims")
// 	}
// 	userID := claims["uid"].(float64)
// 	uid := int64(userID)

// 	log.Info("updating user")

// 	// Retrieve the user from the storage
// 	user, err := a.usrProvider.UserAllData(ctx, uid)
// 	if err != nil {
// 		if errors.Is(err, storage.ErrUserNotFound) {
// 			a.log.Warn("user not found", sl.Err(err))
// 			return ErrInvalidCredentials // or ErrUserNotFound
// 		}

// 		log.Warn("user not found", sl.Err(err))
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	// Update user information
// 	user.FirstName = firstName
// 	user.LastName = lastName
// 	user.PhoneNumber = phoneNumber
// 	user.Email = email

// 	// Hash and update password if provided
// 	// if password != "" {
// 	// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	// 	if err != nil {
// 	// 		log.Error("failed to generate password hash", sl.Err(err))
// 	// 		return fmt.Errorf("%s: %w", op, err)
// 	// 	}
// 	// 	user.PasswordHash = passwordHash
// 	// }

// 	// log.Info("upd: ", sl.Err(fmt.Errorf(fmt.Sprintf("%v", user))))

// 	// Save updated user information to the storage

// 	err = a.usrUpdater.UpdateUser(ctx, user)
// 	if err != nil {
// 		log.Error("failed to update user", sl.Err(err))
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	log.Info("user updated successfully")

// 	return nil
// }

// func (a *Auth) GetUserData(ctx context.Context, appID int64) (models.User, error) {
// 	const op = "auth.GetUserData"

// 	log := a.log.With(
// 		slog.String("op: ", op),
// 		// slog.String("user_email", email),
// 	)

// 	// Extract username from token

// 	app, err := a.appProvider.App(ctx, appID)
// 	if err != nil {
// 		if errors.Is(err, storage.ErrUserNotFound) {
// 			a.log.Warn("user not found", sl.Err(err))

// 			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
// 		}

// 		log.Warn("user not found", sl.Err(err))
// 		return models.User{}, fmt.Errorf("%s: %w", op, err)
// 	}

// 	claims, err := jwtn.ValidateToken(ctx, app)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("invalid token claims")
// 	}
// 	userID := claims["uid"].(float64)
// 	uid := int64(userID)

// 	user, err := a.usrProvider.UserAllData(ctx, uid)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("%s: %w", op, err)
// 	}

// 	return user, nil
// }

// func (a *Auth) SetNewPassword(ctx context.Context, confirmCode, email string, newPassword string) error {
// 	const op = "auth.SetNewPassword"

// 	log := a.log.With(
// 		slog.String("op: ", op),
// 		// slog.String("email: ", email), // do not do that
// 	)

// 	user, err := a.usrProvider.User(ctx, email)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	userAllData, err := a.usrProvider.UserAllData(ctx, user.ID)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Error("failed to generate password hash", sl.Err(err))
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	uid := userAllData.ID

// 	confCodeFromDB, err := a.emailConfirmProvider.ConfirmationCode(ctx, uid)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	location, _ := time.LoadLocation("Asia/Bishkek")
// 	now := time.Now().In(location)
// 	// now := time.Now()

// 	if confirmCode != confCodeFromDB.Code {
// 		return fmt.Errorf("%s: invalid confirm code", op)
// 	}

// 	if elapsedTime := now.Sub(confCodeFromDB.CreatedAt.In(location)); elapsedTime > (5 * time.Minute) {
// 		return fmt.Errorf("%s: confirm code is expired", op)
// 	}

// 	if err = a.usrProvider.UserEmailConfirm(ctx, uid); err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	if err = a.emailConfirmProvider.DeleteConfirmationCode(ctx, confCodeFromDB.ID); err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	if err = a.passwordResetter.ChangePassword(ctx, userAllData.Email, passwordHash); err != nil {
// 		log.Error("failed to change pass", sl.Err(err))
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	return nil
// }
