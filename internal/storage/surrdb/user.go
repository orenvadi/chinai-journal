package surrdb

import (
	"context"
	"fmt"

	"github.com/orenvadi/auth-grpc/internal/domain/models"
	"github.com/surrealdb/surrealdb.go"
)

func (s *Storage) SaveTeacher(ctx context.Context, teacher models.Teacher) (uid string, err error) {
	const op = "storage.surrdb.SaveTeacher"
	tchr := struct {
		ID           string `json:"id,omitempty"`
		Name         models.Name
		Email        string
		PasswordHash []byte
		TeacherCode  string
		Groups       []string
		Subjects     []string
	}{
		Name:         teacher.Name,
		Email:        teacher.Email,
		PasswordHash: teacher.PasswordHash,
		TeacherCode:  teacher.TeacherCode,
		Groups:       []string{},
		Subjects:     []string{},
	}
	rawResult, err := s.db.Create("Teacher", tchr)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	createdUser := make([]models.Teacher, 1)

	err = surrealdb.Unmarshal(rawResult, &createdUser)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	uid = string(createdUser[0].ID)

	return uid, nil
}

func (s *Storage) SaveStudent(ctx context.Context, student models.Student) (uid string, err error) {
	const op = "storage.surrdb.SaveStudent"

	stdnt := struct {
		ID           string `json:"id,omitempty"`
		Name         models.Name
		Email        string
		PasswordHash []byte
		StudentCode  string
		Groups       []string
		Subjects     []string
	}{
		Name:         student.Name,
		Email:        student.Email,
		PasswordHash: student.PasswordHash,
		StudentCode:  student.StudentCode,
		Groups:       []string{},
		Subjects:     []string{},
	}
	rawResult, err := s.db.Create("Student", stdnt)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	createdUser := make([]models.Student, 1)

	err = surrealdb.Unmarshal(rawResult, &createdUser)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	uid = string(createdUser[0].ID)

	return uid, nil
}

func (s *Storage) UpdateTeacher(ctx context.Context, usr models.Teacher, teacherLogin, email string) (err error) {
	const op = "storage.surrdb.UpdateTeacher"
	return
}

func (s *Storage) UpdateStudent(ctx context.Context, usr models.Student, email string) (err error) {
	const op = "storage.surrdb.UpdateStudent"
	return
}

func (s *Storage) GetTeacherProfileData(ctx context.Context, email string) (teacher models.Teacher, err error) {
	const op = "storage.surrdb.GetTeacherProfileData"
	return
}

func (s *Storage) GetStudentProfileData(ctx context.Context, email string) (student models.Student, err error) {
	const op = "storage.surrdb.GetStudentProfileData"
	return
}
