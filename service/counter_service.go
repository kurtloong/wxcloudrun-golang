package service

import (
	"fmt"
	"io"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func GetAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	auth_code, err := get_auth_code(r)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Fprint(w, auth_code)
}

func get_auth_code(r *http.Request) (string, error) {
	// 请求接口 http://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode
	// 请求方式 POST
	request, err := http.NewRequest("POST", "http://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode", nil)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	// {
	// 	"pre_auth_code": "Cx_Dk6qiBE0Dmx4EmlT3oRfArPvwSQ-oa3NL_fwHM7VI08r52wazoZX2Rhpz1dEw",
	// 	"expires_in": 600
	//   }

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
