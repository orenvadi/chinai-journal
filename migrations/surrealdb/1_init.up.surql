DEFINE TABLE Student SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Student TYPE object;
DEFINE FIELD Student_code   ON TABLE Student TYPE string;
DEFINE FIELD Password_hash  ON TABLE Student TYPE string;
DEFINE FIELD Subjects       ON TABLE Student TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Student TYPE array<record<Group>>;





DEFINE TABLE Teacher SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Teacher TYPE object;
DEFINE FIELD Teacher_code   ON TABLE Teacher TYPE string;
DEFINE FIELD Password_hash  ON TABLE Teacher TYPE string;

DEFINE FIELD Subjects       ON TABLE Teacher TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Teacher TYPE array<record<Group>>;







DEFINE TABLE Subject SCHEMAFULL; 

DEFINE FIELD Name ON TABLE Subject TYPE string;












DEFINE TABLE Group SCHEMAFULL; 

DEFINE FIELD Name ON TABLE Group TYPE string;

create Group:tester
    set
        Name = "WIN-1-22"
;







DEFINE TABLE Location SCHEMAFULL; 

DEFINE FIELD Name               ON TABLE Location TYPE string;
DEFINE FIELD PrelimConfirmCode  ON TABLE Location TYPE string;








DEFINE TABLE Schedule SCHEMAFULL; 


DEFINE FIELD Subject  ON TABLE Schedule TYPE record<Subject>;
DEFINE FIELD Group    ON TABLE Schedule TYPE record<Group>;
DEFINE FIELD Teacher  ON TABLE Schedule TYPE record<Teacher>;
DEFINE FIELD Location ON TABLE Schedule TYPE record<Location>;
DEFINE FIELD Dateslot ON TABLE Schedule TYPE datetime;
DEFINE FIELD Timeslot ON TABLE Schedule TYPE datetime;
DEFINE FIELD QrCodes  ON TABLE Schedule TYPE array<record<QrCode>>;








DEFINE TABLE Attendance SCHEMAFULL; 

DEFINE FIELD Subject  ON TABLE Schedule TYPE record<Subject>;


DEFINE FIELD Student              ON TABLE Attendance TYPE record<Student> ;
DEFINE FIELD Schedule             ON TABLE Attendance TYPE record<Schedule>;
DEFINE FIELD IsRoomCodeScanned    ON TABLE Attendance TYPE bool;
DEFINE FIELD IsConfirmCodeScanned ON TABLE Attendance TYPE bool;
DEFINE FIELD IsAttended           ON TABLE Attendance TYPE bool;











DEFINE TABLE QrCode SCHEMAFULL; 


DEFINE FIELD Code   ON TABLE QrCode TYPE string;
DEFINE FIELD IsUsed ON TABLE QrCode TYPE bool;
