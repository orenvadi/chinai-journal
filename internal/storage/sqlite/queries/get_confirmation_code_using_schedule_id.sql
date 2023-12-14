---Запрос для получения QR-кодов для подтверждения присутствия
SELECT
    QRCodes.QRCodeID,
    QRCodes.CodeValue AS ConfirmationQRCode,
    QRCodes.IsUsed
FROM
    ScheduleQRCodes
JOIN
    QRCodes ON ScheduleQRCodes.QRCodeID = QRCodes.QRCodeID
WHERE
    ScheduleQRCodes.ScheduleID = ?;
