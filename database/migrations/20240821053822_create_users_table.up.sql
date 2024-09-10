create table users (
    id serial primary key ,
    email varchar(100) not null unique,
    name varchar(100) not null,
    password varchar(100) not null ,
    phone varchar(20) not null,
    token varchar(200) null ,
    created_at  bigint not null,
    updated_at  bigint not null,
    deleted_at  bigint null
);