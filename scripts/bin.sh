#!/usr/bin/env bash

export GOPRIVATE="gitlab.com/gooride"

SCRIPTPATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"

CURRENT_DIR=$SCRIPTPATH
ROOT_DIR="$(dirname $CURRENT_DIR)"

INFRA_LOCAL_COMPOSE_FILE=$ROOT_DIR/build/docker-compose.yaml

function local_infra() {
  docker-compose -f $INFRA_LOCAL_COMPOSE_FILE $@
}

function infra() {
  case $1 in
  up)
    local_infra up ${@:2}
    ;;
  down)
    local_infra down ${@:2}
    ;;
  build)
    local_infra build ${@:2}
    ;;
  *)
    echo "up|down|build [docker-compose command arguments]"
    ;;
  esac
}

function init() {
    cd $CURRENT_DIR/..
    goimports -w ./..
    go fmt ./...
}

function run_test() {
  # run all unit tests
  echo 'run unit testing'
  go test ./... -short || {
    echo 'unit testing failed'
    exit 1
  }
}

function api_start() {
  echo "Start generating Swagger docs ..."
  swag init -d $ROOT_DIR -g server/init.go -o $ROOT_DIR/docs
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/service/main.go"
  go run $ENTRY_FILE --config-file=$CONFIG_FILE
}

function worker_start() {
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/worker/main.go"
  go run $ENTRY_FILE --config-file=$CONFIG_FILE
}

# generate code from proto
function api_proto_gen() {
    echo ""
    echo "Compiling protobuf file"
    cd gRPC/protos
    # generate go code
    protoc \
        -I. \
        -I/usr/local/include \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:. \
        $1/*.proto
    cd -

    cd gRPC/protos/$1
    mockgen -source=$1.pb.go -destination=$1_mock.go -package=$1
    cd -
}

function api_test() {
    infra up -d
    go clean -testcache ./...
    go generate ./...

    # Call migrate tool first
    api_migrate || {
        echo 'migrate database failed'
        exit 1
    }
    setup_env_variables
    run_test
}

# Add more command 'migrate' for migrate tool
function api_migrate() {
    echo "Starting migration..."
    infra up -d
    setup_env_variables
    ENTRY_FILE="$ROOT_DIR/cmd/migrate/main.go"
    go run $ENTRY_FILE
}

function api_docs_gen() {
   echo "Start generating Swagger docs ..."
   swag init -d $ROOT_DIR/handler -g configure.go -o $ROOT_DIR/docs
}

# Setup variables environment for app
function setup_env_variables() {
    set -a
    export $(grep -v '^#' "$ROOT_DIR/build/.base.env" | xargs -0) >/dev/null 2>&1
    . $ROOT_DIR/build/.base.env
    set +a
    export CONFIG_FILE=$ROOT_DIR/build/app.yaml
}


function api() {
    case $1 in
    test)
        api_test
        ;;
    migrate)
        api_migrate
        ;;
    start)
        api_start
        ;;
    worker_start)
        worker_start
        ;;
    proto_gen)
        api_proto_gen ${@:2}
        ;;
    docs_gen)
        api_docs_gen 
        ;;
    *)
        echo "[test|start|dev_start|proto_gen|docs_gen|migrate]"
        ;;
    esac
}

function lint() {
    export GO111MODULE=on
    command -v golangci-lint >/dev/null 2>&1 || {
      curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.26.0
    }

    golangci-lint run --timeout 5m
}

function code_format() {
    gofmt -w internal/ pkg/ cmd/

    goimports -w internal/ pkg/ cmd/
}

git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

case $1 in
init)
    init
    ;;
infra)
    infra ${@:2}
    ;;
api)
    api ${@:2}
    ;;
lint)
    lint
    ;;
*)
    echo "./scripts/bin.sh [infra|api|lint|add_version]"
    ;;
esac

