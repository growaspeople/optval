#! /bin/bash
set -eux
PATH="$(pwd)/.built:$PATH"

# Test 2. opts should show error when both long and short option names of the same option are specified

opts init "$@"

set +e
opts def -a --opt1
set -eux

[[ "$?" != "0" ]]
