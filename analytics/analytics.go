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

func QueryConversationAggregates(token purecloud.AccessToken, query AggregationQuery) (response AggregateQueryResponse, err error) {
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

	// IMPORTANT: Close response body to allow reuse and avoid memory leak 
	defer res.Body.Close()

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

// QueryConversationDetails retrieves list of conversations that match the query. Sends and receives response from /api/v2/analytics/conversations/details/query
func QueryConversationDetails(token purecloud.AccessToken, query ConversationQuery) (response AnalyticsConversationQueryResponse, err error) {
	var b []byte
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/analytics/conversations/details/query")

	var req *http.Request
	b, _ = json.Marshal(query)
	req, _ = http.NewRequest("POST", u.String(), bytes.NewReader(b))
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

func QueryQueueObservations(token purecloud.AccessToken, query ObservationQuery) (response ObservationQueryResponse, err error) {
	var b []byte
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/analytics/queues/observations/query")

	var req *http.Request
	b, _ = json.Marshal(query)
	req, _ = http.NewRequest("POST", u.String(), bytes.NewReader(b))
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
