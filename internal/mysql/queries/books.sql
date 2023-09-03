-- name: GetBookById :one
SELECT * FROM books
WHERE id = ? LIMIT 1;

-- name: GetAllBooks :many
SELECT * FROM books
ORDER BY title;

-- name: AddNewBook :execresult
INSERT INTO books (
  author,title
) VALUES (
  ?, ?
);

-- name: UpdateBook :exec
UPDATE books set title = ?,author = ? WHERE id = ?;  

-- name: DeleteAuthor :exec
DELETE FROM books
WHERE id = ?;