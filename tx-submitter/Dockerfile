# Build Geth in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

COPY . /tx-submitter

WORKDIR /tx-submitter

RUN make build

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /tx-submitter/tx-submitter /usr/local/bin/

CMD ["tx-submitter"]