FROM golang:1.10-alpine3.7
RUN apk add --no-cache git
RUN apk add --no-cache bash

COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -d -v && go install -v
RUN go get github.com/stretchr/testify/assert
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -buildmode=exe -v

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN echo '#!/bin/bash' >> /usr/local/bin/runtests && \
    echo 'go test' >> /usr/local/bin/runtests && \
    chmod a+x /usr/local/bin/runtests

EXPOSE 80
CMD ["./app"]

# Build-time metadata as defined at http://label-schema.org
ARG BUILD_DATE=unknown
ARG VCS_NAME=unknown
ARG VCS_REF=unknown
ARG VCS_URL=unknown
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name=$VCS_NAME \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url=$VCS_URL \
      org.label-schema.schema-version="1.0"
