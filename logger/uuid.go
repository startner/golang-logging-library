package logger

import (
	"fmt"

	"github.com/google/uuid"
)

func SetUUID() {
	UuidLogger = generateUUID()
}

func generateUUID() string {
	uuidResult, err := uuid.NewRandom()
	if err != nil {
		LogCustom(ERROR, fmt.Sprintf("Unable to generate UUID: %v\n", err), "Logger.generateUUID", "")
		return ""
	}
	return uuidResult.String()
}
