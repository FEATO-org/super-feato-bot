ARG NODE_VERSION

# 開発環境のみしか発生し得ないコンテナ
FROM node:${NODE_VERSION}
RUN apt update && apt install -y \
  git
WORKDIR /opt/app/bot

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
# rebuildするたびにclaspへの再ログインするのはめんどくさいため配置しておく（root user）
COPY ./.clasprc.json /root/.clasprc.json

CMD [ "npm", "run", "dev" ]
