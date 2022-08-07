FROM golang:latest

RUN mkdir /build
WORKDIR /build

ADD go.mod /build
ADD go.sum /build
ADD main.go /build

RUN go build -o /app/main /build/main.go

CMD /app/main