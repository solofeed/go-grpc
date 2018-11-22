
FROM znly/protoc as protobuf

WORKDIR /go/proto

COPY proto .

RUN protoc --go_out=. *.proto


FROM golang:1.11.1-stretch as builder

ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS="linux"

WORKDIR /go/src

COPY services/user-service/* .

COPY --from=protobuf /go/proto proto

RUN go build -a -installsuffix nocgo -o /app .

ENTRYPOINT ["/app"]

FROM scratch  

COPY --from=builder /app ./

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 80

ENTRYPOINT ["./app"]
