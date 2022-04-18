package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/gasRequest"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
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
		errMsg := "code template is nil"
		logger.ErrorMessage(errors.New(errMsg))
		utility.ResponseByQRCode(w, http.StatusInternalServerError, errMsg)
		return
	}

	hashName := utility.GetStringFromURL(r, "code", "")
	if hashName == "" {
		errMsg := "there is no parameter [code]"
		logger.ErrorMessage(errors.New(errMsg))
		utility.ResponseByQRCode(w, http.StatusBadRequest, errMsg)
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
		logger.ErrorMessage(err)
		utility.ResponseByQRCode(w, statusCode, err.Error())
		return
	}
	pageData.Validtion()

	imgFileName := utility.HashBySha256(pageData.String() + ", hashName: " + hashName)
	imgFilePath := "static/catch/img/" + imgFileName + ".png"
	if utility.IsFileOrDirExist(imgFilePath) {
		var mImg []byte
		if mImg, err = utility.OpenPngAsByte(imgFilePath); err != nil {
			logger.ErrorMessage(err)
			utility.ResponseByQRCode(w, http.StatusInternalServerError, err.Error())
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
		logger.ErrorMessage(err)
		utility.ResponseByQRCode(w, statusCode, err.Error())
		return
	}
	if len(codeContent) < 2 {
		errMsg := "there is no code content"
		logger.ErrorMessage(errors.New(errMsg))
		utility.ResponseByQRCode(w, http.StatusInternalServerError, errMsg)
		return
	}
	pageData.Code = codeContent[1]

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
		logger.ErrorMessage(err)
		utility.ResponseByQRCode(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlToPng)
	if err = utility.SaveBytesAsPng(imgFilePath, htmlToPng); err != nil {
		logger.ErrorMessage(err)
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

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
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
		logger.ErrorMessage(err)
		utility.ResponseByPage(w, r, statusCode, err.Error())
		return
	}

	codeContent := <-chCodeContent
	statusCode = <-chStatusCode
	err = <-chError
	if err != nil {
		logger.ErrorMessage(err)
		utility.ResponseByPage(w, r, statusCode, err.Error())
		return
	}
	if len(codeContent) < 2 {
		errMsg := "there is no code content"
		logger.ErrorMessage(errors.New(errMsg))
		utility.ResponseByPage(w, r, http.StatusInternalServerError, errMsg)
		return
	}
	pageData.Code = codeContent[1]

	w.Write([]byte(templatePage.Parse(pageData)))
}

func ShowCodeContent(w http.ResponseWriter, r *http.Request) {
	var codeContent []string
	var statusCode int
	var err error
	if codeContent, statusCode, err = gasRequest.GetCodeData(r); err != nil {
		logger.ErrorMessage(err)
		utility.ResponseByJSON(w, statusCode, err.Error())
		return
	}
	if len(codeContent) < 2 {
		errMsg := "there is no code content"
		logger.ErrorMessage(errors.New(errMsg))
		utility.ResponseByJSON(w, http.StatusInternalServerError, errMsg)
		return
	}

	utility.ResponseByJSON(w, http.StatusOK, codeContent[1])
}

func UploadCode(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var hashName string
	var err error

	if statusCode, hashName, err = gasRequest.UploadCodeData(r); err != nil {
		logger.ErrorMessage(err)
		utility.ResponseByJSON(w, statusCode, err.Error())
		return
	}

	utility.ResponseByJSON(w, statusCode, hashName)
}
