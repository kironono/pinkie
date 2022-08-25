FROM golang:1.18.5 as develop

RUN apt-get update && apt-get install -y \
    default-mysql-client \
 && apt-get clean

WORKDIR /app

RUN go install github.com/rubenv/sql-migrate/...@v1.1.2
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
