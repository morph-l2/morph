# Build Geth in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

COPY .. /node
RUN cd /node && echo $(ls -1) && make build


FROM alpine:latest as initializer
RUN apk add --no-cache ca-certificates

COPY --from=builder /node/build/bin/tendermint /usr/local/bin/
RUN echo "Initializing tendermint."
RUN tendermint init --home /data
RUN sed -i  's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' /data/config/config.toml


FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /node/build/bin/morphnode /usr/local/bin/
COPY --from=initializer /data /data
