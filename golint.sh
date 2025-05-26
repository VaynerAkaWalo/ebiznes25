#!/bin/zsh

set -e

cd zadanie_04

exec golangci-lint run ./...
