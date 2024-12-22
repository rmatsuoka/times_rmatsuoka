create table users (
  id integer primary key not null,
  code varchar(256) not null unique,
  name varchar(256) not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  updated_at timestamp default CURRENT_TIMESTAMP not null
);

