package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

type Sort struct {
	Field string
	Type  string
}

func Dump(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func GenerateID() string {
	randomID := time.Now().Nanosecond() + rand.Intn(10000)
	return strconv.Itoa(randomID)
}
