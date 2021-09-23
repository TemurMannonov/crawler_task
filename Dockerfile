# workspace (GOPATH) configured at /go
FROM golang:1.17 as builder


RUN mkdir -p $GOPATH/src/crawler_task
WORKDIR $GOPATH/src/crawler_task

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/crawler_task /


FROM alpine

EXPOSE 8000

COPY --from=builder crawler_task .

ENTRYPOINT ["/crawler_task"]
