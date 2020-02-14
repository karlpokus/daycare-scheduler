#!/bin/bash

go test
TEST_RESULT=`echo $?`
if test $TEST_RESULT -ne 0; then
  echo "tests failed. Exiting"
  exit $TEST_RESULT
fi

VERSION=`cat version`

docker build -t pokus2000/daycare-scheduler:$VERSION .
