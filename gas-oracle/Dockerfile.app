FROM ubuntu:22.04 as builder

ENV TZ Asia/Shanghai
RUN echo "${TZ}" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && apt update \
    && apt install -y tzdata \
    && rm -rf /var/lib/apt/lists/*

COPY ./app/target/release/app ./

CMD ["./app"]
