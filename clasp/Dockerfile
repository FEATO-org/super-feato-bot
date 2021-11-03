ARG NODE_VERSION
FROM node:${NODE_VERSION}

WORKDIR /opt/app/bot

# ** [Optional] Uncomment this section to install additional packages. **
RUN apt upgrade

COPY . .
