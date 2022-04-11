package templatePage

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
)

type CodePage struct {
	FontsCssUrl     string
	CssUrl          string
	Code            string
	BackgroundColor string
	ContainerColor  string
	ContainerWidth  string
	FontSize        string
}

var Page string

func init() {
	var err error
	var fileVal []byte

	fileVal, err = os.ReadFile("static/codePage.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Page = string(fileVal)
}

func (p CodePage) String() string {
	var result string
	val := reflect.Indirect(reflect.ValueOf(p))
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldVal := val.Field(i).String()
		if i == 0 {
			result += fieldName + " : " + fieldVal
		} else {
			result += ", " + fieldName + " : " + fieldVal
		}
	}
	return result
}

func Parse(data CodePage) string {
	tempPage := Page
	val := reflect.Indirect(reflect.ValueOf(data))

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldVal := val.Field(i).String()
		m1 := regexp.MustCompile("{{( *)" + fieldName + "( *)}}")
		tempPage = m1.ReplaceAllString(tempPage, fieldVal)
	}
	return tempPage
}

func (data *CodePage) Validtion() {

	data.FontsCssUrl = utility.FontsStaticUrl() + utility.DefaultFontStyle + ".css"

	if data.CssUrl == "" {
		data.CssUrl = utility.CssStaticUrl() + utility.DefaultCssStyle + ".css"
	} else {
		splitString := strings.Split(data.CssUrl, "/")
		numOfElems := len(splitString)
		if numOfElems == 0 {
			data.CssUrl = utility.CssStaticUrl() + utility.DefaultCssStyle + ".css"
		} else {
			cssFileName := strings.ReplaceAll(splitString[numOfElems-1], ".css", "")
			if _, exist := utility.CssStyle[cssFileName]; !exist {
				data.CssUrl = utility.CssStaticUrl() + utility.DefaultCssStyle + ".css"
			}
		}
	}

	colorReg := regexp.MustCompile("#[0-9A-Fa-f]{6}|#[0-9A-Fa-f]{3}")
	if !colorReg.MatchString(data.BackgroundColor) {
		data.BackgroundColor = utility.BackgroundColor
	}
	if !colorReg.MatchString(data.ContainerColor) {
		data.ContainerColor = utility.ContainerColor
	}

	pixelReg := regexp.MustCompile("[0-9]+px")
	if !pixelReg.MatchString(data.ContainerWidth) {
		data.ContainerWidth = utility.ContainerWidth
	}
	data.ContainerWidth = strconv.Itoa(utility.Clamp(utility.PixelToInt(data.ContainerWidth), 100, 1000)) + "px"
	if !pixelReg.MatchString(data.FontSize) {
		data.FontSize = utility.FontSize
	}
	data.FontSize = strconv.Itoa(utility.Clamp(utility.PixelToInt(data.FontSize), 5, 36)) + "px"

	data.Code = strings.ReplaceAll(data.Code, "\t", "    ")
}
