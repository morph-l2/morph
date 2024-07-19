ARG GO_VERSION=1.22

FROM ubuntu:20.04

RUN apt-get update --fix-missing && ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime

# Install basic packages
RUN apt-get install build-essential curl wget git pkg-config -y
# Install dev-packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends libclang-dev libssl-dev llvm && \
    rm -rf /var/lib/apt/lists/*

# Install Go
ARG GO_VERSION
RUN rm -rf /usr/local/go
RUN if [ "$(uname -m)" = "x86_64" ]; then \
    echo amd64 >/tmp/arch; \
    elif [ "$(uname -m)" = "aarch64" ]; then \
    echo arm64 >/tmp/arch; \
    else \
    echo "Unsupported architecture"; exit 1; \
    fi
RUN wget https://go.dev/dl/go${GO_VERSION}.1.linux-$(cat /tmp/arch).tar.gz
RUN tar -C /usr/local -xzf go${GO_VERSION}.1.linux-$(cat /tmp/arch).tar.gz
RUN rm go${GO_VERSION}.1.linux-$(cat /tmp/arch).tar.gz && rm /tmp/arch
ENV PATH="/usr/local/go/bin:${PATH}"