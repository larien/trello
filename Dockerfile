FROM golang

ADD . /go/src/trello

# RUN go get github.com/external-package

RUN go install trello

ENTRYPOINT /go/bin/trello