#!/bin/bash
set -e

go build --tags="production" -o app main.go
tar -czvf package.tar.gz static/ app