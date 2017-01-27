opts
=======

[![Build Status](https://travis-ci.org/phanect/optval.svg?branch=master)](https://travis-ci.org/phanect/optval)

Simplified getopt / getopts alternative

**Currently work in progress**

Usage
-----

e.g. Run following command:

```sh
./tetete.sh production notel --github-user phanect -b cloudsql
```

tetete.sh

```sh
github_user=$(opts "--github-user" -- "$@")
echo $github_user # phanect

branch=$(opts "--branch" "-b" -- "$@")
echo $branch # cloudsql

args=$(opts --args)
echo $args # "production notel"; values should be separated by $IFS
```
