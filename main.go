package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "os/user"
  "path"
  "path/filepath"
  "strconv"
  "strings"
)

var jsonFile string

func main() {
  var optsArgs []string = os.Args[1:] // First argument is skipped because it is program name

  // Define global variables
  scriptPid := os.Getppid() // Process ID of shell script which is calling opts
  currentUser, err := user.Current()
  catch(err)
  jsonFile = path.Join("/var/run/user", currentUser.Uid, "opts", strconv.Itoa(scriptPid) + ".json")

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







  //
  // var optvalArgs []string = os.Args[1:] // First argument is skipped because it is program name
  // /** All arguments given to the command. (== $@) */
  // args := []string{} // = strings.Split(optvalArgs[len(optvalArgs)-1], " ") // Last arg is arguments to analyze
  // /** Option names the user wants to get value */
  // optNames := []string{} // = optvalArgs[:len(optvalArgs)-1] // Option names; e.g. --foo, -b
  //
  // //
  // // Analyze given arguments to optval
  // //
  // for i := 0; i < len(optvalArgs); i++ {
  //   if optvalArgs[i] == "--" { // when reached to --
  //     for j := i + 1; j < len(optvalArgs); j++ {
  //       args = append(args, optvalArgs[j])
  //     }
  //
  //     break
  //   }
  //
  //   if strings.HasPrefix(optvalArgs[i], "-") {
  //     optNames = append(optNames, optvalArgs[i])
  //   } else {
  //     fmt.Fprintln(os.Stderr, "Error: '%[1]d' is not option; options have to starts with - or --. Also check if you quote args to analyze: optval --foo \"$@\"", args[i])
  //     os.Exit(1)
  //   }
  // }
  //
  // //
  // // Analyze options
  // //
  // for i := 0; i < len(args); i++ {
  //   var _optName, value string
  //
  //   if (!strings.HasPrefix(args[i], "-")) { // if not option
  //     continue
  //   } else if (strings.Contains(args[i], "=")) { // --foo=bar style option
  //     var tmp = strings.Split(args[i], "=")
  //     _optName = tmp[0]
  //     value = tmp[1]
  //   } else { // --foo bar style option
  //     _optName = args[i]
  //     value = args[i+1]
  //
  //     i++ // Skip next arg (== the option's value)
  //   }
  //
  //   for _, optName := range(optNames) {
  //     if optName == _optName {
  //       os.Stdout.WriteString(value)
  //       os.Exit(0)
  //     }
  //   }
  // }
}

func initialize(cliArgs []string) {
  jsonStr, err := json.Marshal(cliArgs)
  catch(err)

  os.Mkdir(filepath.Dir(jsonFile), 0700)

  err = ioutil.WriteFile(jsonFile, jsonStr, 0600)
  catch(err)
}

func def(options []string) {
  /** Value to return */
  var value string

  jsonByte, err := ioutil.ReadFile(jsonFile)
  catch(err)

  for _, option := range options {
    var cliArgs []string
    var index int

    json.Unmarshal(jsonByte, &cliArgs)

    //   TODO option が- 又は -- で始まらない場合、- が1つなのにその後に文字列が続いている場合、- が2つなのにその後に1文字しかない場合にエラー

    index = indexOf(option, cliArgs)

    if 0 <= index { // if option exists in cliArgs
      if len(value) > 0 {
        value = cliArgs[index + 1]
        continue // Search next option to validate if both short and long options are not specified redundantly
      } else { // value already assigned == short and long options are specified redundantly
        die("Duplicate options: you can only specify one of " + strings.Join(options, ", "))
      }
    }
  }
}

func end(cliArgs []string) {
  // TODO
}

func catch(err error) {
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}

func die(message string) {
  fmt.Fprintln(os.Stderr, message)
  os.Exit(1)
}

func help(isError bool) {
  // TODO echo help

  if isError == true {
    os.Exit(1)
  }
}

/**
 * Get index of `search` in `array`
 */
func indexOf(search string, array []string) int {
  for i, element := range array {
    if search == element {
      return i
    }
  }

  // If there is no matching element
  return -1
}
