
DEFINE TABLE Student SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Student TYPE object;
DEFINE FIELD Student_code   ON TABLE Student TYPE string;
DEFINE FIELD Password_hash  ON TABLE Student TYPE string;
DEFINE FIELD Subjects       ON TABLE Student TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Student TYPE array<record<Group>>;


create
  Student:tester
    set
        Name.Last       = "narashi",
        Name.First      = "kuga",
        Name.Patronimic = "baktybekovich",
        Name.Full       = string::join(' ', Name.Last, Name.First, Name.Patronimic),

        Student_code    = "WIN-1-22-472",
        Password_hash   = "fwoieiwocdkfekj20934ge",
        Subjects        = [Subject:tester1],
        Groups          = [Group:tester]
;


DEFINE TABLE Teacher SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Teacher TYPE object;
DEFINE FIELD Teacher_code   ON TABLE Teacher TYPE string;
DEFINE FIELD Password_hash  ON TABLE Teacher TYPE string;

DEFINE FIELD Subjects       ON TABLE Teacher TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Teacher TYPE array<record<Group>>;


create
  Teacher:tester
    set
        Name.Last       = "karabakirov",
        Name.First      = "kubat",
        Name.Patronimic = "rymbekovich",
        Name.Full       = string::join(' ', Name.Last, Name.First, Name.Patronimic),
        Teacher_code    = "karabakirov",
        Password_hash   = "fwoieiwocdkfekj20934ge",
        Subjects        = [Subject:tester1, Subject:tester2],
        Groups          = [Group:tester1]
;



DEFINE TABLE Subject SCHEMAFULL; 

DEFINE FIELD Name ON TABLE Subject TYPE string;

create Subject:tester1
    set 
        Name = "Operating Systems"
;


create Subject:tester2
    set 
        Name = "Math"
;






DEFINE TABLE Group SCHEMAFULL; 

DEFINE FIELD Name ON TABLE Group TYPE string;

create Group:tester
    set
        Name = "WIN-1-22"
;




DEFINE TABLE Location SCHEMAFULL; 

DEFINE FIELD Name               ON TABLE Location TYPE string;
DEFINE FIELD PrelimConfirmCode  ON TABLE Location TYPE string;

create Location:tester
    set 
        Name                = "222",
        PrelimConfirmCode   = "vownlekir-3495045(&$(&349))"
;






DEFINE TABLE Schedule SCHEMAFULL; 


DEFINE FIELD Subject  ON TABLE Schedule TYPE record<Subject>;
DEFINE FIELD Group    ON TABLE Schedule TYPE record<Group>;
DEFINE FIELD Teacher  ON TABLE Schedule TYPE record<Teacher>;
DEFINE FIELD Location ON TABLE Schedule TYPE record<Location>;
DEFINE FIELD Dateslot ON TABLE Schedule TYPE datetime;
DEFINE FIELD Timeslot ON TABLE Schedule TYPE datetime;
DEFINE FIELD QrCodes  ON TABLE Schedule TYPE array<record<QrCode>>;

create Schedule:tester
    set 
        Subject     = Subject:tester1,
        Group       = Group:tester,
        Teacher     = Teacher:tester,
        Location    = Location:tester,
        Dateslot    = time::now(),
        Timeslot    = time::now(),
        QrCodes     = [QrCode:tester1, QrCode:tester2]
;







DEFINE TABLE Attendance SCHEMAFULL; 

DEFINE FIELD Subject  ON TABLE Schedule TYPE record<Subject>;


DEFINE FIELD Student              ON TABLE Attendance TYPE record<Student> ;
DEFINE FIELD Schedule             ON TABLE Attendance TYPE record<Schedule>;
DEFINE FIELD IsRoomCodeScanned    ON TABLE Attendance TYPE bool;
DEFINE FIELD IsConfirmCodeScanned ON TABLE Attendance TYPE bool;
DEFINE FIELD IsAttended           ON TABLE Attendance TYPE bool;

create Attendance:tester
    set 
        Student              = Student:tester,
        Schedule             = Schedule:tester,
        IsRoomCodeScanned    = false,
        IsConfirmCodeScanned = false,
        IsAttended           = IsRoomCodeScanned && IsConfirmCodeScanned
;


DEFINE TABLE QrCode SCHEMAFULL; 


DEFINE FIELD Code   ON TABLE QrCode TYPE string;
DEFINE FIELD IsUsed ON TABLE QrCode TYPE bool;


create QrCode:tester1
    set 
        Code    = "foiwejfoierr",
        IsUsed  = false
;

create QrCode:tester2
    set 
        Code    = "veorjoeirofer",
        IsUsed  = false
;






function attendance_journal() {
  return select {
    S.id as Student_id,
    S.first_name,
    S.last_name,
    A.schedule_id,
    S1.Subject_name,
    concat(T.first_name, ' ', T.last_name) as teacher_name,
    A.status,
    A.first_scan_time,
    A.second_scan_time
  }
  from attendance A
  join Student S on A.Student_id = S.id
  join schedule S1 on A.schedule_id = S1.id
  join teacher T on S1.teacher_id = T.id
}


-- update Student:arxqjwkxrq1sfgg7km3v set Subjects = [Subject:lte0fk9n9atcgt5dl0tu];

-- select * from Student where array::find_index(Subjects, 'Subject:lte0fk9n9atcgt5dl0tu') != null;


-- update Student:arxqjwkxrq1sfgg7km3v set Subjects = array::append(Subjects, Subject:04800vi0mveg4berlo6h);


-- update Teacher:s0c8199kzme5rzadmxbf set Subjects = [Subject:y7s93exo00oek3rbrfxr];
