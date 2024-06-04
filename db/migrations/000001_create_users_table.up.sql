CREATE TABLE IF NOT EXISTS users (
    user_id serial primary key,
    username varchar(255) UNIQUE,
    password varchar(255)
);