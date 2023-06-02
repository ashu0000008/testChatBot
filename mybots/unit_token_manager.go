package mybots

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//TokenResponse struct for json
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

var mUnitAccessToken = ""
var mAccessTokenCreate int64 = 0

func GetUnitToken() string {
	now := time.Now().Unix()

	//30分钟更新一下key
	if "" == mUnitAccessToken || now-mAccessTokenCreate > 1800 {
		mUnitAccessToken = request4Token()
		mAccessTokenCreate = time.Now().Unix()
	}

	return mUnitAccessToken
}

func request4Token() string {
	url := "https://aip.baidubce.com/oauth/2.0/token?client_id=TbI6BQ4U0Qq0TH98y5gbb82T&client_secret=56TBxle6Fp5EfHqtMU3E8HgXas58tP9d&grant_type=client_credentials"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))

	//parse json
	var resp TokenResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return resp.AccessToken
}
