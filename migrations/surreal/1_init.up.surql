DEFINE TABLE Student SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Student FLEXIBLE TYPE object;
DEFINE FIELD StudentCode    ON TABLE Student TYPE string;
DEFINE INDEX StudentCodeIndex     ON TABLE Student COLUMNS StudentCode UNIQUE;
DEFINE FIELD Email          ON TABLE Student TYPE string;
DEFINE INDEX StudentEmailIndex    ON TABLE Student COLUMNS Email UNIQUE;
DEFINE FIELD PasswordHash   ON TABLE Student TYPE string;
DEFINE FIELD Subjects       ON TABLE Student TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Student TYPE array<record<Group>>;



DEFINE TABLE Teacher SCHEMAFULL; 

DEFINE FIELD Name           ON TABLE Teacher FLEXIBLE TYPE object;
DEFINE FIELD TeacherCode    ON TABLE Teacher TYPE string;
DEFINE INDEX TeacherCodeIndex    ON TABLE Teacher COLUMNS TeacherCode UNIQUE;
DEFINE FIELD Email          ON TABLE Teacher TYPE string;
DEFINE INDEX TeacherEmailIndex    ON TABLE Teacher COLUMNS Email UNIQUE;
DEFINE FIELD PasswordHash   ON TABLE Teacher TYPE string;
DEFINE FIELD Subjects       ON TABLE Teacher TYPE array<record<Subject>>;
DEFINE FIELD Groups         ON TABLE Teacher TYPE array<record<Group>>;


DEFINE TABLE Subject SCHEMAFULL; 

DEFINE FIELD Name ON TABLE Subject TYPE string;



DEFINE TABLE Group SCHEMAFULL; 

DEFINE FIELD Name       ON TABLE Group TYPE string;
DEFINE INDEX NameIndex  ON TABLE Group COLUMNS Name UNIQUE;



DEFINE TABLE Location SCHEMAFULL; 

DEFINE FIELD Name               ON TABLE Location TYPE string;
DEFINE INDEX NameIndex          ON TABLE Location COLUMNS Name UNIQUE;
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

DEFINE FIELD Subject              ON TABLE Schedule   TYPE record<Subject>;
DEFINE FIELD Student              ON TABLE Attendance TYPE record<Student> ;
DEFINE FIELD Schedule             ON TABLE Attendance TYPE record<Schedule>;
DEFINE FIELD IsRoomCodeScanned    ON TABLE Attendance TYPE bool;
DEFINE FIELD IsConfirmCodeScanned ON TABLE Attendance TYPE bool;
DEFINE FIELD IsAttended           ON TABLE Attendance TYPE bool;



DEFINE TABLE QrCode SCHEMAFULL; 

DEFINE FIELD Code         ON TABLE QrCode TYPE string;
DEFINE FIELD FirstUseTime ON TABLE QrCode TYPE datetime;
