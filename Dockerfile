FROM golang:1.9

ADD . /go/src/go-api-scaffolding

WORKDIR /go/src/go-api-scaffolding

RUN go get github.com/golang/dep \ 
    && go install github.com/golang/dep/cmd/dep \
    && dep ensure -v
    
RUN go install go-api-scaffolding

RUN ls

RUN cp /go/bin/go-api-scaffolding /go/src/go-api-scaffolding

EXPOSE 9000

ENTRYPOINT ["./go-api-scaffolding"]