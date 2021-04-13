FROM golang:1.15-buster

WORKDIR /go/src/mag

COPY . .

RUN go mod download
# RUN go get -d -v ./...
# RUN go install -v ./...
