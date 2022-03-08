#!/bin/bash

IP=$1
PORT=$2

filepath=$(cd "$(dirname "$0")"; pwd)

export PPROF_TMPDIR=${filepath}/profile
export PPROF_BINARY_PATH=${filepath}/profile

go tool pprof -raw http://$1:$2/debug/pprof/profile?seconds=5
