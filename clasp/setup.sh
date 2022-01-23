#!/bin/bash
set -euxo pipefail
WORKDIR=$(pwd)
cd "$(dirname "$0")"

npx clasp login
mv /root/.clasprc.json ./.clasprc.json
