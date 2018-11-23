FROM golang:1.11.1-stretch as builder

ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS="linux"

WORKDIR /go/src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

COPY data.csv .

RUN go build -a -installsuffix nocgo -o /app .

ENTRYPOINT ["/app"]

FROM scratch  

COPY --from=builder /app ./

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./app"]
