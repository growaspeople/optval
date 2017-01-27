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
