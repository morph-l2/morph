# Build Geth in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

COPY . /staking-oracle

WORKDIR /staking-oracle

RUN make staking-oracle

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /staking-oracle/staking-oracle /usr/local/bin/

CMD ["staking-oracle"]