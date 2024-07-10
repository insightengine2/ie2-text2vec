$env:GOOS="linux"
$env:CGO_ENABLED="0"
$env:GOARCH="amd64"

# -tags lambda.norpc exludes the remote procedure call component of the lambda library
# which reduces the binary size

# -ldflags="-s -w" removes debug libraries from the binary, making it smaller
go build -tags lambda.norpc -ldflags="-s -w" -o ./dist/bootstrap main.go
cd ./dist
build-lambda-zip -output bootstrap.zip bootstrap
