package main

import mylogger "github.com/jagdevs/testlib/jlogger"

func main() {
  mylogger.LogInfo("This is an info message!")
  mylogger.LogWarning("This is a warning message!")
  mylogger.LogError("This is an error message!")
}
