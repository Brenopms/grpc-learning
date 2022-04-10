FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app/grpc-learning
COPY . .

RUN go get -d -v ./cmd
RUN go build -o /go/bin/product cmd/main.go
COPY ./pkg/config/envs /go/bin/

FROM alpine as runtime

WORKDIR /app
# Copy our static executable.
COPY --from=builder /go/bin /app


RUN ls
RUN pwd
# Run the hello binary.
EXPOSE 50052
ENTRYPOINT ["/app/product"]