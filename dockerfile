FROM golang:alpine

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/word-ladder
COPY ./src .
COPY ./data /workspaces/wordchain/data

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/word-ladder

RUN cp /go/bin/word-ladder ./word-ladder

ENV start="sport"
ENV end="crate"

ENTRYPOINT [ "sh", "-c", "word-ladder -start=$start -end=$end" ]