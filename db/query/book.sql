-- name: GetBook :one
select * from books 
where id=$1 limit 1;

-- name: ListBook :many
select * from books 
order by id;

-- name: InsertBook :one
insert into books 
(title,author,is_read,year) 
values($1,$2, $3, $4) 
RETURNING *;

-- name: DeleteBook :one
delete from books 
where id = $1 
RETURNING *;

-- name: UpdateBook :one
update books 
set title=$2, author=$3, is_read=$4, year=$5 
where id=$1 
RETURNING *;

-- name: ToggleBookStatus :one
update books 
set is_read = not is_read
where id=$1 
RETURNING *;