#!/usr/bin/env bash

set -e

# $project_namespace="github.com/cloudfoundry-incubator/bosh-certification-receipt-generator"
: ${project_namespace:?must be set}

export $(cat <(workspace-generator))
fullpath=$GOPATH/src/$project_namespace
mkdir -p $(dirname $fullpath)
mv src $fullpath

pushd $fullpath > /dev/null
  echo -e "\n Vetting packages for potential issues..."
  go vet ${project_namespace}/...

  echo -e "\n Checking with golint..."
  golint ${project_namespace}/...

  echo -e "\n Testing packages..."
  ginkgo -r -race .
popd > /dev/null
