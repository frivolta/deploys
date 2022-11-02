#!/bin/sh
set -e
source /app/app.env

echo "start the app------------------------------------------------"
echo "run db migration on: $ENV_VARIABLE_RANDOM"
exec "$@"