---Запрос для получения QR-кода аудитории
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
            ScheduleID = ?
    );
