package gasRequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
)

// map[ hashName ] codeContent
var codeContentCatch map[string]string

func init() {
	codeContentCatch = make(map[string]string)
}

func closeResponse(res *http.Response) {
	if err := res.Body.Close(); err != nil {
		logger.ErrorMessage(err)
	}
}

func getCodeContentInLocal(hashName string) string {
	if result, exist := codeContentCatch[hashName]; exist {
		return result
	}
	return ""
}

func setCodeContentInLocal(hashName, codeContent string) {
	if _, exist := codeContentCatch[hashName]; exist {
		return
	}
	if len(codeContent) >= utility.MaxLocalCatchOfCode {
		codeContentCatch = make(map[string]string)
	}
	codeContentCatch[hashName] = codeContent
}

func isCodeContentValid(codeContent string) error {
	if numOfNewLines := strings.Count(codeContent, "\n"); numOfNewLines > utility.MaxCodeLines {
		return errors.New("too many lines")
	}
	codeReg := regexp.MustCompile("(<span class=\"token[a-zA-Z0-9-_ ]*\">)|(</span>)")
	codeText := codeReg.ReplaceAllString(codeContent, "")
	if len(codeText) > utility.MaxCodeLength {
		return errors.New("too many text")
	}
	return nil
}

func GetCodeData(r *http.Request) ([]string, int, error) {
	var err error
	var req *http.Request
	client := &http.Client{}

	hashName := utility.GetStringFromURL(r, "code", "")
	if hashName == "" {
		return nil, http.StatusBadRequest, errors.New("there is no parameter [code]")
	}

	if localCodeContent := getCodeContentInLocal(hashName); localCodeContent != "" {
		return []string{hashName, localCodeContent}, http.StatusOK, nil
	}

	mURL := fmt.Sprintf(utility.GetGasUrl()+"?hash_name=%s", hashName)
	if req, err = http.NewRequest("GET", mURL, nil); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer closeResponse(res)

	resultInJSON := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&resultInJSON); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if resultInJSON["message"] == "there is no code data" {
		return nil, http.StatusBadRequest, errors.New("there is no code data")
	}

	var parseResult []interface{}
	var parseOk bool
	if parseResult, parseOk = resultInJSON["message"].([]interface{}); !parseOk {
		return nil, http.StatusInternalServerError, errors.New("resultJson parse to []interface{} error")
	}

	var resultInStrSlice []string
	for _, v := range parseResult {
		var vStr string
		if vStr, parseOk = v.(string); !parseOk {
			return nil, http.StatusInternalServerError, errors.New("parseResult parse to string error")
		}
		resultInStrSlice = append(resultInStrSlice, vStr)
	}

	setCodeContentInLocal(hashName, resultInStrSlice[1])

	return resultInStrSlice, http.StatusOK, nil
}

func UploadCodeData(r *http.Request) (int, string, error) {
	var err error
	postParamInJSON := map[string]string{}

	if err = json.NewDecoder(r.Body).Decode(&postParamInJSON); err != nil {
		return http.StatusBadRequest, "", err
	}

	var languageParam, codeContentParam, hashNameParam string
	var exist bool
	if codeContentParam, exist = postParamInJSON["code_content"]; !exist {
		return http.StatusBadRequest, "", errors.New("there is no parameter [code_content]")
	}
	if err = isCodeContentValid(codeContentParam); err != nil {
		return http.StatusBadRequest, "", err
	}
	languageParam = postParamInJSON["code_language"]
	hashNameParam = utility.HashBySha256(codeContentParam)

	postToGasData := url.Values{
		"hash_name":     {hashNameParam},
		"code_content":  {codeContentParam},
		"code_language": {languageParam},
	}
	var res *http.Response
	if res, err = http.PostForm(utility.GetGasUrl(), postToGasData); err != nil {
		return http.StatusInternalServerError, "", err
	}

	defer closeResponse(res)

	resultInJSON := map[string]string{}
	if err = json.NewDecoder(res.Body).Decode(&resultInJSON); err != nil {
		return http.StatusInternalServerError, "", err
	}

	var resultMsg string
	if resultMsg, exist = resultInJSON["message"]; !exist {
		return http.StatusInternalServerError, "", errors.New("upload code failed, there is no message return")
	}
	if strings.Contains(resultMsg, "no parameter") ||
		strings.Contains(resultMsg, "failed") {
		return http.StatusInternalServerError, "", errors.New(resultMsg)
	}

	setCodeContentInLocal(hashNameParam, codeContentParam)

	return http.StatusOK, hashNameParam, nil
}
