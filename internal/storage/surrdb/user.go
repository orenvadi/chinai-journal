package surrdb

import (
	"context"
	"errors"
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
		PasswordHash string
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
		PasswordHash string
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
	return fmt.Errorf("%s: %w", op, errors.New("unimplemented UpdateTeacher"))
}

func (s *Storage) UpdateStudent(ctx context.Context, usr models.Student, email string) (err error) {
	const op = "storage.surrdb.UpdateStudent"
	return fmt.Errorf("%s: %w", op, errors.New("unimplemented UpdateStudent"))
}

type DbResGetTeacherProfileData struct {
	Result []models.Teacher `json:"result"`
}

func (s *Storage) GetTeacherProfileData(ctx context.Context, teacherLogin string) (teacher models.Teacher, err error) {
	const op = "storage.surrdb.GetTeacherProfileData"
	res := []DbResGetTeacherProfileData{}

	var data interface{}

	println(teacherLogin)
	data, err = s.db.Query("SELECT * FROM Teacher WHERE TeacherCode = $teacherLogin;", map[string]string{
		"teacherLogin": teacherLogin,
	})
	if err != nil {
		return models.Teacher{}, fmt.Errorf("%s: %w", op, err)
	}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return models.Teacher{}, fmt.Errorf("%s: %w", op, err)
	}

	if len(res[0].Result) == 0 {
		return models.Teacher{}, fmt.Errorf("%s: %w", op, errors.New("invalid username"))
	}
	teacher = res[0].Result[0]

	return teacher, nil
}

type DbResGetStudentProfileData struct {
	Result []models.Student `json:"result"`
}

func (s *Storage) GetStudentProfileData(ctx context.Context, studentLogin string) (student models.Student, err error) {
	const op = "storage.surrdb.GetStudentProfileData"
	res := []DbResGetStudentProfileData{}

	var data interface{}

	data, err = s.db.Query("SELECT * FROM Student WHERE StudentCode = $studentLogin;", map[string]string{
		"studentLogin": studentLogin,
	})
	if err != nil {
		return models.Student{}, fmt.Errorf("%s: %w", op, err)
	}

	err = surrealdb.Unmarshal(data, &res)
	if err != nil {
		return models.Student{}, fmt.Errorf("%s: %w", op, err)
	}

	if len(res[0].Result) == 0 {
		return models.Student{}, fmt.Errorf("%s: %w", op, errors.New("invalid username"))
	}

	student = res[0].Result[0]

	return student, nil
}
