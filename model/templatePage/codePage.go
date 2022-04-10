package templatePage

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
)

type CodePage struct {
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
