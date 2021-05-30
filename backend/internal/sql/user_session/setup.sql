create table user_session
(
  id serial not null,
  sid text,
  expiration timestamp
);