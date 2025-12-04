package files

import (
	"encoding/json"
	"fmt"
	"log"
)



func DecodeJsonTypeFromBytes[T any](bytes []byte) (T, error) {
	var v T
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return v, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return v, nil
}

func EncodeJsonTypeToBytes[T any](v T) ([]byte, error) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}
	return bytes, nil
}

func PrintJsonType[T any](v T) {
	bytes, err := EncodeJsonTypeToBytes(v)
	if err != nil {
		log.Fatalf("failed to encode json: %v", err)
	}
	fmt.Println(string(bytes))
}