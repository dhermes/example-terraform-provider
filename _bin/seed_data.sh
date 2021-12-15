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

requireEnvVar "BOOKS_ADDR"
exists "curl"
exists "jq"

## Add Author
echo ':: Adding authors:'
RESPONSE_AUTHOR1=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "Anne", "last_name": "Rice"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID1=$(echo "${RESPONSE_AUTHOR1}" | jq '.author_id' -r)
echo "Added author Anne Rice: ${AUTHOR_ID1}"

RESPONSE_AUTHOR2=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "John", "last_name": "Steinbeck"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID2=$(echo "${RESPONSE_AUTHOR2}" | jq '.author_id' -r)
echo "Added author John Steinbeck: ${AUTHOR_ID2}"

RESPONSE_AUTHOR3=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "JK", "last_name": "Rowling"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID3=$(echo "${RESPONSE_AUTHOR3}" | jq '.author_id' -r)
echo "Added author JK Rowling: ${AUTHOR_ID3}"

RESPONSE_AUTHOR4=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "Ernest", "last_name": "Hemingway"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID4=$(echo "${RESPONSE_AUTHOR4}" | jq '.author_id' -r)
echo "Added author Ernest Hemingway: ${AUTHOR_ID4}"

RESPONSE_AUTHOR5=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "Kurt", "last_name": "Vonnegut"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID5=$(echo "${RESPONSE_AUTHOR5}" | jq '.author_id' -r)
echo "Added author Kurt Vonnegut: ${AUTHOR_ID5}"

RESPONSE_AUTHOR6=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "Agatha", "last_name": "Christie"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID6=$(echo "${RESPONSE_AUTHOR6}" | jq '.author_id' -r)
echo "Added author Agatha Christie: ${AUTHOR_ID6}"

RESPONSE_AUTHOR7=$(curl \
  --silent --show-error --fail \
  --data-binary '{"first_name": "James", "last_name": "Joyce"}' \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/author")
AUTHOR_ID7=$(echo "${RESPONSE_AUTHOR7}" | jq '.author_id' -r)
echo "Added author James Joyce: ${AUTHOR_ID7}"

## Get Author By ID
echo '--------------------------------------------------'
echo ':: Getting author James Joyce by ID:'
GET_RESULT=$(curl \
  --silent --show-error --fail \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/authors/${AUTHOR_ID7}")
echo "${GET_RESULT}" | jq

## Get All Authors
echo '--------------------------------------------------'
echo ':: Getting all authors:'
GET_RESULT=$(curl \
  --silent --show-error --fail \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/authors")
echo "${GET_RESULT}" | jq

## Add Books
echo '--------------------------------------------------'
echo ':: Adding books for recently added authors:'
RESPONSE_BOOK1=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID1}\", \"title\": \"The Wolf Gift\", \"publish_date\": \"2012-02-14T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID1=$(echo "${RESPONSE_BOOK1}" | jq '.book_id' -r)
echo "Added book The Wolf Gift by Anne Rice: ${BOOK_ID1}"

RESPONSE_BOOK2=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID1}\", \"title\": \"Interview with the Vampire\", \"publish_date\": \"1976-05-05T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID2=$(echo "${RESPONSE_BOOK2}" | jq '.book_id' -r)
echo "Added book Interview with the Vampire by Anne Rice: ${BOOK_ID2}"

RESPONSE_BOOK3=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID1}\", \"title\": \"The Queen of the Damned\", \"publish_date\": \"1988-09-12T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID3=$(echo "${RESPONSE_BOOK3}" | jq '.book_id' -r)
echo "Added book The Queen of the Damned by Anne Rice: ${BOOK_ID3}"

RESPONSE_BOOK4=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID2}\", \"title\": \"East of Eden\", \"publish_date\": \"1952-09-19T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID4=$(echo "${RESPONSE_BOOK4}" | jq '.book_id' -r)
echo "Added book East of Eden by John Steinbeck: ${BOOK_ID4}"

RESPONSE_BOOK5=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID3}\", \"title\": \"Harry Potter and the Goblet of Fire\", \"publish_date\": \"2000-07-08T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID5=$(echo "${RESPONSE_BOOK5}" | jq '.book_id' -r)
echo "Added book Harry Potter and the Goblet of Fire by JK Rowling: ${BOOK_ID5}"

RESPONSE_BOOK6=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID6}\", \"title\": \"Murder on the Orient Express\", \"publish_date\": \"1934-01-01T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID6=$(echo "${RESPONSE_BOOK6}" | jq '.book_id' -r)
echo "Added book Murder on the Orient Express by Agatha Christie: ${BOOK_ID6}"

RESPONSE_BOOK7=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID7}\", \"title\": \"Ulysses\", \"publish_date\": \"1922-02-02T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID7=$(echo "${RESPONSE_BOOK7}" | jq '.book_id' -r)
echo "Added book Ulysses by James Joyce: ${BOOK_ID7}"

RESPONSE_BOOK8=$(curl \
  --silent --show-error --fail \
  --data-binary "{\"author_id\": \"${AUTHOR_ID7}\", \"title\": \"Finnegans Wake\", \"publish_date\": \"1939-05-04T00:00:00Z\"}" \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/book")
BOOK_ID8=$(echo "${RESPONSE_BOOK8}" | jq '.book_id' -r)
echo "Added book Finnegans Wake by James Joyce: ${BOOK_ID8}"

## Get All Books by Author
echo '--------------------------------------------------'
echo ':: Getting all books by James Joyce:'
GET_RESULT=$(curl \
  --silent --show-error --fail \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/books/${AUTHOR_ID7}")
echo "${GET_RESULT}" | jq

## Get Author By ID (again)
echo '--------------------------------------------------'
echo ':: Getting author James Joyce by ID (again):'
GET_RESULT=$(curl \
  --silent --show-error --fail \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/authors/${AUTHOR_ID7}")
echo "${GET_RESULT}" | jq

## Get All Authors (again)
echo '--------------------------------------------------'
echo ':: Getting all authors (again):'
GET_RESULT=$(curl \
  --silent --show-error --fail \
  --header 'Content-Type: application/json' \
  "${BOOKS_ADDR}/v1alpha1/authors")
echo "${GET_RESULT}" | jq
