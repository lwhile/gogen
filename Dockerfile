FROM golang

RUN go get github.com/lwhile/gogen

WORKDIR $GOPATH/src/github.com/lwhile/gogen

RUN go build -o $GOPATH/src/github.com/lwhile/gogen/gogen cmd/cmd.go

EXPOSE 4928

CMD ["./gogen", "-web"]