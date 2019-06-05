FROM golang:alpine

COPY ./src /go/src/word-ladder
COPY ./data /workspaces/wordchain/data

RUN go get -d -v /go/src/word-ladder
RUN go install -v /go/src/word-ladder

ENV start="hot"
ENV end="ice"

ENTRYPOINT [ "sh", "-c", "word-ladder -start=$start -end=$end" ]
