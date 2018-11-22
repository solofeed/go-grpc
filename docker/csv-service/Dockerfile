
FROM znly/protoc as protobuf

WORKDIR /go/proto

COPY proto .

RUN protoc --go_out=. *.proto


FROM golang:1.11.1-stretch as builder

ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS="linux"

WORKDIR /go/src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./

RUN go mod download

COPY services/csv-service/* .

COPY data.csv .

RUN mkdir -p /go/src/github.com/go-grpc/proto

COPY --from=protobuf /go/proto /go/src/github.com/go-grpc/proto

RUN go build -a -installsuffix nocgo -o /app .

ENTRYPOINT ["/app"]

FROM scratch  

COPY --from=builder /app ./

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./app"]
