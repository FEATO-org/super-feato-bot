ARG GO_VERSION
FROM golang:${GO_VERSION}

WORKDIR /opt/app/bot

# ** [Optional] Uncomment this section to install additional packages. **
RUN apt upgrade

COPY . .
