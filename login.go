package purecloud

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// AccessToken contains the token for further PureCloud API calls
type AccessToken struct {
	Token  string
	ApiURL string
}

type loginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   uint32 `json:"expires_in,omitempty"`
	Error       string `json:"error,omitempty"`
}

const (
	// AustraliaRegion = mypurecloud.com.au
	AustraliaRegion string = "mypurecloud.com.au"
	// USRegion = mypurecloud.com
	USRegion string = "mypurecloud.com"
	// JapanRegion = mypurecloud.com.jp
	JapanRegion string = "mypurecloud.com.jp"
	// EuropeRegion = mypurecloud.com.ie
	EuropeRegion string = "mypurecloud.com.ie"
)

var apiURL = map[string]string{
	AustraliaRegion: "https://api.mypurecloud.com.au",
	USRegion:        "https://api.mypurecloud.com",
	JapanRegion:     "https://api.mypurecloud.com.jp",
	EuropeRegion:    "https://api.mypurecloud.com.ie",
}

// LoginWithClientCredentials logs into PureCloud uses Client Credentials Authorization and returns a token
func LoginWithClientCredentials(region string, clientID string, clientSecret string) (token AccessToken, err error) {
	var req *http.Request
	req, _ = http.NewRequest("POST", "https://login."+region+"/oauth/token", strings.NewReader("grant_type=client_credentials"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return // failed to connect?
	}

	// Connected successfully but PureCloud may return an error
	if res.StatusCode != http.StatusOK {
		var b []byte
		b, _ = ioutil.ReadAll(res.Body)
		err = errors.New(res.Status + " " + string(b))
		return
	}

	// No errors, decode JSON response body to struct
	var decoder = json.NewDecoder(res.Body)
	var l loginResponse
	if err = decoder.Decode(&l); err != nil {
		return
	}

	token.Token = l.AccessToken
	token.ApiURL = apiURL[region]
	return
}
