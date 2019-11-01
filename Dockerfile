FROM golang

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/flog

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /bin/flog \
    && chmod +x docker-entrypoint.sh

FROM alpine

ENV SLEEP_TIME=1 \
    LOG_FILE_PATH=/dev/stdout

COPY --from=0 /bin/flog /bin/flog
COPY --from=0 /go/src/flog/docker-entrypoint.sh /bin/docker-entrypoint.sh
ENTRYPOINT ["/bin/docker-entrypoint.sh"]
