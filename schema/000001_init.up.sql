CREATE TABLE users(
    id varchar(255) constraint users_pk primary key,
    login varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255)

)