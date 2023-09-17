#!/bin/bash
set -euxo pipefail
WORKDIR=$(pwd)
cd "$(dirname "$0")"

ARCH=$(arch)
if [ ${ARCH} = "aarch64" ]; then
  ARCH_TYPE="arm64"
elif [ ${ARCH} = "x86_64" ]; then
  ARCH_TYPE="amd64"
else
  echo "unexpected arch type"
  exit 1
fi

curl -OL https://github.com/k0kubun/sqldef/releases/download/v0.16.7/mysqldef_linux_${ARCH_TYPE}.tar.gz
tar xf mysqldef_linux_${ARCH_TYPE}.tar.gz -C /usr/local/bin/
rm mysqldef_linux_${ARCH_TYPE}.tar.gz
