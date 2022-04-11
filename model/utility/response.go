package utility

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/skip2/go-qrcode"
)

type ResponseMethod int

const (
	Json   ResponseMethod = 0
	QRCode ResponseMethod = 1
	Page   ResponseMethod = 2
)

var GolangImg []byte

func init() {
	var err error
	if GolangImg, err = os.ReadFile("static/golang.jpg"); err != nil {
		InternalErrorHandler(err)
		GolangImg = make([]byte, 0, 1)
		return
	}
}

func ResponesByJson(w http.ResponseWriter, statusCode int, message string) {
	mapMessage := map[string]string{}
	mapMessage["message"] = message
	byteMessage, err := json.Marshal(mapMessage)
	if err != nil {
		InternalErrorHandler(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(byteMessage)
}

func ResponesByQRCode(w http.ResponseWriter, statusCode int, message string) {
	base64Msg := base64.URLEncoding.EncodeToString([]byte(message))
	mUrl := Hostname() + "/show"
	if statusCode >= 400 {
		mUrl += "/error/"
	} else {
		mUrl += "/msg/"
	}

	qImg, err := qrcode.Encode(mUrl+base64Msg, qrcode.Medium, 256)
	if err != nil {
		InternalErrorHandler(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(GolangImg)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(qImg)
}

func ResponesByPage(w http.ResponseWriter, statusCode int, message string) {

}
