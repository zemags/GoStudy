# multi-stage builds for two containers
# first build container fetches our dependencies
# using the GOlang runtime image as the base image
FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service_consignment

# Run container, execute a compiled binary
# takes the binary from build container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/service_consignment .

CMD ["./service_consignment"]
