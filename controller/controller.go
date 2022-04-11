package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"runtime"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/templatePage"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
	wk "github.com/Zhang-Yu-Bo/friendly-pancake/model/wkhtmltoimage"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s\n", name)
	fmt.Fprintf(w, "OS: %s\n", runtime.GOOS)
	fmt.Fprintf(w, "Max Process: %d\n", runtime.GOMAXPROCS(0))
	fmt.Fprintf(w, "Your IP is: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Forwarded for: %s\n", r.Header.Get("X-FORWARDED-FOR"))
}

func ShowRawImage(w http.ResponseWriter, r *http.Request) {

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	codeFileName := utility.GetStringFromURL(r, "code", "")
	if codeFileName == "" {
		fmt.Fprintln(w, "There is no parameter [code]")
		return
	}
	codeFilePath := "static/catch/code/" + codeFileName + ".json"
	if !utility.IsFileOrDirExist(codeFilePath) {
		fmt.Fprintln(w, "There is no code file")
		return
	}

	var codeFile *os.File
	var err error
	defer func() {
		codeFile.Close()
	}()
	if codeFile, err = os.Open(codeFilePath); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	mCode := map[string]string{}
	if err = json.NewDecoder(codeFile).Decode(&mCode); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	codeContent := mCode["code"]
	if codeContent == "" {
		codeContent = utility.DefaultCode
	}

	var tempByte []byte
	if tempByte, err = base64.StdEncoding.DecodeString(codeContent); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	codeContent = string(tempByte)

	backgroundColor := utility.GetStringFromURL(r, "backgroundColor", utility.BackgroundColor)
	containerColor := utility.GetStringFromURL(r, "containerColor", utility.ContainerColor)
	containerWidth := utility.GetStringFromURL(r, "containerWidth", utility.ContainerWidth)
	fontSize := utility.GetStringFromURL(r, "fontSize", utility.FontSize)
	cssStyle := utility.GetStringFromURL(r, "cssStyle", utility.DefaultCssStyle)

	data := templatePage.CodePage{
		FontsCssUrl:     utility.FontsStaticUrl() + utility.DefaultFontStyle + ".css",
		CssUrl:          utility.CssStaticUrl() + cssStyle + ".css",
		Code:            codeContent,
		BackgroundColor: backgroundColor,
		ContainerColor:  containerColor,
		ContainerWidth:  containerWidth,
		FontSize:        fontSize,
	}

	data.Validtion()

	imgFileName := utility.HashBySha256(data.String())
	imgFilePath := "static/catch/img/" + imgFileName + ".png"
	if utility.IsFileOrDirExist(imgFilePath) {
		var err error
		var mImg image.Image
		if mImg, err = utility.OpenPngAsImage(imgFilePath); err != nil {
			fmt.Fprintf(w, "%s\n", err.Error())
			return
		}
		if err = png.Encode(w, mImg); err != nil {
			fmt.Fprintf(w, "%s\n", err.Error())
			return
		}
		return
	}

	html := templatePage.Parse(data)

	c := wk.ImageOptions{
		BinaryPath: utility.BinPath(),
		Input:      "-",
		HTML:       html,
		Format:     "png",
		Width:      utility.PixelToInt(containerWidth) + 50,
	}

	if out, err := wk.GenerateImage(&c); err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
	} else {
		w.Write(out)
		if err = utility.SaveBytesAsPng(imgFilePath, out); err != nil {
			fmt.Fprintf(w, "%s\n", err.Error())
		}
	}

}

func TestPage(w http.ResponseWriter, r *http.Request) {

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	codeFileName := utility.GetStringFromURL(r, "code", "")
	if codeFileName == "" {
		fmt.Fprintln(w, "There is no parameter [code]")
		return
	}
	codeFilePath := "static/catch/code/" + codeFileName + ".json"
	if !utility.IsFileOrDirExist(codeFilePath) {
		fmt.Fprintln(w, "There is no code file")
		return
	}

	var codeFile *os.File
	var err error
	defer func() {
		codeFile.Close()
	}()
	if codeFile, err = os.Open(codeFilePath); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	mCode := map[string]string{}
	if err = json.NewDecoder(codeFile).Decode(&mCode); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	codeContent := mCode["code"]
	if codeContent == "" {
		codeContent = utility.DefaultCode
	}

	var tempByte []byte
	if tempByte, err = base64.StdEncoding.DecodeString(codeContent); err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	codeContent = string(tempByte)

	backgroundColor := utility.GetStringFromURL(r, "backgroundColor", utility.BackgroundColor)
	containerColor := utility.GetStringFromURL(r, "containerColor", utility.ContainerColor)
	containerWidth := utility.GetStringFromURL(r, "containerWidth", utility.ContainerWidth)
	fontSize := utility.GetStringFromURL(r, "fontSize", utility.FontSize)
	cssStyle := utility.GetStringFromURL(r, "cssStyle", utility.DefaultCssStyle)

	data := templatePage.CodePage{
		FontsCssUrl:     utility.FontsStaticUrl() + utility.DefaultFontStyle + ".css",
		CssUrl:          utility.CssStaticUrl() + cssStyle + ".css",
		Code:            codeContent,
		BackgroundColor: backgroundColor,
		ContainerColor:  containerColor,
		ContainerWidth:  containerWidth,
		FontSize:        fontSize,
	}

	data.Validtion()

	w.Write([]byte(templatePage.Parse(data)))
}

func UploadCode(w http.ResponseWriter, r *http.Request) {
	var err error
	receive := map[string]string{}

	if err = json.NewDecoder(r.Body).Decode(&receive); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	fileName := utility.HashBySha256(receive["code"])
	filePath := "static/catch/code/" + fileName + ".json"
	if utility.IsFileOrDirExist(filePath) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Code already exist")
		return
	}

	var txtFile *os.File
	if txtFile, err = utility.CreateFile(filePath); err != nil {
		txtFile.Close()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	var jsonByte []byte
	if jsonByte, err = json.Marshal(receive); err != nil {
		txtFile.Close()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	if _, err = txtFile.Write(jsonByte); err != nil {
		txtFile.Close()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	txtFile.Close()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", receive)
}
