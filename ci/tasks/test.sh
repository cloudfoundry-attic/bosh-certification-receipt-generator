#!/usr/bin/env bash

set -e

: ${project_namespace:?must be set}

eval $(workspace-generator)
fullpath=$GOPATH/src/$project_namespace
mkdir -p $fullpath
cp -r src/* $fullpath

pushd $fullpath > /dev/null
  echo -e "\n Vetting & Linting packages for potential issues..."

  for i in $(go list ./... | grep -v vendor); do
    go vet $i
    golint $i | grep -v "should have comment or be unexported"
  done

  echo -e "\n Testing packages..."
  ginkgo -r -race .
popd > /dev/null
