package utils

import (
	"fmt"
	"time"
)

func Nonce() string {
	return fmt.Sprintf("%v", (time.Now().Unix() * 7))
}
