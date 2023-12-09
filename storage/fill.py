import random
import sqlite3

from faker import Faker

# Initialize Faker
fake = Faker()


# Function to generate random student data
def generate_student():
    return {
        "FirstName": fake.first_name(),
        "LastName": fake.last_name(),
        "StudentCode": fake.unique.random_number(digits=8),
        "PasswordHash": fake.sha256(),
        "Role": "student",
    }


# Function to generate random teacher data
def generate_teacher():
    return {
        "FirstName": fake.first_name(),
        "LastName": fake.last_name(),
        "TeacherCode": fake.unique.random_number(digits=8),
        "PasswordHash": fake.sha256(),
        "Role": "teacher",
    }


# Function to generate random subject data
def generate_subject():
    return {"SubjectName": fake.word()}


# Function to generate random group data
def generate_group():
    return {"GroupName": fake.word(), "Schedule": fake.sentence()}


# Function to generate random location data
def generate_location():
    return {
        "LocationName": fake.word(),
        "InitialQRCode": fake.uuid4(),
        "CurrentQRCode": fake.uuid4(),
        "CodeChangeInterval": random.randint(1, 10),
        "CodeExpirationTime": random.randint(30, 1800),
        "LastCodeUpdateTime": fake.date_time_this_month().timestamp(),
    }


# Function to generate random schedule data
def generate_schedule(subject_id, group_id, teacher_id, location_id):
    return {
        "SubjectID": subject_id,
        "GroupID": group_id,
        "TeacherID": teacher_id,
        "LocationID": location_id,
        "DayOfWeek": fake.random_element(
            elements=("Monday", "Tuesday", "Wednesday", "Thursday", "Friday")
        ),
        "TimeSlot": fake.time(),
        "SubjectName": fake.word(),
    }


# Function to generate random attendance data
def generate_attendance(student_id, schedule_id):
    return {
        "StudentID": student_id,
        "ScheduleID": schedule_id,
        "Status": fake.random_element(elements=("Present", "Absent")),
        "FirstScanTime": fake.date_time_this_month().timestamp(),
        "SecondScanTime": fake.date_time_this_month().timestamp(),
    }


# Connect to the SQLite database
conn = sqlite3.connect("./storage.db")
cursor = conn.cursor()

# Generate and insert random student data
for _ in range(50):
    student_data = generate_student()
    cursor.execute(
        "INSERT INTO Students (FirstName, LastName, StudentCode, PasswordHash, Role) VALUES (?, ?, ?, ?, ?)",
        (
            student_data["FirstName"],
            student_data["LastName"],
            student_data["StudentCode"],
            student_data["PasswordHash"],
            student_data["Role"],
        ),
    )

# Generate and insert random teacher data
for _ in range(10):
    teacher_data = generate_teacher()
    cursor.execute(
        "INSERT INTO Teachers (FirstName, LastName, TeacherCode, PasswordHash, Role) VALUES (?, ?, ?, ?, ?)",
        (
            teacher_data["FirstName"],
            teacher_data["LastName"],
            teacher_data["TeacherCode"],
            teacher_data["PasswordHash"],
            teacher_data["Role"],
        ),
    )

# Generate and insert random subject data
for _ in range(20):
    subject_data = generate_subject()
    cursor.execute(
        "INSERT INTO Subjects (SubjectName) VALUES (?)", (subject_data["SubjectName"],)
    )

# Generate and insert random group data
for _ in range(5):
    group_data = generate_group()
    cursor.execute(
        "INSERT INTO Groups (GroupName, Schedule) VALUES (?, ?)",
        (group_data["GroupName"], group_data["Schedule"]),
    )

# Generate and insert random location data
for _ in range(10):
    location_data = generate_location()
    cursor.execute(
        "INSERT INTO Locations (LocationName, InitialQRCode, CurrentQRCode, CodeChangeInterval, CodeExpirationTime, LastCodeUpdateTime) VALUES (?, ?, ?, ?, ?, ?)",
        (
            location_data["LocationName"],
            location_data["InitialQRCode"],
            location_data["CurrentQRCode"],
            location_data["CodeChangeInterval"],
            location_data["CodeExpirationTime"],
            location_data["LastCodeUpdateTime"],
        ),
    )

# Get IDs for subjects, groups, teachers, and locations
subject_ids = [
    row[0] for row in cursor.execute("SELECT SubjectID FROM Subjects").fetchall()
]
group_ids = [row[0] for row in cursor.execute("SELECT GroupID FROM Groups").fetchall()]
teacher_ids = [
    row[0] for row in cursor.execute("SELECT TeacherID FROM Teachers").fetchall()
]
location_ids = [
    row[0] for row in cursor.execute("SELECT LocationID FROM Locations").fetchall()
]

# Generate and insert random schedule data
for _ in range(50):
    schedule_data = generate_schedule(
        random.choice(subject_ids),
        random.choice(group_ids),
        random.choice(teacher_ids),
        random.choice(location_ids),
    )
    cursor.execute(
        "INSERT INTO Schedule (SubjectID, GroupID, TeacherID, LocationID, DayOfWeek, TimeSlot, SubjectName) VALUES (?, ?, ?, ?, ?, ?, ?)",
        (
            schedule_data["SubjectID"],
            schedule_data["GroupID"],
            schedule_data["TeacherID"],
            schedule_data["LocationID"],
            schedule_data["DayOfWeek"],
            schedule_data["TimeSlot"],
            schedule_data["SubjectName"],
        ),
    )

# Get IDs for students and schedules
student_ids = [
    row[0] for row in cursor.execute("SELECT StudentID FROM Students").fetchall()
]
schedule_ids = [
    row[0] for row in cursor.execute("SELECT ScheduleID FROM Schedule").fetchall()
]

# Generate and insert random attendance data
for _ in range(200):
    attendance_data = generate_attendance(
        random.choice(student_ids), random.choice(schedule_ids)
    )
    cursor.execute(
        "INSERT INTO Attendance (StudentID, ScheduleID, Status, FirstScanTime, SecondScanTime) VALUES (?, ?, ?, ?, ?)",
        (
            attendance_data["StudentID"],
            attendance_data["ScheduleID"],
            attendance_data["Status"],
            attendance_data["FirstScanTime"],
            attendance_data["SecondScanTime"],
        ),
    )

# Commit changes and close the connection
conn.commit()
conn.close()
