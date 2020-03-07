-- This file was generated by dump_schema.sh.
-- DO NOT MANUALLY MODIFY THIS.

CREATE TABLE bookshelves (
    bookshelf_id integer NOT NULL,
    group_id integer NOT NULL,
    name character varying(128) NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE SEQUENCE book_shelves_book_shelf_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE book_shelves_book_shelf_id_seq OWNED BY bookshelves.bookshelf_id;

CREATE TABLE books (
    book_id integer NOT NULL,
    bookshelf_id integer NOT NULL,
    title character varying(128) NOT NULL,
    borrowed_by integer,
    isbn character varying(64),
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE SEQUENCE books_book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE books_book_id_seq OWNED BY books.book_id;

CREATE TABLE groups (
    group_id integer NOT NULL,
    name character varying(128) NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE SEQUENCE groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE groups_group_id_seq OWNED BY groups.group_id;

CREATE TABLE user_group_memberships (
    user_id integer NOT NULL,
    group_id integer NOT NULL
);

CREATE TABLE users (
    user_id integer NOT NULL,
    name character varying(128) NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    email character varying(128) NOT NULL
);

CREATE SEQUENCE users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE users_user_id_seq OWNED BY users.user_id;

ALTER TABLE ONLY books ALTER COLUMN book_id SET DEFAULT nextval('books_book_id_seq'::regclass);

ALTER TABLE ONLY bookshelves ALTER COLUMN bookshelf_id SET DEFAULT nextval('book_shelves_book_shelf_id_seq'::regclass);

ALTER TABLE ONLY groups ALTER COLUMN group_id SET DEFAULT nextval('groups_group_id_seq'::regclass);

ALTER TABLE ONLY users ALTER COLUMN user_id SET DEFAULT nextval('users_user_id_seq'::regclass);

ALTER TABLE ONLY bookshelves
    ADD CONSTRAINT book_shelves_pkey PRIMARY KEY (bookshelf_id);

ALTER TABLE ONLY books
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);

ALTER TABLE ONLY groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (group_id);

ALTER TABLE ONLY user_group_memberships
    ADD CONSTRAINT user_group_memberships_pkey PRIMARY KEY (user_id, group_id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);

CREATE UNIQUE INDEX users_email_uindex ON users USING btree (email);

ALTER TABLE ONLY bookshelves
    ADD CONSTRAINT book_shelves_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(group_id) ON DELETE CASCADE;

ALTER TABLE ONLY books
    ADD CONSTRAINT books_book_shelf_id_fkey FOREIGN KEY (bookshelf_id) REFERENCES bookshelves(bookshelf_id) ON DELETE CASCADE;

ALTER TABLE ONLY books
    ADD CONSTRAINT books_borrowed_by_fkey FOREIGN KEY (borrowed_by) REFERENCES users(user_id) ON DELETE SET NULL;

ALTER TABLE ONLY user_group_memberships
    ADD CONSTRAINT user_group_memberships_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(group_id) ON DELETE CASCADE;

ALTER TABLE ONLY user_group_memberships
    ADD CONSTRAINT user_group_memberships_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE;

