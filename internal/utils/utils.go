package utils

import (
	"strconv"
	"time"
)

func GenerateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}
