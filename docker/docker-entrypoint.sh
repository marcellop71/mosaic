#!/bin/bash
set -e

redis-server &

exec "$@"
