drop table posts cascade if exists;
drop table comments if exists;

create table posts (
    id serial primary key, -- 主キー
    content text,
    author varchar(255)
);

create table comments (
    id serial primary key, -- 主キー
    content text,
    author varchar(255),
    post_id integer references posts(id) -- 外部キー
)