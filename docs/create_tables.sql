-- Создание таблицы студентов
CREATE TABLE IF NOT EXISTS Students (
    StudentID INTEGER PRIMARY KEY,
    FirstName TEXT,
    LastName TEXT,
    StudentCode TEXT UNIQUE,
    PasswordHash TEXT,
    Role TEXT DEFAULT 'student' CHECK (Role IN ('student', 'teacher', 'admin'))
);

-- Создание таблицы преподавателей
CREATE TABLE IF NOT EXISTS Teachers (
    TeacherID INTEGER PRIMARY KEY,
    FirstName TEXT,
    LastName TEXT,
    TeacherCode TEXT UNIQUE,
    PasswordHash TEXT,
    Role TEXT DEFAULT 'teacher' CHECK (Role IN ('student', 'teacher', 'admin'))
);

-- Создание таблицы предметов
CREATE TABLE IF NOT EXISTS Subjects (
    SubjectID INTEGER PRIMARY KEY,
    SubjectName TEXT
);

-- Создание таблицы, связывающей студентов с их дополнительными предметами
CREATE TABLE IF NOT EXISTS StudentSubjects (
    StudentID INTEGER,
    SubjectID INTEGER,
    PRIMARY KEY (StudentID, SubjectID),
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID) ON DELETE CASCADE,
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID) ON DELETE CASCADE
);

-- Создание таблицы, связывающей преподавателей с предметами
CREATE TABLE IF NOT EXISTS TeacherSubjects (
    TeacherID INTEGER,
    SubjectID INTEGER,
    PRIMARY KEY (TeacherID, SubjectID),
    FOREIGN KEY (TeacherID) REFERENCES Teachers(TeacherID) ON DELETE CASCADE,
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID) ON DELETE CASCADE
);

-- Создание таблицы групп
CREATE TABLE IF NOT EXISTS Groups (
    GroupID INTEGER PRIMARY KEY,
    GroupName TEXT,
    Schedule TEXT -- Расписание группы
);

-- Создание таблицы, связывающей студентов с группами
CREATE TABLE IF NOT EXISTS StudentGroups (
    StudentID INTEGER,
    GroupID INTEGER,
    PRIMARY KEY (StudentID, GroupID),
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID) ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID) ON DELETE CASCADE
);

-- Создание таблицы, связывающей преподавателей с группами
CREATE TABLE IF NOT EXISTS TeacherGroups (
    TeacherID INTEGER,
    GroupID INTEGER,
    PRIMARY KEY (TeacherID, GroupID),
    FOREIGN KEY (TeacherID) REFERENCES Teachers(TeacherID) ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID) ON DELETE CASCADE
);

-- Создание таблицы мест проведения занятий
CREATE TABLE IF NOT EXISTS Locations (
    LocationID INTEGER PRIMARY KEY,
    LocationName TEXT,
    InitialQRCode TEXT UNIQUE,
    CurrentQRCode TEXT,
    CodeChangeInterval INTEGER,
    CodeExpirationTime INTEGER,
    LastCodeUpdateTime INTEGER
);

-- Создание таблицы расписания
CREATE TABLE IF NOT EXISTS Schedule (
    ScheduleID INTEGER PRIMARY KEY,
    SubjectID INTEGER,
    GroupID INTEGER,
    TeacherID INTEGER,
    LocationID INTEGER,
    DayOfWeek TEXT,
    TimeSlot TEXT,
    SubjectName TEXT, -- Добавлено поле SubjectName
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID) ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID) ON DELETE CASCADE,
    FOREIGN KEY (TeacherID) REFERENCES Teachers(TeacherID) ON DELETE CASCADE,
    FOREIGN KEY (LocationID) REFERENCES Locations(LocationID) ON DELETE CASCADE
);

-- Создание таблицы посещаемости
CREATE TABLE IF NOT EXISTS Attendance (
    AttendanceID INTEGER PRIMARY KEY,
    StudentID INTEGER,
    ScheduleID INTEGER,
    Status TEXT,
    FirstScanTime INTEGER,
    SecondScanTime INTEGER,
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID) ON DELETE CASCADE,
    FOREIGN KEY (ScheduleID) REFERENCES Schedule(ScheduleID) ON DELETE CASCADE
);

-- Создание таблицы QR-кодов
CREATE TABLE IF NOT EXISTS QRCodes (
    QRCodeID INTEGER PRIMARY KEY,
    CodeValue TEXT UNIQUE,
    IsUsed BOOLEAN DEFAULT 0,
    ExpirationTime INTEGER
);

-- Создание таблицы для связи расписания с QR-кодами
CREATE TABLE IF NOT EXISTS ScheduleQRCodes (
    ScheduleQRCodeID INTEGER PRIMARY KEY,
    ScheduleID INTEGER,
    QRCodeID INTEGER,
    FOREIGN KEY (ScheduleID) REFERENCES Schedule(ScheduleID) ON DELETE CASCADE,
    FOREIGN KEY (QRCodeID) REFERENCES QRCodes(QRCodeID) ON DELETE CASCADE
);

-- Создание представления для журнала
CREATE VIEW IF NOT EXISTS AttendanceJournal AS
SELECT
    S.StudentID,
    S.FirstName,
    S.LastName,
    A.ScheduleID,
    S1.SubjectName,
    T.FirstName || ' ' || T.LastName AS TeacherName,
    A.Status,
    A.FirstScanTime,
    A.SecondScanTime

FROM
    Attendance A

    JOIN Students S ON A.StudentID = S.StudentID
    JOIN Schedule S1 ON A.ScheduleID = S1.ScheduleID
    JOIN Teachers T ON S1.TeacherID = T.TeacherID;
