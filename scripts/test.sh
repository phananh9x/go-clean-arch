#!/usr/bin/env bash

echo "--- Script Test Application With Avariable. ---"

function mockCustomer() {
    echo "mock genarate mock customer application"
}

function mock() {
    echo "mockgen coding"
}

function unitTest() {
    echo "exec unit-test"
}



case $1 in
mock)
    echo "... [mock]"
    mock
    ;;
test)
    echo "... [test]"
    unitTest
    ;;
*)
    echo "... [mock|test]"
esac 