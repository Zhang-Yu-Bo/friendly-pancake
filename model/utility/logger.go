package utility

import (
	"fmt"
	"time"
)

func LogMessage(message string) {
	message = fmt.Sprintf("[Log][%s]: %s", time.Now(), message)
	fmt.Println(message)
}
