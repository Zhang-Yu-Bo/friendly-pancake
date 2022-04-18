package gasRequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
)

func closeResponse(res *http.Response) {
	if err := res.Body.Close(); err != nil {
		logger.ErrorMessage(err)
	}
}

// TODO: 可以在本地 memory中緩存 code data，就不用換個 style也要等太久
func GetCodeData(r *http.Request) ([]string, int, error) {
	var err error
	var req *http.Request
	client := &http.Client{}

	hashName := utility.GetStringFromURL(r, "code", "")
	if hashName == "" {
		return nil, http.StatusBadRequest, errors.New("there is no parameter [code]")
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

	return resultInStrSlice, http.StatusOK, nil
}

func UploadCodeData(r *http.Request) (int, error) {
	var err error
	postParamInJSON := map[string]string{}

	if err = json.NewDecoder(r.Body).Decode(&postParamInJSON); err != nil {
		return http.StatusBadRequest, err
	}

	var languageParam, codeContentParam, hashNameParam string
	var exist bool
	if codeContentParam, exist = postParamInJSON["code_content"]; !exist {
		return http.StatusBadRequest, errors.New("there is no parameter [code_content]")
	}
	languageParam = postParamInJSON["code_language"]
	hashNameParam = utility.HashBySha256(codeContentParam)

	// code 長度沒有被驗證不太好, 驗證行數
	postToGasData := url.Values{
		"hash_name":     {hashNameParam},
		"code_content":  {codeContentParam},
		"code_language": {languageParam},
	}
	var res *http.Response
	if res, err = http.PostForm(utility.GetGasUrl(), postToGasData); err != nil {
		return http.StatusInternalServerError, err
	}

	defer closeResponse(res)

	resultInJSON := map[string]string{}
	if err = json.NewDecoder(res.Body).Decode(&resultInJSON); err != nil {
		return http.StatusInternalServerError, err
	}

	var resultMsg string
	if resultMsg, exist = resultInJSON["message"]; !exist {
		return http.StatusInternalServerError, errors.New("upload code failed, there is no message return")
	}
	if strings.Contains(resultMsg, "no parameter") ||
		strings.Contains(resultMsg, "failed") {
		return http.StatusInternalServerError, errors.New(resultMsg)
	}

	return http.StatusOK, nil
}
