

## Получение кода аудитории на основе id занятия

**напиши мне пожалуйста sql запрос для того чтобы получить qr код аудитории
в котором проводится занятие и qr коды для подтверждения присутствия
которое должен транслировать преподаватель**


```sql
-- Запрос для получения QR-кода аудитории
SELECT
    LocationID,
    LocationName,
    InitialQRCode AS AuditoriumQRCode
FROM
    Locations
WHERE
    LocationID IN (
        SELECT
            LocationID
        FROM
            Schedule
        WHERE
            ScheduleID = :your_schedule_id
    );

-- Запрос для получения QR-кодов для подтверждения присутствия
SELECT
    QRCodes.QRCodeID,
    QRCodes.CodeValue AS ConfirmationQRCode
FROM
    ScheduleQRCodes
JOIN
    QRCodes ON ScheduleQRCodes.QRCodeID = QRCodes.QRCodeID
WHERE
    ScheduleQRCodes.ScheduleID = :your_schedule_id;
```

## Получение идентификаторов всех занятий студента

**напиши мне пожалуйста sql запрос для Получения расписания 
занятий конкретного студента**

```sql
-- Запрос для получения расписания занятий конкретного студента
SELECT
    Schedule.ScheduleID,
    Subjects.SubjectName,
    Teachers.FirstName || ' ' || Teachers.LastName AS TeacherName,
    Groups.GroupName,
    Locations.LocationName,
    Schedule.DayOfWeek,
    Schedule.TimeSlot
FROM
    Attendance
JOIN
    Students ON Attendance.StudentID = Students.StudentID
JOIN
    Schedule ON Attendance.ScheduleID = Schedule.ScheduleID
JOIN
    Subjects ON Schedule.SubjectID = Subjects.SubjectID
JOIN
    Teachers ON Schedule.TeacherID = Teachers.TeacherID
JOIN
    Groups ON Schedule.GroupID = Groups.GroupID
JOIN
    Locations ON Schedule.LocationID = Locations.LocationID
WHERE
    Students.StudentID = :your_student_id;
```



## Получение идентификаторов всех занятий преподователя

**напиши мне пожалуйста sql запрос для Получения расписания занятий конкретного преподавателя**

```sql
-- Запрос для получения расписания занятий конкретного преподавателя
SELECT
    Schedule.ScheduleID,
    Subjects.SubjectName,
    Students.FirstName || ' ' || Students.LastName AS StudentName,
    Groups.GroupName,
    Locations.LocationName,
    Schedule.DayOfWeek,
    Schedule.TimeSlot
FROM
    Attendance
JOIN
    Schedule ON Attendance.ScheduleID = Schedule.ScheduleID
JOIN
    Subjects ON Schedule.SubjectID = Subjects.SubjectID
JOIN
    Teachers ON Schedule.TeacherID = Teachers.TeacherID
JOIN
    Groups ON Schedule.GroupID = Groups.GroupID
JOIN
    Locations ON Schedule.LocationID = Locations.LocationID
JOIN
    Students ON Attendance.StudentID = Students.StudentID
WHERE
    Teachers.TeacherID = :your_teacher_id;
```




## Создание кодов для подтверждения в конце занятия

**напиши мне пожалуйста sql запрос для создания нового занятия в расписании таким образом чтобы также создавались коды для подтверждения в конце занятия**

```sql
-- Шаг 2: Создание нового занятия в расписании
INSERT INTO Schedule (SubjectID, GroupID, TeacherID, LocationID, DayOfWeek, TimeSlot, SubjectName)
VALUES (1, 1, 1, 1, 1, '22:22:22', 'sanjar');

-- Получение ID нового занятия
SELECT last_insert_rowid() AS NewScheduleID;

-- Шаг 2: Создание QR-кодов для подтверждения присутствия на новом занятии
INSERT INTO QRCodes (CodeValue, IsUsed, ExpirationTime)
VALUES ('123456789', 0, '11:11:11');

INSERT INTO QRCodes (CodeValue, IsUsed, ExpirationTime)
VALUES ('987654321', 0, '22:22:22');

-- Получение ID новых QR-кодов
SELECT last_insert_rowid() AS FirstQRCodeID, '123456789' AS FirstQRCodeValue;
SELECT last_insert_rowid() AS SecondQRCodeID, '987654321' AS SecondQRCodeValue;

-- Шаг 3: Связывание QR-кодов с новым занятием
INSERT INTO ScheduleQRCodes (ScheduleID, QRCodeID)
VALUES (51, 1);

INSERT INTO ScheduleQRCodes (ScheduleID, QRCodeID)
VALUES (51, 1);
```


