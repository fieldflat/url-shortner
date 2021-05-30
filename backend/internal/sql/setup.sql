create table url_pairs
(
  id serial not null,
  short_url text,
  origin_url text,
  sid text
);

create table user_session
(
  id int primary key,
  sid text,
  expiration timestamp
);