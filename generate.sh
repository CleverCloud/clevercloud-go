#! /usr/bin/env bash

set -eu

SCHEMA_URL="https://api.clever-cloud.com/v2/openapi.json"
OUTPUT_FOLDER="clevercloud"
PACKAGE_VERSION="0.1.0"

GOFMT_PATH=$(which gofmt)

export GO_POST_PROCESS_FILE="${GOFMT_PATH} -w"

./openapi-generator.sh generate \
  -g go \
  -i "${SCHEMA_URL}" \
  -o "${OUTPUT_FOLDER}" \
  --skip-validate-spec \
  --enable-post-process-file \
  --minimal-update \
  --git-user-id clevercloud \
  --git-repo-id clevercloud-go \
  --additional-properties=packageName="clevercloud",packageVersion="${PACKAGE_VERSION}",isGoSubmodule=true