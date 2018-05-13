FROM golang

ENV PACKAGE=github.com/owenoclee/gext

# gext-server run time configuration
ENV GEXT_ADDRESS=localhost \
    GEXT_PORT=8080 \
    GEXT_DATASTORE=mysql \
    GEXT_DATASTORE_MYSQL_DSN=root@/gext \
    GEXT_VIEWS_PATH=/go/src/github.com/owenoclee/gext/views/ \
    GEXT_PUBLIC_PATH=/go/src/github.com/owenoclee/gext/public/

ADD . /go/src/${PACKAGE}

RUN go get ${PACKAGE} && \
    go install ${PACKAGE}

ENTRYPOINT /go/bin/gext

EXPOSE 8080
