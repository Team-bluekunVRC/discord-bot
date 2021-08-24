# syntax=docker/dockerfile:1.0-experimental
FROM golang:1.17-alpine as builder
ARG VERSION
ENV GOCACHE "/go-build-cache"
WORKDIR /src

# Why: Pinning versions in a docker container isn't
# really worth the effort.
# hadolint ignore=DL3018
RUN apk add --no-cache make bash git libc-dev gcc

COPY . .

# --mount here allows us to cache the packages even if
# it's invalidated by go.mod,go.sum
RUN --mount=type=cache,target=/go/pkg make dep

# --mount here allows us to save go build cache across builds
# but also needed to use the package cache above
RUN --mount=type=cache,target=/go-build-cache --mount=type=cache,target=/go/pkg make build APP_VERSION=${VERSION}

FROM alpine:3.14
ENTRYPOINT ["/usr/bin/discord-bot"]

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates

COPY --from=builder /src/bin/discord-bot /usr/bin/discord-bot