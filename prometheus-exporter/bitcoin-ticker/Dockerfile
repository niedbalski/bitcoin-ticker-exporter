FROM golang:alpine as build
MAINTAINER  Jorge Niedbalski <jnr@metaklass.org>

RUN apk --no-cache add git

WORKDIR /go/src/github.com/niedbalski/bitcoin-ticker-exporter
COPY . .
RUN go get github.com/tools/godep && godep restore && go build -o bitcoin-ticker-exporter

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/src/github.com/niedbalski/bitcoin-ticker-exporter/bitcoin-ticker-exporter .
COPY ./docker/config/exporter.yml /exporter.yml

ENTRYPOINT ["./bitcoin-ticker-exporter"] 
