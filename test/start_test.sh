#! /bin/bash

set -eu

cd $(dirname ${BASH_SOURCE[0]})/../
go build -o ./.built/optval

cd $(dirname ${BASH_SOURCE[0]})

./target.sh arg1 arg2 --opt1 value1 -a shortoptval
