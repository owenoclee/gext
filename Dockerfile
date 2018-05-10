FROM golang

ENV PACKAGE=github.com/owenoclee/gext-server

# gext-server run time configuration
ENV ADDRESS=localhost \
    PORT=8080 \
    DATASTORE=mysql \
    DATASTORE_MYSQL_DSN=root@/gext

ADD . /go/src/${PACKAGE}

RUN apt-get update && apt-get install -y bsdtar && \
    mkdir -p /protoc && \
    wget -qO- https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip | \
        bsdtar -xvf- -C /protoc && \
    chmod +x /protoc/bin/protoc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

ENV PATH=/protoc/bin/:$PATH

RUN protoc --go_out=/go/src/${PACKAGE}/models \
        --proto_path=/go/src/${PACKAGE}/models/protos \
        /go/src/${PACKAGE}/models/protos/*.proto && \
    go get ${PACKAGE} && \
    go install ${PACKAGE}

ENTRYPOINT /go/bin/gext-server

EXPOSE 8080

