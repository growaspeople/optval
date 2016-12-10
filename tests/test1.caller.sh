#! /bin/bash

set -eux

TESTS_ROOT="$(readlink --canonicalize "$(dirname "${BASH_SOURCE[0]}")/../tests")"

# Test 1. opts should receive option

"$TESTS_ROOT/test1.callee.sh" --opt1 value1 --opt2=value2 -a value3 -b value4 --opt5 value5
