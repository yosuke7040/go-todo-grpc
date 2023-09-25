# https://docs.docker.com/language/golang/build-images/

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/main .

EXPOSE 8080

# USER 1001
USER nonroot:nonroot

ENTRYPOINT ["/app/main"]