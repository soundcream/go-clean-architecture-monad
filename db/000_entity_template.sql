
-- base entity
create table _base_entity
(
    id                          serial primary key,
    created_date                timestamp not null default now(),
    created_by                  varchar(50) not null,
    updated_date                timestamp null,
    updated_by                  varchar(50) null
);

-- soft delete entity
create table _deleter
(
    is_delete                   boolean not null default true,
    deleted_date                timestamp null,
    deleted_by                  varchar(50) null,
    deleted_reason              varchar(100) null,
    restore_date                timestamp null,
    restore_by                  varchar(50) null,
    restored_reason             varchar(100) null
);

-- activate entity
create table _activator
(
    is_active                   boolean not null default true,
    activated_date              timestamp null,
    activated_by                varchar(50) null,
    deactivated_date            timestamp null,
    deactivated_by              varchar(50) null
);