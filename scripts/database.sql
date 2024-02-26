-- DATABASE DESIGN MOVED TO /data/design.txt


\c notetaker-ntc02
--
--DROP TABLE IF EXISTS notes;
--CREATE SEQUENCE notes_id_seq;
--CREATE TABLE notes (
--    id  int PRIMARY KEY NOT NULL DEFAULT nextval('notes_id_seq'),
--    title varchar(200),
--    content text,
--    created_date timestamp,
--    updated_date timestamp
--);
--
--ALTER SEQUENCE notes_id_seq
--OWNED BY notes.id;
--
--DROP TABLE IF EXISTS tags;
--CREATE SEQUENCE tags_id_seq;
--CREATE TABLE tags (
--    id int PRIMARY KEY,
--    title varchar(200),
--    note_id int,
--   
--    CONSTRAINT fk_notes_tags FOREIGN KEY (note_id) REFERENCES tags(id)
--);
--ALTER SEQUENCE tags_id_seq
--OWNED BY tags.id;


--DROP TABLE IF EXISTS sessions;
--CREATE TABLE sessions (
--    token char(43) PRIMARY KEY,
--    data bytea not null,
--    expiry timestamp not null
--);

DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id int PRIMARY KEY NOT NULL DEFAULT nextval('users_id_seq'),
    name varchar(255) not null,
    email varchar(255) not null,
    hashed_password char(60) not null,
    created timestamp not null
);
ALTER SEQUENCE users_id_seq
OWNED BY users.id;
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE(email);


