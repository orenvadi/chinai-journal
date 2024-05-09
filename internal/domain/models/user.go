package models

type Name struct {
	First      string
	Last       string
	Patronimic string
}

type Student struct {
	ID           string `json:"id,omitempty"`
	Name         Name
	Email        string
	PasswordHash string // should be bytes, but cant unmarshall to []bytes from surrdb
	StudentCode  string
	Groups       []string
	Subjects     []string
}

type Teacher struct {
	ID           string `json:"id,omitempty"`
	Name         Name
	Email        string
	PasswordHash string // should be bytes, but cant unmarshall to []bytes from surrdb
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
	return s.Name.First + " " + s.Name.Last
}

func (t Teacher) FullName() string {
	return t.Name.First + " " + t.Name.Last
}

func (t Teacher) GetEmail() string {
	return t.Email
}

func (s Student) GetEmail() string {
	return s.Email
}

func (t Teacher) GetUserCode() string {
	return t.TeacherCode
}

func (s Student) GetUserCode() string {
	return s.StudentCode
}
