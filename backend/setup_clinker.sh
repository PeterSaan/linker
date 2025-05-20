#!/usr/bin/env bash

if [ -f $GOPATH/bin/clinker ]; then
	echo -e "Clinker already found in $GOPATH/bin"
	exit 0
fi
echo -e "Clinker not found. Building and moving..."
go build ./cmd/cli/clinker.go
mv ./clinker $GOPATH/bin/
exit 0
