FROM golang:1.16.1 as builder

WORKDIR /home
RUN git clone https://github.com/monandkey/expedition3gpp

WORKDIR /home/expedition3gpp
RUN go mod tidy \
    && mkdir ./bin/ \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/expedition3gpp.exe main.go \
    && GOOS=linux GOARCH=amd64 go build -o ./bin/expedition3gpp main.go
