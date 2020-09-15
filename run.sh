#!/bin/bash

go run cmd/face-detection-app-server/main.go \
  --port=5555 \
  --pigo-cascade-file=internal/pigo/cascade/facefinder \
  --pigo-puploc-file=internal/pigo/cascade/puploc \
  --pigo-flploc-dir=internal/pigo/cascade/lps \
  --http-client-timeout=5s
