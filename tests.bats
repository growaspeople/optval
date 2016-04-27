#!/usr/bin/env bats

setup() {
  TMP_SCRIPT="$(mktemp --directory)/tmp.sh"
  touch "$TMP_SCRIPT"
  chmod +x "$TMP_SCRIPT"
}

teardown() {
  rm -rf "$(dirname $TMP_SCRIPT)"
}

@test "optval should receive option" {
  echo $(pwd)'/.built/optval "--opt1" -- "$@"' > "$TMP_SCRIPT"

  result="$($TMP_SCRIPT arg1 arg2 --opt1 value1 --opt2=value2 -a shortoptval)"

  [ "$result" = "value1" ]
}

@test "optval should receive option with equal (--opt2=value2 style)" {
  echo $(pwd)'/.built/optval "--opt2" -- "$@"' > "$TMP_SCRIPT"

  result="$($TMP_SCRIPT arg1 arg2 --opt1 value1 --opt2=value2 -a shortoptval)"

  [ "$result" = "value2" ]
}

@test "optval should receive option with equal (--opt2=value2 style)" {
  echo $(pwd)'/.built/optval "--shortopt" "-a" -- "$@"' > "$TMP_SCRIPT"

  result="$($TMP_SCRIPT arg1 arg2 --opt1 value1 --opt2=value2 -a shortoptval)"

  [ "$result" = "shortoptval" ]
}

# optval should receive arguments
# TODO not implemented yet
# [ "arg1${IFS}arg2" != "$(optval --args)" ] && exit 1
