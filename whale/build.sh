#!/usr/bin/env bash

govendor sync
GOOS=linux go build
zip -r whale.zip whale
