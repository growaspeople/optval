#! /bin/bash

set -eux

TESTS_ROOT="$(readlink --canonicalize "$(dirname "${BASH_SOURCE[0]}")/../tests")"

# Test 2. opts should show error when both long and short option names of the same option are specified

"$TESTS_ROOT/test2.callee.sh" --opt1 value1 -a value2
