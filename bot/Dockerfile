ARG GO_VERSION

# 共通のbuilder
FROM golang:${GO_VERSION} as build
WORKDIR /opt/app/bot
COPY go.mod go.sum ./
RUN go mod tidy && \
  go mod vendor
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
  && rm -rf /var/lib/apt/lists/*

WORKDIR /opt/app/bot
RUN useradd -m golang
COPY --chown=golang:golang --from=build /opt/app/bot ./
RUN ./tools/setup-psqldef.sh

USER golang
RUN go install github.com/cosmtrek/air@v1.27.8 && \
  go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.11.0

CMD ["make", "dev"]

# 本番用コンテナ（Herokuを想定）
# hadolint ignore=DL3006
FROM gcr.io/distroless/static-debian11 as production
ENV DISCORD_TOKEN=""
ENV GUILD_IDS=""

COPY --chown=nonroot:nonroot --from=build /opt/app/bot/bin /opt/app
USER nonroot
CMD [ "/opt/app/server" ]
