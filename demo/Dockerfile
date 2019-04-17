FROM golang:1.12-alpine
ENV GO111MODULE on

ENV VERSION v2.0.0
EXPOSE 443

WORKDIR /go/src/app
COPY . .

RUN go build -mod vendor -o echo-passthrough-demo .

CMD ["./echo-passthrough-demo"]