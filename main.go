package main

import (
	"fmt"

	"github.com/startner/golang-logging-library/logger"
)

func main() {
	logger.InitLogger("Asia/Makassar")

	logger.SetUUID()
	logger.LogInfo("testdoang")
	logger.LogWarnWithFunctionName("nahloh warning", "main")
	fmt.Println("============================================================")
	logger.SetUUID()
	logger.LogInfo("aaa")
	logger.LogWarnWithFunctionName("bbb", "main")
}
