package utility

import (
	"os"
	"runtime"
)

var CssStyle = map[string]string{
	"TomorrowNight-dark":   "TomorrowNight-dark",
	"Coy-light":            "Coy-light",
	"Default-light":        "Default-light",
	"Funky-dark":           "Funky-dark",
	"Okaidia-dark":         "Okaidia-dark",
	"SolarizedLight-light": "SolarizedLight-light",
	"Twilight-dark":        "Twilight-dark",
}

const (
	BackgroundColor  = "#2885D3"
	ContainerColor   = "#151718"
	ContainerWidth   = "700px"
	FontSize         = "18px"
	DefaultCssStyle  = "TomorrowNight-dark"
	DefaultFontStyle = "fontsFace"
)

const (
	MaxLocalCatchOfImg  = 50
	MaxCodeLength       = 2000
	MaxCodeLines        = 50
	MaxLocalCatchOfCode = 50
)

func Hostname() string {
	if runtime.GOOS == "linux" {
		return "https://friendly-pancake.herokuapp.com"
	}
	return "http://localhost"
}

func CssStaticUrl() string {
	return Hostname() + "/static/css/"
}

func FontsStaticUrl() string {
	return Hostname() + "/static/fonts/"
}

func GetGasUrl() string {
	url := os.Getenv("GAS_URL")
	if url == "" {
		url = "YOUR_GOOGLE_APPLACATION_SCRIPT_URL"
	}
	return url
}

func BinPath() string {
	if runtime.GOOS == "linux" {
		return "./bin/wkhtmltoimage"
	}
	return "C:\\Users\\Lykoi\\Desktop\\html2image-master\\wkhtmltopdf\\bin\\wkhtmltoimage.exe"
}
