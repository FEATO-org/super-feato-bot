#!/bin/bash
set -euxo pipefail
WORKDIR=$(pwd)
cd "$(dirname "$0")" && cd ../

set +e
psql -U ${PGUSER} -H ${PGHOST} -f ./db/init/setup.sql
set -e

make migrate
psql -U ${PGUSER} -H ${PGHOST} -d ${APP_ENV} -f ./db/seed.sql
