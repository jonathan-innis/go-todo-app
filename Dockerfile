FROM golang:latest

RUN mkdir /build
WORKDIR /build

ADD go.mod /build
ADD go.sum /build
ADD pkg /build/pkg
ADD main.go /build

RUN go mod download

RUN go build -o /app/main /build/main.go

CMD /app/main