
-- Запрос для отметки QR-кода как использованного
UPDATE
    QRCodes
SET
    IsUsed = 1
WHERE
    QRCodeID = ?;
