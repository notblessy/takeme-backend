package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

// Dump :nodoc:
func Dump(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

// GenerateID :nodoc:
func GenerateID() string {
	randomID := time.Now().Nanosecond() + rand.Intn(10000)
	return strconv.Itoa(randomID)
}
