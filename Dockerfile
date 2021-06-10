FROM golang:latest

WORKDIR $GOPATH/src/github.com/EDDYCJY/QAPlatform
COPY . $GOPATH/src/github.com/EDDYCJY/QAPlatform
RUN go build .

EXPOSE 3000
ENTRYPOINT ["./QAPlatform"]