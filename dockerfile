FROM golang:alpine as build

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/word-ladder
COPY ./src .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/word-ladder

FROM alpine:latest
COPY --from=build /go/bin/word-ladder /bin/word-ladder
COPY ./data /workspaces/wordchain/data

ENV start="sport"
ENV end="crate"

ENTRYPOINT [ "/bin/word-ladder", "--", "-start=$start -end=$end" ]