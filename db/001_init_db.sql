

create table users
(
    id       serial constraint users_pk primary key,
    name     varchar(50),
    username varchar(50),
    email    varchar(100),
    point    int
);

create table customer_groups
(
    id              serial constraint users_pk primary key,
    name            varchar(50),
    level           int,
    is_active       boolean default true,
    created_date    timestamp default now()
);


