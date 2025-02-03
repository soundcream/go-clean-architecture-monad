


set TIMEZONE = 'Asia/Bangkok';

create table users
(
    id       serial primary key,
    name     varchar(50) not null,
    username varchar(50) not null,
    email    varchar(100) not null,
    point    int null,
    user_group_id int null
);

create table user_groups
(
    id              serial primary key,
    name            varchar(50) not null,
    level           int not null default 0,
    is_active       boolean not null default true,
    created_date    timestamp not null default now()
);

