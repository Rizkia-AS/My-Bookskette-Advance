create table "books" (
    "id" bigserial not null primary key,
    "title" varchar(45) not null, 
    "author" varchar(45) not null, 
    "is_read" boolean not null, 
    "year" bigint not null,
    "created_at" timestamptz not null default (now())
);

create index on "books" ("author");