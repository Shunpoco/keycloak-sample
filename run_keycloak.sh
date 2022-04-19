#!/bin/bash

docker run \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  -p 8080:8080 \
  -it quay.io/keycloak/keycloak:17.0.1 \
  start-dev
