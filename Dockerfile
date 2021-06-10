FROM golang:latest

WORKDIR $GOPATH/src/github.com/EDDYCJY/QAPlatform
COPY . $GOPATH/src/github.com/EDDYCJY/QAPlatform
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./QAPlatform"]