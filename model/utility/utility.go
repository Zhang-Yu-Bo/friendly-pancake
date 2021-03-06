package utility

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
)

func IsFileOrDirExist(path string) bool {
	_, err := os.Stat(path)
	// 好像不太能用 os.IsExist(err) ...在判斷上會出現問題 = =
	// 當路徑存在時， os.IsExist(err)跟 os.IsNotExist(err)同時判定為 false
	// 當路徑不存在時 os.IsExist(err) == false, os.IsNotExist(err) == true
	// 所以這邊先用 !os.IsNotExist(err)判斷
	return !os.IsNotExist(err)
}

func CreateFile(fileName string) (*os.File, error) {
	path := filepath.Dir(fileName)
	if IsFileOrDirExist(path) {
		// if file path exist
		return os.Create(fileName)
	}

	// if file path not exist
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(fileName)
}

func HashBySha256(val string) string {
	hashVal := sha256.Sum256([]byte(val))
	return hex.EncodeToString(hashVal[:])
}

func Clamp(val, min, max int) int {
	if val < min {
		val = min
	}
	if val > max {
		val = max
	}
	return val
}

func GetStringFromURL(r *http.Request, key string, defaultVal string) string {
	val := r.URL.Query().Get(key)
	if val == "" && defaultVal != "" {
		val = defaultVal
	}
	return val
}

func PixelToInt(val string) int {
	valInt, err := strconv.Atoi(strings.ReplaceAll(val, "px", ""))
	if err != nil {
		valInt = 0
		fmt.Printf("%s\n", err.Error())
	}
	return valInt
}

func OpenPngAsByte(filePath string) ([]byte, error) {
	if !IsFileOrDirExist(filePath) {
		return nil, errors.New(filePath + " does not exist.")
	}
	return os.ReadFile(filePath)
}

func SaveBytesAsPng(filePath string, data []byte) error {
	var err error
	var imgFile *os.File
	var mImg image.Image
	var fileInDir []string

	if fileInDir, err = filepath.Glob(filepath.Dir(filePath) + "/*.*"); err != nil {
		return err
	}
	if len(fileInDir) >= MaxLocalCatchOfImg {
		if err = os.RemoveAll(filepath.Dir(filePath)); err != nil {
			return err
		}
		logger.LogMessage("delete file at: " + filepath.Dir(filePath))
	}

	if imgFile, err = CreateFile(filePath); err != nil {
		return err
	}
	defer CloseFile(imgFile)

	if mImg, err = png.Decode(bytes.NewReader(data)); err != nil {
		return err
	}
	if err = png.Encode(imgFile, mImg); err != nil {
		return err
	}
	return nil
}

func CloseFile(mFile *os.File) {
	if err := mFile.Close(); err != nil {
		logger.ErrorMessage(err)
	} else {
		logger.LogMessage(mFile.Name() + " is closed.")
	}
}
