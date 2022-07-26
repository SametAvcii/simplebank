FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

COPY db/migration ./migration



EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]