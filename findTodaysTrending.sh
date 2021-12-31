#!/bin/sh

curl https://github.com/trending > results 2>&1

go run .