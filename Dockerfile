FROM golang

ENV PACKAGE=github.com/owenoclee/gext

# gext-server run time configuration
ENV GEXT_ADDRESS=0.0.0.0 \
    GEXT_PORT=80 \
    GEXT_DATASTORE=mysql \
    GEXT_DATASTORE_MYSQL_DSN=root:password1@tcp(db:3306)/gext \
    GEXT_VIEWS_PATH=/go/src/github.com/owenoclee/gext/views/ \
    GEXT_PUBLIC_PATH=/go/src/github.com/owenoclee/gext/public/

ADD . /go/src/${PACKAGE}

RUN go get ${PACKAGE} && \
    go install ${PACKAGE}

ENTRYPOINT /go/bin/gext

EXPOSE 8080
