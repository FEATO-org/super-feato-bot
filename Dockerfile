ARG GO_VERSION

FROM golang:${GO_VERSION} as build
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod vendor && go mod verify
COPY . .
# ./binの下にserverバイナリが吐かれる
RUN make clean build

# 開発用コンテナとしてairでhot reloadさせる
FROM golang:${GO_VERSION} as develop
ENV DISCORD_TOKEN=""
ENV GUILD_IDS=""

# hadolint ignore=DL3008
RUN apt-get update && apt-get install -y --no-install-recommends \
  git postgresql-client \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go install github.com/cosmtrek/air@v1.27.8 \
  && go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.11.0

WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod vendor && go mod verify
COPY . .
RUN ./tools/setup-psqldef.sh

CMD ["make", "dev"]

# 本番用コンテナ
# hadolint ignore=DL3006
FROM gcr.io/distroless/base-debian11
ENV DISCORD_TOKEN=""
ENV GUILD_IDS=""

COPY --chown=nonroot:nonroot --from=build /opt/app/bin /opt/app

USER nonroot
CMD [ "/opt/app/server" ]
