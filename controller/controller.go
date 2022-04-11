package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"

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

	var err error
	var statusCode int
	var pageData templatePage.CodePage
	if statusCode, err = pageData.GetDataFromURL(r); err != nil {
		utility.ExternalErrorHandler(w, statusCode, err, utility.QRCode)
		return
	}

	pageData.Validtion()

	imgFileName := utility.HashBySha256(pageData.String())
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

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	var err error
	var statusCode int
	var pageData templatePage.CodePage
	if statusCode, err = pageData.GetDataFromURL(r); err != nil {
		utility.ExternalErrorHandler(w, statusCode, err, utility.Json)
		return
	}
	pageData.Validtion()

	w.Write([]byte(templatePage.Parse(pageData)))
}

func UploadCode(w http.ResponseWriter, r *http.Request) {
	var err error
	receive := map[string]string{}

	// get code data from request body
	if err = json.NewDecoder(r.Body).Decode(&receive); err != nil {
		utility.ExternalErrorHandler(w, http.StatusBadRequest, err, utility.Json)
		return
	}

	// check code data
	if _, exist := receive["code"]; !exist {
		mErr := errors.New("bad request, there is no code data in [receive]")
		utility.ExternalErrorHandler(w, http.StatusBadRequest, mErr, utility.Json)
		return
	}
	if _, err = base64.StdEncoding.DecodeString(receive["code"]); err != nil {
		mErr := errors.New(err.Error() +
			". bad request, the formate of the code data is not std base64(RFC 4648)")
		utility.ExternalErrorHandler(w, http.StatusBadRequest, mErr, utility.Json)
		return
	}

	// search file by the hash value of code data
	fileName := utility.HashBySha256(receive["code"])
	filePath := "static/catch/code/" + fileName + ".json"
	if utility.IsFileOrDirExist(filePath) {
		utility.ResponesByJson(w, http.StatusAccepted, "Code already exist")
		return
	}

	// if there is no code file, then create it
	var txtFile *os.File
	if txtFile, err = utility.CreateFile(filePath); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.Json)
		return
	}
	defer utility.CloseFile(txtFile)

	// json encode
	var jsonByte []byte
	if jsonByte, err = json.Marshal(receive); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.Json)
		return
	}
	// write data in json formate
	if _, err = txtFile.Write(jsonByte); err != nil {
		utility.ExternalErrorHandler(w, http.StatusInternalServerError, err, utility.Json)
		return
	}

	utility.ResponesByJson(w, http.StatusOK, fileName)
}
