CREATE TABLE users
(
    id       bigserial not null unique,
    email    varchar unique,
    password varchar
);

CREATE SEQUENCE users_seq