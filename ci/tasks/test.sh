#!/usr/bin/env bash

set -e

: ${project_namespace:?must be set}

eval $(workspace-generator)
fullpath=$GOPATH/src/$project_namespace
mkdir -p $fullpath
cp -r src/* $fullpath

pushd $fullpath > /dev/null
  echo -e "\n Vetting packages for potential issues..."
  go vet ${project_namespace}/...

  echo -e "\n Checking with golint..."
  golint ${project_namespace}/...

  echo -e "\n Testing packages..."
  ginkgo -r -race .
popd > /dev/null
