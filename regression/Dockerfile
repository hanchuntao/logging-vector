#
# LADING
#
FROM ghcr.io/datadog/lading@sha256:557e5ebc5ebbb5702504ae01c9f149c950c31e542ccca2debaf26b8dee78f5c3 as lading

#
# VECTOR BUILDER
#
FROM ghcr.io/vectordotdev/vector/soak-builder@sha256:b598dc1eeb8661e4928238093a979597fd98f80554e5276d8c17cf4781ea3b82 as builder
WORKDIR /vector
COPY . .
RUN bash scripts/environment/install-protoc.sh
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/usr/local/cargo/git \
    --mount=type=cache,target=/vector/target \
    /usr/local/bin/mold -run cargo build --bin vector --release && \
    cp target/release/vector .

#
# TARGET
#
FROM docker.io/debian:bookworm-slim@sha256:39868a6f452462b70cf720a8daff250c63e7342970e749059c105bf7c1e8eeaf
RUN apt-get update && apt-get dist-upgrade -y && apt-get -y --no-install-recommends install zlib1g ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=lading /usr/bin/lading /usr/local/bin/lading
COPY --from=builder /vector/vector /usr/local/bin/vector
RUN mkdir --parents --mode=0777 /var/lib/vector

# Smoke test
RUN ["/usr/local/bin/lading", "--help"]
RUN ["/usr/local/bin/vector", "--version"]

ENTRYPOINT ["/usr/local/bin/lading"]
