#!/usr/bin/env sh

case $1 in
app)
  chmod +x /app/app
  /app/app --config-file /app/app.yaml
  ;;
*)
  echo "./scripts/run.sh [app]"
  ;;
esac
