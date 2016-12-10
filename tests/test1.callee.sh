#! /bin/bash
set -eux
PATH="$(pwd)/.built:$PATH"

# Test 1. opts should receive option

opts init "$@"

ARG1="$(opts def --opt1)" # Option without equal ("="); e.g. --opt1 value1
ARG2="$(opts def --opt2)" # Option with equal ("="); e.g. --opt2=value2
ARG3="$(opts def -a)" # Short option name; e.g. -a value3
# Test if opts can receive short option name value
# if short and long option names are both specified; e.g. -b value4
ARG4="$(opts def -b --opt4)"
# Test if opts can receive long option name value
# if short and long option names are both specified; e.g. --opt5 value5
ARG5="$(opts def -c --opt5)"

opts end

[[ "$ARG1" = "value1" ]]
[[ "$ARG2" = "value2" ]]
[[ "$ARG3" = "value3" ]]
[[ "$ARG4" = "value4" ]]
[[ "$ARG5" = "value5" ]]
