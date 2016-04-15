#! /bin/sh

set -eu

# optval should receive option
[ "value1" != "$(../.built/optval "--opt1" -- "$@")" ] && echo "Failed to receive long option (--blahblah): Expected: value1, @='$@'; Actual: $(../.built/optval "--opt1" -- "$@")" && exit 1

# optval should receive short option
[ "shortoptval" != "$(../.built/optval "--shortopt" "-a" -- "$@")" ] &&  echo "Failed to receive short option (-b)" && exit 1

# optval should receive arguments
# TODO not implemented yet
# [ "arg1${IFS}arg2" != "$(optval --args)" ] && exit 1
