create domain ulid as varchar(26);

create table users
(
    user_id    ulid         not null primary key,
    name       varchar(128) not null,
    email      varchar(128) not null unique,
    updated_at timestamptz  not null,
    created_at timestamptz  not null
);

create table groups
(
    group_id   ulid         not null primary key,
    name       varchar(128) not null,
    updated_at timestamptz  not null,
    created_at timestamptz  not null
);

create table user_group_memberships
(
    user_id  ulid not null
        references users
            on delete cascade,
    group_id ulid not null
        references groups
            on delete cascade,
    primary key (user_id, group_id)
);

create table bookshelves
(
    bookshelf_id  ulid         not null primary key,
    group_id      ulid         not null
        references groups
            on delete cascade,
    name          varchar(128) not null,
    updated_at    timestamptz  not null,
    created_at    timestamptz  not null
);

create table books
(
    book_id       ulid         not null primary key,
    bookshelf_id  ulid         not null
        references bookshelves (bookshelf_id)
            on delete cascade,
    title         varchar(128) not null,
    borrowed_by   ulid
        references users
            on delete set null,
    isbn          varchar(64),
    updated_at    timestamptz  not null,
    created_at    timestamptz  not null
);
