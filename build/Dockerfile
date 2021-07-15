FROM golang:1.16.1 as builder

WORKDIR /home
RUN git clone https://github.com/monandkey/expedition3gpp

ARG SRC=main.go
ARG DST=expedition3gpp
WORKDIR /home/expedition3gpp
RUN go mod inti github.com/monandkey/expedition3gpp \
    && go mod tidy \
    && mkdir ./bin/ \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/${DST}.exe ${SRC}.go \
    && GOOS=linux GOARCH=amd64 go build -o ./bin/${DST}-linux,x64 ${SRC}.go

FROM alpine:latest as runner

WORKDIR /expedition3gpp
COPY --from=builder /home/expedition3gpp/bin/expedition3gpp ./
COPY --from=builder /home/expedition3gpp/bin/expedition3gpp.exe ./
