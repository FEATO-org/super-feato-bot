#!/bin/bash
set -euxo pipefail
WORKDIR=$(pwd)
cd "$(dirname "$0")" && cd ../

make migrate
mysql -u ${DBUSER} -h ${DBHOST} -p${DBPASSWORD} sfs < ./db/seed.sql
