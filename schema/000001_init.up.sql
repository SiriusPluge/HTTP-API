CREATE TABLE person
(
    id serial not null unique,
    name varchar(256) not null,
    last_name varchar(256) not null,
    phone int not null
)