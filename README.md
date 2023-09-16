# interixa-golang-logging-library

Startner Golang Library is a logging library powered by zap logger. 
It is used to support microservice development using golang so that the logging will be synchronized into one package only

## Installing

```
go get github.com/startner/golang-logging-library.git
```

### With Version

```
go get github.com/startner/golang-logging-library.git@v1.0.0
```

## Initial Setup

When using the golang logging library in a module, set up this function.
It is needed to initialize the timezone and configuration when the logging library is used.
```golang
logger.InitLogger("Asia/Jakarta")
```

## UUID Usage

UUID is used to identify and trace the logs per request.
In the current interixa structure, this function should be put in each controller functions after the defer function
```golang
logger.SetUUID()
```

## Example Usage in 1 file only

Example Usage in main package:

```golang
package main

import (
	"gitlab.com/backend-logging-library/logging-library/logger"
)

func main() {
	logger.InitLogger("Asia/Jakarta") // only needs to be set in main / routes file
	logger.SetUUID() // needs to be put everytime a request is made so the request can be tracked
	// example usage in a regular file / any other file
	logger.LogInfo("test print info level log")
	logger.LogWarnWithFunctionName("test warn level log", "main")
}
```
