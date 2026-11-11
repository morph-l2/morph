# Build from the repository root so the local bindings module is available:
#   docker build -f challenger/Dockerfile -t morph/challenger .
FROM golang:1.24-bookworm AS builder

RUN apt-get update && apt-get install -y --no-install-recommends gcc libc6-dev && rm -rf /var/lib/apt/lists/*

WORKDIR /src
COPY bindings ./bindings
COPY challenger ./challenger

WORKDIR /src/challenger
RUN go mod download
RUN CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go build -o /out/challenger ./cmd

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /out/challenger /usr/local/bin/challenger

EXPOSE 8080 6060

ENTRYPOINT ["challenger"]
