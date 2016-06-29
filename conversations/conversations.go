package conversations

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/szemin-ng/purecloud"
)

// GetConversation gets details about a conversation that is live or completed. Sends and receives response from /api/v2/conversations/{conversationId}
func GetConversation(token purecloud.AccessToken, conversationID string) (response Conversation, err error) {
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/conversations/" + conversationID)

	var req *http.Request
	req, _ = http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.Token)

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return // failed to connect?
	}

	// IMPORTANT: Close response body to allow reuse and avoid memory leak
	defer res.Body.Close()

	// Connected successfully but PureCloud may return an error
	if res.StatusCode != http.StatusOK {
		var b []byte
		b, _ = ioutil.ReadAll(res.Body)
		err = errors.New(res.Status + " " + string(b))
		return
	}

	// No errors, decode JSON response body to struct
	var decoder = json.NewDecoder(res.Body)
	if err = decoder.Decode(&response); err != nil {
		return
	}

	return
}
