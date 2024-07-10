#!/bin/bash

LAMBDA_NAME="poc"

declare output="./dist/${LAMBDA_NAME}/${CODEBUILD_WEBHOOK_TRIGGER}/${CODEBUILD_BUILD_NUMBER}"

echo ${output}

# -tags lambda.norpc
#   reduces binary size by excluding the remote procedure call component
# -ldflags="-s -w"
#   reduces binary size by removing debug libraries
env GOOS="linux" CGO_ENABLED="0" GOARCH="arm64" GOTOOLCHAIN="local" \
    go build \
    -tags lambda.norpc \
    -ldflags="-s -w" \
    -o ${output}/bootstrap main.go

# copy the config
cp ie2-config.yml ${output}/ie2-config.yml

cd ${output}

chmod +x bootstrap

zip bootstrap.zip bootstrap