alter table users
    add email varchar(128) not null;

create unique index users_email_uindex
    on users (email);
