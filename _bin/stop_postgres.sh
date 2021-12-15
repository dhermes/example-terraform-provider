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

exists "docker"
requireEnvVar "DB_CONTAINER_NAME"
CONTAINER_EXISTS=$(docker ps --quiet --filter "name=${DB_CONTAINER_NAME}")
if [ -z "${CONTAINER_EXISTS}" ]; then
  echo "Container ${DB_CONTAINER_NAME} is not currently running."
else
  docker rm --force "${DB_CONTAINER_NAME}" > /dev/null
  echo "Container ${DB_CONTAINER_NAME} stopped."
fi

requireEnvVar "DB_NETWORK_NAME"
NETWORK_EXISTS=$(docker network ls --quiet --filter "name=${DB_NETWORK_NAME}")
if [ -z "${NETWORK_EXISTS}" ]; then
  echo "Network ${DB_NETWORK_NAME} is not currently running."
  exit
fi

NETWORK_CONTAINERS=$(docker network inspect --format "{{len .Containers}}" "${DB_NETWORK_NAME}")
if [ "${NETWORK_CONTAINERS}" = "0" ]; then
  docker network rm "${DB_NETWORK_NAME}" > /dev/null
  echo "Network ${DB_NETWORK_NAME} stopped."
else
  echo "Network ${DB_NETWORK_NAME} still has ${NETWORK_CONTAINERS} container(s) running."
fi
