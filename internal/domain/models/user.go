package models

type User interface {
	Student | Teacher
}

type Name struct {
	FirstName  string
	LastName   string
	Patronimic string
}

type Student struct {
	ID           string `json:"id"`
	Name         Name
	Email        string
	PasswordHash []byte
	StudentCode  string
	Groups       []string
	Subjects     []string
}

type Teacher struct {
	ID           string `json:"id"`
	Name         Name
	Email        string
	PasswordHash []byte
	TeacherCode  string
	Groups       []string
	Subjects     []string
}

func (s Student) Id() string {
	return s.ID
}

func (t Teacher) Id() string {
	return t.ID
}

func (s Student) FullName() string {
	return s.Name.FirstName + " " + s.Name.LastName
}

func (t Teacher) FullName() string {
	return t.Name.FirstName + " " + t.Name.LastName
}

func (t Teacher) GetEmail() string {
	return t.Email
}

func (s Student) GetEmail() string {
	return s.Email
}
