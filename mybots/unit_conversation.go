package mybots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var mSessionId = ""

func GetAnswerFromUnit(question string, usrId string) string {
	orgData := "UNIT received: " + question
	fmt.Println(orgData)

	response := doRequest(question, usrId)
	fmt.Println(response)

	return orgData
}

func doRequest(question string, usrId string) string {
	url := "https://aip.baidubce.com/rpc/2.0/unit/service/v3/chat?access_token=" + GetUnitToken()
	requestData := newConversationRequest(question, mSessionId, usrId)

	//post request
	jsonParamBytes, err := json.Marshal(requestData)
	if err != nil {
		err = fmt.Errorf("Marshal json error:%s ", err.Error())
		return "Internal error"
	}
	paramData := bytes.NewBuffer(jsonParamBytes)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", paramData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//todo 解析出sessionId

	return string(body)
}
