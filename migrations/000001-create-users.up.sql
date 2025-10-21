-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
   id            serial primary key,
   first_name    varchar(100) not null,
   last_name     varchar(100) not null,
   email         varchar(255) unique not null,
   password      varchar(255) not null,
   is_shop_owner boolean default false,
   created_at    timestamp with time zone default current_timestamp,
   updated_at    timestamp with time zone default current_timestamp
);