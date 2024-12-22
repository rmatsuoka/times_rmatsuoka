create table channels (
  id integer primary key not null,
  code varchar(256) unique not null,
  created_at timestamp default CURRENT_TIMESTAMP not null
);

create table members (
  channel_id integer references channels(id) not null,
  user_id integer references users(id) not null,
  role integer not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  unique(channel_id, user_id)
);

create index members_channel_id on members (channel_id);
create index members_user_id on members (user_id);
