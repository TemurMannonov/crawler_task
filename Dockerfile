# workspace (GOPATH) configured at /go
FROM golang:1.16 as builder


RUN mkdir -p $GOPATH/src/github.com/TemurMannonov/crawler_task
WORKDIR $GOPATH/src/github.com/TemurMannonov/crawler_task

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/github.com/TemurMannonov/crawler_task /

FROM alpine
COPY --from=builder github.com/TemurMannonov/crawler_task .
ENTRYPOINT ["/github.com/TemurMannonov/crawler_task"]
