package main

import (
  "os"
)

func main() {
  var optsArgs []string = os.Args[1:] // First argument is skipped because it is program name

  if len(optsArgs) <= 0 {
    help(true)
  }

  switch optsArgs[0] {
    case "init":
      initialize(optsArgs[1:])
    case "def":
      def(optsArgs[1:])
    case "end":
      end(optsArgs[1:])
    case "--help":
      help(false)
    default:
      help(true)
  }
}

func initialize(cliArgs []string) {
  // TODO
}

func def(cliArgs []string) {
  // TODO
}

func end(cliArgs []string) {
  // TODO
}

func help(isError bool) {
  // TODO echo help

  if isError == true {
    os.Exit(1)
  }
}
