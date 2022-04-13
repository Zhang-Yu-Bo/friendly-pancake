package logger

import (
	"fmt"
	"time"
)

func LogMessage(message string) {
	message = fmt.Sprintf("[Log][%s]: %s", time.Now(), message)
	fmt.Println(message)
}

func ErrorMessage(err error) {
	message := fmt.Sprintf("[Error][%s]: %s", time.Now(), err.Error())
	fmt.Println(message)
}
