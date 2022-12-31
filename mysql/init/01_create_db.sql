CREATE DATABASE echo;
CREATE USER echo IDENTIFIED BY 'echo';
GRANT ALL PRIVILEGES ON echo.* TO 'echo'@'%';
USE echo

DROP TABLE IF EXISTS member;
CREATE TABLE member
(
  id               VARCHAR(64),
  last_access_date DATE
);
