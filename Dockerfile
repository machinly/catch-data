# First stage: complete build environment
FROM golang:1.15.2 AS builder

ADD ./ /src/

WORKDIR /src

RUN go build -v ./

# Second stage: minimal runtime environment
FROM scratch
# copy jar from the first stage
COPY --from=builder src/main main
# run jar
CMD ["main"]
