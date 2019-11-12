FROM golang:1.12-alpine
ENV GO111MODULE on

ENV VERSION v1.0.0
EXPOSE 8080
EXPOSE 8443

WORKDIR /go/src/app
COPY . .

RUN go build -mod vendor -o echo-passthrough-demo .

CMD ["./echo-passthrough-demo"]