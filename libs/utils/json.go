package baseUtils

import (
	"encoding/json"
	"log"
)

func ParseJSON(data string, v any) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		log.Printf("Cannot parse from JSON: %v", err)
		return err
	}
	return nil
}

func EncodeJSON(v any) (string, error) {
	res, err := json.Marshal(v)
	if err != nil {
		log.Printf("Cannot encode to JSON: %v", err)
		return "", err
	}
	return string(res), err
}
