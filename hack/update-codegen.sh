#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
GOPATH="/home/wangb/go_projects/crddemo"
SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}

vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/bingerambo/crd-code-generation/pkg/nodecache_client github.com/bingerambo/crd-code-generation/pkg/apis \
  inspur.com:v1 \
  --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt
