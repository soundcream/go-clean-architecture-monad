
set TIMEZONE = 'Asia/Bangkok';

create table users
(
    id                          serial primary key,
    name                        varchar(50) not null,
    username                    varchar(50) not null,
    email                       varchar(100) not null,
    point                       int null,
    user_group_id               int null,
    created_date                timestamp not null default now(),
    created_by                  varchar(50) not null,
    updated_date                timestamp null,
    updated_by                  varchar(50) null
);

create table user_groups
(
    id                          serial primary key,
    name                        varchar(50) not null,
    level                       int not null default 0,
    is_active                   boolean not null default true,
    activated_date              timestamp null,
    activated_by                varchar(50) null,
    deactivated_date            timestamp null,
    deactivated_by              varchar(50) null,
    created_date                timestamp not null default now(),
    created_by                  varchar(50) not null,
    updated_date                timestamp null,
    updated_by                  varchar(50) null
);
