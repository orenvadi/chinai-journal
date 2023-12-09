---Запрос для получения расписания занятий конкретного студента
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
    Students.StudentID = ?;
