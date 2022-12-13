package utils

import "fmt"

func Print(msg string, data any) {
	fmt.Println(fmt.Sprintf("%v %#+v", msg, data))
}
