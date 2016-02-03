#!/usr/bin/env bash

set -e

: ${project_namespace:?must be set}

export $(workspace-generator)
fullpath=$GOPATH/src/$project_namespace
mkdir -p $fullpath
cp -r src/* $fullpath

semver=`cat version/number`
timestamp=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
output_dir=${PWD}/out

pushd $fullpath > /dev/null
  git_rev=`git rev-parse --short HEAD`
  version="${semver}-${git_rev}-${timestamp}"

  echo -e "\n building artifact..."
  go build -ldflags "-X main.version=${version}" -o "out/certify-artifacts" $project_namespace

  echo -e "\n sha1 of artifact..."
  sha1sum out/certify-artifacts

  mv out/certify-artifacts ${output_dir}/
popd > /dev/null
