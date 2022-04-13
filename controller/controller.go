package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/gasRequest"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/templatePage"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
	wk "github.com/Zhang-Yu-Bo/friendly-pancake/model/wkhtmltoimage"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "stranger"
	}

	var message string
	message += fmt.Sprintf("Hello, %s\n", name)
	message += fmt.Sprintf("OS: %s\n", runtime.GOOS)
	message += fmt.Sprintf("Max Process: %d\n", runtime.GOMAXPROCS(0))
	message += fmt.Sprintf("Your IP is: %s\n", r.RemoteAddr)
	message += fmt.Sprintf("Forwarded for: %s\n", r.Header.Get("X-FORWARDED-FOR"))

	fmt.Fprint(w, templatePage.ShowMessage(message))
}

func FaviconIco(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon.ico", http.StatusSeeOther)
}

// if error happend, show qr code
func ShowRawImage(w http.ResponseWriter, r *http.Request) {

	if templatePage.Page == "" {
		mErr := errors.New("code template is nil")
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, mErr, utility.QRCode)
		return
	}

	hashName := utility.GetStringFromURL(r, "code", "")
	if hashName == "" {
		mErr := errors.New("there is no parameter [code]")
		utility.ExternalErrorHandler(w, http.StatusBadRequest, mErr, utility.QRCode)
		return
	}

	chCodeContent := make(chan []string)
	chStatusCode := make(chan int)
	chError := make(chan error)
	go func() {
		tempCodeContent, tempStatusCode, tempErr := gasRequest.GetCodeData(r)
		chCodeContent <- tempCodeContent
		chStatusCode <- tempStatusCode
		chError <- tempErr
	}()

	var err error
	var statusCode int
	var pageData templatePage.CodePage
	if statusCode, err = pageData.GetStyleFromURL(r); err != nil {
		utility.ExternalErrorHandler(w, statusCode, err, utility.QRCode)
		return
	}
	pageData.Validtion()

	imgFileName := utility.HashBySha256(pageData.String() + ", hashName: " + hashName)
	imgFilePath := "static/catch/img/" + imgFileName + ".png"
	if utility.IsFileOrDirExist(imgFilePath) {
		var mImg []byte
		if mImg, err = utility.OpenPngAsByte(imgFilePath); err != nil {
			utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.QRCode)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(mImg)
		return
	}

	codeContent := <-chCodeContent
	statusCode = <-chStatusCode
	err = <-chError
	if err != nil {
		utility.ExternalErrorHandler(w, statusCode, err, utility.QRCode)
		return
	}
	if len(codeContent) < 2 {
		mErr := errors.New("there is no code content")
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, mErr, utility.QRCode)
		return
	}

	var codeContentBytes []byte
	if codeContentBytes, err = base64.StdEncoding.DecodeString(codeContent[1]); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.QRCode)
		return
	}
	if pageData.Code, err = url.QueryUnescape(string(codeContentBytes)); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.QRCode)
		return
	}

	html := templatePage.Parse(pageData)

	htmlToPngOptions := wk.ImageOptions{
		BinaryPath: utility.BinPath(),
		Input:      "-",
		HTML:       html,
		Format:     "png",
		Width:      utility.PixelToInt(pageData.ContainerWidth) + 50,
	}

	var htmlToPng []byte
	if htmlToPng, err = wk.GenerateImage(&htmlToPngOptions); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.QRCode)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlToPng)
	if err = utility.SaveBytesAsPng(imgFilePath, htmlToPng); err != nil {
		utility.InternalErrorHandler(err)
	}
}

func ShowMessagePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	byteMsg, err := base64.URLEncoding.DecodeString(vars["message"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, templatePage.ShowMessage(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, templatePage.ShowMessage(string(byteMsg)))
}

func TestPage(w http.ResponseWriter, r *http.Request) {

	// get Code Data From GAS
	// var err error
	// var mData []string

	// if mData, err = gasRequest.GetCodeData(r); err != nil {
	// 	fmt.Fprintf(w, "%s\n", err.Error())
	// 	return
	// }

	// for k, v := range mData {
	// 	fmt.Fprintf(w, "%d %s\n", k, v)
	// }

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	var err error
	var statusCode int
	var pageData templatePage.CodePage
	if statusCode, err = pageData.GetStyleFromURL(r); err != nil {
		utility.ExternalErrorHandler(w, statusCode, err, utility.Json)
		return
	}
	pageData.Validtion()

	w.Write([]byte(templatePage.Parse(pageData)))
}

func UploadCode(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var err error

	if statusCode, err = gasRequest.UploadCodeData(r); err != nil {
		utility.ResponesByJson(w, statusCode, err.Error())
		return
	}

	utility.ResponesByJson(w, statusCode, "upload code success")
}
