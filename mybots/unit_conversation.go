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
	fmt.Println("UNIT received: " + response)

	return response
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
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	//todo 解析出sessionId
	//parse json string
	var responseData ConversationResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	if responseData.ErrorCode != 0 {
		fmt.Println(responseData.ErrorCode)
		return responseData.ErrorMsg
	}

	if len(responseData.Result.Responses) == 0 {
		fmt.Println("no response")
		return "no response"
	}

	if len(responseData.Result.Responses[0].Actions) == 0 {
		fmt.Println("no actions")
		return "no actions"
	}

	mSessionId = responseData.Result.SessionId
	return responseData.Result.Responses[0].Actions[0].Say
}
