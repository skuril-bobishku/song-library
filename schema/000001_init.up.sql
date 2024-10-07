CREATE TABLE songs
(
    s_id serial primary key,
    group_name varchar(255) not null,
    song_name varchar(255) not null,
    release_date date,
    text text,
    link varchar(255)
);