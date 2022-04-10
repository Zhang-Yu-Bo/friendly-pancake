package utility

import "runtime"

func Hostname() string {
	if runtime.GOOS == "linux" {
		return "https://friendly-pancake.herokuapp.com"
	}
	return "http://localhost"
}
