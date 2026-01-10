FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN ["go", "mod", "download"]
COPY . .
RUN ["go", "build", "-o", "/app/main", "./cmd/main"]
FROM alpine
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=build /app/main .
CMD ["./main"]
