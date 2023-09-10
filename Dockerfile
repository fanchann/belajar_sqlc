FROM golang:latest as BuildStage

WORKDIR /book_app
COPY . .
RUN go mod download
EXPOSE 9000
RUN go test -v -run=. latihan_sqlc/internal/repositories
RUN go test -v -run=. latihan_sqlc/internal/service
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/book_app cmd/books_app/main.go

# deploy stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /bin/book_app /app/
EXPOSE 9000
CMD ["./book_app"]