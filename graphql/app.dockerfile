FROM golang:1.23-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/ranjannkumar/go-grpc-graphql-microservice
COPY go.mod go.sum ./
COPY account account
COPY catalog catalog
COPY order order
COPY graphql graphql
RUN GO111MODULE=on go build -v -o /go/bin/app ./graphql

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]