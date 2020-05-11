#!/bin/bash

PROJECT_NAME=mosaic

MACHINE=`uname -m`
KERNEL=`uname -s`
TARGET_BASE=centos
TARGET_VERSION=latest

OPTS=`getopt -o h --long help -n 'setup.sh' -- "$@"`
if [ $? != 0 ] ; then echo "failed parsing options" >&2 ; exit 1 ; fi
eval set -- "$OPTS"

HELP=false

while true; do
  case "$1" in
    -h | --help ) HELP=true; shift ;;
    -- ) shift; break ;;
    * ) break ;;
  esac
done

echo ">>> docker build for ${TARGET_BASE}:${TARGET_VERSION} on ${MACHINE}:${KERNEL}"

docker pull ${TARGET_BASE}:${TARGET_VERSION}
docker build -t ${PROJECT_NAME}:latest --file docker/Dockerfile .
docker image prune -f
docker container prune -f
