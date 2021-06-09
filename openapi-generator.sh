#! /usr/bin/env bash

set -e

OPENAPI_GENERATOR_VERSION="5.1.1"
OPENAPI_GENERATOR_URL="https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${OPENAPI_GENERATOR_VERSION}/openapi-generator-cli-${OPENAPI_GENERATOR_VERSION}.jar"
OPENAPI_GENERATOR_PATH="./openapi-generator-cli.jar"

if [[ ! -f ${OPENAPI_GENERATOR_PATH} ]]; then
    wget ${OPENAPI_GENERATOR_URL} -O ${OPENAPI_GENERATOR_PATH}
fi

java -ea \
  ${JAVA_OPTS} \
  -Xms512M \
  -Xmx1024M \
  -server \
  -jar ${OPENAPI_GENERATOR_PATH} \
  "$@"