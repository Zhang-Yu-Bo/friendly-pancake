package utility

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
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
		logger.ErrorMessage(err)
		GolangImg = make([]byte, 0, 1)
		return
	}
}

func ResponesByJSON(w http.ResponseWriter, statusCode int, message string) {
	mapMessage := map[string]string{}
	mapMessage["message"] = message
	byteMessage, err := json.Marshal(mapMessage)
	if err != nil {
		logger.ErrorMessage(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(byteMessage)
}

func ResponesByQRCode(w http.ResponseWriter, statusCode int, message string) {
	qImg, err := qrcode.Encode(getMsgPageURL(statusCode, message), qrcode.Medium, 256)
	if err != nil {
		logger.ErrorMessage(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(GolangImg)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(qImg)
}

func ResponesByPage(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	http.Redirect(w, r, getMsgPageURL(statusCode, message), http.StatusSeeOther)
}

func getMsgPageURL(statusCode int, message string) string {
	base64Msg := base64.URLEncoding.EncodeToString([]byte(message))
	mUrl := Hostname() + "/show"
	if statusCode >= 400 {
		mUrl += "/error/"
	} else {
		mUrl += "/msg/"
	}
	return mUrl + base64Msg
}
