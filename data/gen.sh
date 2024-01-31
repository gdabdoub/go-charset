#!/bin/bash -e

go run generate.go && goimports -w .