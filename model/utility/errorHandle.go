package utility

import (
	"fmt"
	"net/http"
	"time"
)

func InternalErrorHandler(err error) {
	message := fmt.Sprintf("[Error][%s]: %s", time.Now(), err.Error())
	fmt.Println(message)
}

func ExternalErrorHandler(w http.ResponseWriter, statusCode int, err error, method ResponseMethod) {
	InternalErrorHandler(err)
	if method == Json {
		ResponesByJson(w, statusCode, err.Error())
	} else if method == QRCode {
		ResponesByQRCode(w, statusCode, err.Error())
	} else if method == Page {
		ResponesByPage(w, statusCode, err.Error())
	}
}
