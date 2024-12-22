create table messages (
  id integer primary key not null,
  user_id integer references users(id) not null,
  channel_id integer references channels(id) not null,
  text text not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  modified_at timestamp default CURRENT_TIMESTAMP
);

create index messages_created_at on messages (created_at);
create index messages_channel_id on messages (channel_id);
