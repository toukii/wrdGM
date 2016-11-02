FROM golang:latest

ENV GOPATH /app/gopath

RUN go get github.com/toukii/wrdGM && go get github.com/go-macaron/macaron

WORKDIR /wkdir
RUN git clone --depth 1 git://github.com/toukii/wrdGM.git .
WORKDIR /wkdir
RUN go build -o wrdGM main.go

EXPOSE 4000:80

CMD ["/wkdir/wrdGM"]
