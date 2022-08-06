CREATE TABLE users
(
    id   BIGINT      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(32) NOT NULL
);


CREATE TABLE user_details
(
    id    BIGINT       NOT NULL PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE
);