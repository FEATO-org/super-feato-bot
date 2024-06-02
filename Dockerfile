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
  git default-mysql-client \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go install github.com/cosmtrek/air@v1.51.0 \
  && go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0

WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod vendor && go mod verify
COPY . .
RUN ./tools/setup-mysqldef.sh

CMD ["make", "dev"]

# 本番用コンテナ
# hadolint ignore=DL3006
FROM gcr.io/distroless/static-debian11
ENV DISCORD_TOKEN=""
ENV GUILD_IDS=""

COPY --chown=nonroot:nonroot --from=build /opt/app/bin /opt/app
COPY --chown=nonroot:nonroot ./credentials.json /opt/app

USER nonroot
CMD [ "/opt/app/server" ]
