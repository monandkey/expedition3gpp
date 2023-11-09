FROM node:18.6.0
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8

RUN apt update && apt install -y \
    build-essential \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    curl

ARG GO_VERSION="1.21.4"
ENV GOPATH $HOME/go
ENV GOBIN $GOPATH/bin
ENV PATH $PATH:/usr/local/go/bin:$GOBIN

RUN <<EOF
    curl -OL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
    rm -rf /usr/local/go
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    rm -rf go${GO_VERSION}.linux-amd64.tar.gz

    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    go get github.com/steebchen/prisma-client-go
EOF
