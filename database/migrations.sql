create table if not exists books (
    id serial primary key,
    title varchar(255) not null,
    author varchar(255) not null
)
