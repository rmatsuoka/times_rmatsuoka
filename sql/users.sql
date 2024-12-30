create table users (
  id integer primary key not null,
  code varchar(256) not null unique,
  name varchar(256) not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  updated_at timestamp default CURRENT_TIMESTAMP not null
);

create table deleted_users (
  user_id integer not null,
  code varchar(256) not null,
  name varchar(256) not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  deleted_at timestamp default CURRENT_TIMESTAMP not null
);

create table usercodes (
  code varchar(256) unique not null
);
