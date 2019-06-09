FROM golang

ADD . /go/src/github.com/leepuppychow/fishies
WORKDIR /go/src/github.com/leepuppychow/fishies

RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go install github.com/leepuppychow/fishies

ENTRYPOINT /go/bin/fishies

EXPOSE 8080