package logger

import (
	"encoding/json"
)

func generatePayload(model any) (string, error) {
	payloadJSON, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	return string(payloadJSON), nil
}
