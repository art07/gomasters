CREATE TABLE IF NOT EXISTS users
(
    id         uuid PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name  varchar(40) NOT NULL,
    email      varchar(40) NOT NULL UNIQUE,
    age        int         NOT NULL,
    created    date        NOT NULL
);

INSERT INTO users (id, first_name, last_name, email, age, created)
VALUES (gen_random_uuid(), 'FirstUser', 'LastNameA', 'user1@gmail.com', 20, now()),
       (gen_random_uuid(), 'SecondUser', 'LastNameB', 'user2@gmail.com', 21, now()),
       (gen_random_uuid(), 'ThirdUser', 'LastNameC', 'user3@gmail.com', 22, now());


CREATE TABLE IF NOT EXISTS admins
(
    id         uuid PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name  varchar(40) NOT NULL,
    email      varchar(40) NOT NULL UNIQUE,
    age        int         NOT NULL,
    created    date        NOT NULL
);

INSERT INTO admins (id, first_name, last_name, email, age, created)
VALUES (gen_random_uuid(), 'SuperUser', 'SuperLastName', 'admin1@gmail.com', 50, now());