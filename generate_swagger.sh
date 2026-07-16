#!/bin/bash
# Generate Swagger docs
# Run this script whenever you change handler annotations
swag init --parseDependency --parseInternal -g cmd/main.go -o ./docs

