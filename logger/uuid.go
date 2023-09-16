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
		LogError(fmt.Sprintf("Unable to generate UUID: %v\n", err))
		return ""
	}
	return uuidResult.String()
}
