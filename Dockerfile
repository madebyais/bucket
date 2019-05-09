FROM golang:1.12 as builder
LABEL maintainer="Faris <madebyais@gmail.com>"

ARG APP_NAME=bucket
ARG CONF_DIR=/etc/bucket
ARG LOCAL_UPLOAD_DIR=/data/bucket

RUN mkdir -p ${CONF_DIR}
RUN mkdir -p ${LOCAL_UPLOAD_DIR}

WORKDIR $GOPATH/src/github.com/madebyais/bucket
COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/bucket .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

ARG CONF_DIR=/etc/bucket
ARG LOCAL_UPLOAD_DIR=/data/bucket

RUN mkdir -p ${CONF_DIR}
RUN mkdir -p ${LOCAL_UPLOAD_DIR}

WORKDIR /root/

COPY --from=builder /go/bin/bucket .

EXPOSE 8700

VOLUME [ "/etc/bucket", "/opt/bucket" ]

CMD [ "./bucket", "server", "-c", "/etc/bucket/bucket.yaml" ]