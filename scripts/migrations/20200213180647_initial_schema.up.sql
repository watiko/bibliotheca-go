create table users
(
  user_id    serial       not null primary key,
  name       varchar(128) not null,
  updated_at timestamptz  not null,
  created_at timestamptz  not null
);

create table groups
(
  group_id   serial       not null primary key,
  name       varchar(128) not null,
  updated_at timestamptz  not null,
  created_at timestamptz  not null
);

create table user_group_memberships
(
  user_id  int not null
    references users
      on delete cascade,
  group_id int not null
    references groups
      on delete cascade,
  primary key (user_id, group_id)
);

create table book_shelves
(
  book_shelf_id serial       not null primary key,
  group_id      int          not null
    references groups
      on delete cascade,
  name          varchar(128) not null,
  updated_at    timestamptz  not null,
  created_at    timestamptz  not null
);

create table books
(
  book_id       serial       not null primary key,
  book_shelf_id int          not null
    references book_shelves (book_shelf_id)
      on delete cascade,
  title         varchar(128) not null,
  borrowed_by   int
    references users
      on delete set null,
  isbn          varchar(64),
  updated_at    timestamptz  not null,
  created_at    timestamptz  not null
);
