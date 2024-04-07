# Use base golang image from Docker Hub
FROM golang:1.22-alpine AS build
RUN apk add --update --no-cache git
WORKDIR /src/secrets-manager
COPY . ./
RUN go mod download
RUN go build -o /app -v ./cmd/Secrets-manager

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build /app /.
ENTRYPOINT ["/app"]