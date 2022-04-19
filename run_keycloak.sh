#!/bin/bash

docker run \
  -e KEYCLOAK_USER=admin \
  -e KEYCLOAK_PASSWORD=admin \
  -p 8080 \
  -it quay.io/keycloak/keycloak:17.0.1 \
  start-dev
