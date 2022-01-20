ARG NODE_VERSION

# 開発環境のみしか発生し得ないコンテナ
FROM node:${NODE_VERSION}
RUN apt-get update && apt-get install -y --no-install-recommends \
  git \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
WORKDIR /opt/app/bot

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
# rebuildするたびにclaspへの再ログインするのはめんどくさいため配置しておく（root user）
COPY ./.clasprc.json /root/.clasprc.json

CMD [ "npm", "run", "dev" ]
