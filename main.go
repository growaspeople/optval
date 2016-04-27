package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  var optvalArgs []string = os.Args[1:] // First argument is skipped because it is program name
  args := []string{} // = strings.Split(optvalArgs[len(optvalArgs)-1], " ") // Last arg is arguments to analyze
  optNames := []string{} // = optvalArgs[:len(optvalArgs)-1] // Option names; e.g. --foo, -b

  //
  // Analyze given arguments to optval
  //
  for i := 0; i < len(optvalArgs); i++ {
    if optvalArgs[i] == "--" { // when reached to --
      for j := i + 1; j < len(optvalArgs); j++ {
        args = append(args, optvalArgs[j])
      }

      break
    }

    if strings.HasPrefix(optvalArgs[i], "-") {
      optNames = append(optNames, optvalArgs[i])
    } else {
      fmt.Fprintln(os.Stderr, "Error: '%[1]d' is not option; options have to starts with - or --. Also check if you quote args to analyze: optval --foo \"$@\"", args[i])
      os.Exit(1)
    }
  }

  //
  // Analyze options
  //
  for i := 0; i < len(args); i++ {
    var _optName, value string

    if (!strings.HasPrefix(args[i], "-")) { // if not option
      continue
    } else if (strings.Contains(args[i], "=")) { // --foo=bar style option
      var tmp = strings.Split(args[i], "=")
      _optName = tmp[0]
      value = tmp[1]
    } else { // --foo bar style option
      _optName = args[i]
      value = args[i+1]

      i++ // Skip next arg (== the option's value)
    }

    for _, optName := range(optNames) {
      if optName == _optName {
        os.Stdout.WriteString(value)
        os.Exit(0)
      }
    }
  }
}
