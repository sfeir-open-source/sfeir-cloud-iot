#!/bin/sh

go install github.com/cespare/reflex@latest

exec "$@"
