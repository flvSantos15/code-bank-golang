FROM golang:1.25-bookworm

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    protobuf-compiler \
    librdkafka-dev \
    wget \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install github.com/spf13/cobra-cli@latest

RUN wget https://github.com/ktr0731/evans/releases/download/0.9.1/evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    mv evans /go/bin/evans && \
    rm -f evans_linux_amd64.tar.gz

CMD ["tail", "-f", "/dev/null"]
