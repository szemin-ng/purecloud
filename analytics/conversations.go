package analytics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/szemin-ng/purecloud"
)

func QueryConversationAggregates(token purecloud.AccessToken, query purecloud.AggregationQuery) (response purecloud.AggregateQueryResponse, err error) {
	var b []byte
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/analytics/conversations/aggregates/query")

	var req *http.Request
	b, _ = json.Marshal(query)
	req, _ = http.NewRequest("POST", u.String(), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.Token)

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return // failed to connect?
	}

	// Connected successfully but PureCloud may return an error
	if res.StatusCode != http.StatusOK {
		b, _ = ioutil.ReadAll(res.Body)
		err = errors.New(res.Status + " " + string(b))
		return
	}

	// No errors, decode JSON response body to struct
	var decoder = json.NewDecoder(res.Body)
	decoder.UseNumber()
	if err = decoder.Decode(&response); err != nil {
		return
	}

	return
}
