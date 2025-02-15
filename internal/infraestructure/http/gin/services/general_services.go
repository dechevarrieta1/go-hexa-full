package GinServices

import (
	"encoding/json"
	"log"
)

func Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("[LOG][ERROR] Marshaling JSON: %v", err)
		return nil, err
	}
	return data, nil
}
