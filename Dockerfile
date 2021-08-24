FROM golang:1.17-alpine3.14 as builder
ARG VERSION
ENV GOCACHE "/go-build-cache"
WORKDIR /src

# Why: Pinning versions in a docker container isn't
# really worth the effort.
# hadolint ignore=DL3018
RUN apk add --no-cache make bash git libc-dev gcc py3-pip jq openssh \
  &&  pip3 install yq \
  &&  git config --global url."https://github.com/".insteadOf git@github.com: \
  &&  git config --global url."https://".insteadOf git:// \
  &&  git config --global url."https://".insteadOf ssh://

COPY . .

# --mount here allows us to cache the packages even if
# it's invalidated by go.mod,go.sum
RUN make dep

# --mount here allows us to save go build cache across builds
# but also needed to use the package cache above
RUN make build APP_VERSION=${VERSION}

FROM alpine:3.14
ENTRYPOINT ["/usr/bin/discord-bot"]

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates

COPY --from=builder /src/bin/discord-bot /usr/bin/discord-bot
