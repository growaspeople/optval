package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
  "os/user"
  "path"
  "path/filepath"
  "strconv"
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

func def(cliArgs []string) {
  // TODO
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

func help(isError bool) {
  // TODO echo help

  if isError == true {
    os.Exit(1)
  }
}
