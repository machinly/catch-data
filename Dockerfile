FROM golang:1.15.2 AS builder

ADD ./ /src/

WORKDIR /src

RUN go build -o catch-data ./ 

FROM ubuntu:xenial
COPY --from=builder /src/catch-data /
CMD ["/catch-data"]
