FROM debian:12 as builder
ARG TARGETARCH
RUN apt-get update && apt-get install -y curl clang gcc llvm make libbpf-dev -y
RUN curl -LO https://go.dev/dl/go1.20.linux-${TARGETARCH}.tar.gz && tar -C /usr/local -xzf go*.linux-${TARGETARCH}.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /app
COPY . .
RUN make build

# FROM gcr.io/distroless/base-debian12@sha256:d64f5483d2fd0cec2260941c443cb2947102e46e1a9fe36a321d0a788c1a49e0
# FROM gcr.io/distroless/base-debian12:debug
FROM alpine:3.18
COPY --from=builder /app/otel-go-instrumentation /
RUN apk add --no-cache bash

CMD ["/otel-go-instrumentation"]
