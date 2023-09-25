FROM golang:1.20.1-bullseye

WORKDIR /app
COPY ./ /app

RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 3000

# CMD ["tail", "-f", "/dev/null"]
CMD ["sleep", "infinity"]
