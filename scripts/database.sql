-- DATABASE DESIGN

-- ┌─────────────────────┐
-- │notes                │          ┌─────────────┐
-- ├─────────────────────┤          │tags         │
-- │id int               │1        n├─────────────┤
-- │title string         ├──────────┤id int       │
-- │content string       │          │name string  │
-- │created_date datetime│          └─────────────┘
-- │updated_date datetime│
-- └─────────────────────┘

DROP DATABASE IF EXISTS notetaker;

CREATE DATABASE notetaker;

\c notetaker;

CREATE SEQUENCE notes_id_seq;
CREATE TABLE notes (
    id  int PRIMARY KEY NOT NULL DEFAULT nextval('notes_id_seq'),
    title varchar(200),
    content text,
    created_date timestamp,
    updated_date timestamp
);

ALTER SEQUENCE notes_id_seq
OWNED BY notes.id;

CREATE SEQUENCE tags_id_seq;
CREATE TABLE tags (
    id int PRIMARY KEY,
    title varchar(200),
    note_id int,
   
    CONSTRAINT fk_notes_tags FOREIGN KEY (note_id) REFERENCES tags(id)
);
ALTER SEQUENCE tags_id_seq
OWNED BY tags.id;

