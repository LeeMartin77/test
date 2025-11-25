#!/bin/bash
echo "Running core validation tests..."
go test -count=1 ./validation/core
