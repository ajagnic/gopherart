#!/bin/bash
env GOOS=js GOARCH=wasm go build -o static/js/main.wasm
