package notifications

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/szemin-ng/purecloud"
)

type Channel struct {
	ID         string `json:"id,omitempty"`
	ConnectURI string `json:"connectUri,omitempty"`
}

func CreateNotificationChannel(token purecloud.AccessToken) (resp Channel, err error) {
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/notifications/channels")

	var req *http.Request
	req, _ = http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.Token)

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
	if err = decoder.Decode(&resp); err != nil {
		return
	}

	return
}
