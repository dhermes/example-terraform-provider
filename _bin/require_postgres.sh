#!/bin/sh
# Copyright 2021 Danny Hermes
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

. "$(dirname "${0}")/exists.sh"
. "$(dirname "${0}")/require_env_var.sh"

requireEnvVar "DB_HOST"
requireEnvVar "DB_PORT"
requireEnvVar "DB_ADMIN_DSN"

##########################################################
## Don't exit until `pg_isready` returns 0 in container ##
##########################################################

# NOTE: This is used strictly for the status code to determine readiness.
pgIsReady() {
  pg_isready --dbname "${DB_ADMIN_DSN}" > /dev/null 2>&1
}

exists "pg_isready"
# Cap at 20 retries.
i=0; while [ ${i} -le 20 ]
do
  if pgIsReady
  then
    exit 0
  fi
  i=$((i+1))
  echo "Checking if PostgresSQL is ready on ${DB_HOST}:${DB_PORT} (attempt $i)"
  sleep "0.1"
done

exit 1
