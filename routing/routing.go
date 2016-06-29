package routing

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/szemin-ng/purecloud"
)

// GetQueueParams represents query parameters for getting information on a list of queues
type GetQueueParams struct {
	PageSize   int
	PageNumber int
	SortBy     string
	Name       string
	Active     bool
}

type Queue struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

type QueueEntityListing struct {
	Entities    []Queue `json:"entities,omitempty"`
	Total       int     `json:"total,omitempty"`
	PageSize    int     `json:"pageSize,omitempty"`
	PageCount   int     `json:"pageCount,omitempty"`
	PageNumber  int     `json:"pageNumber,omitempty"`
	SelfURI     string  `json:"selfUri,omitempty"`
	FirstURI    string  `json:"firstUri,omitempty"`
	LastURI     string  `json:"lastUri,omitempty"`
	PreviousURI string  `json:"previousUri,omitempty"`
	NextURI     string  `json:"nextUri,omitempty"`
}

// GetListOfQueues retrieves a list of queues configured in the organization
func GetListOfQueues(token purecloud.AccessToken, params GetQueueParams) (list QueueEntityListing, err error) {
	var u *url.URL
	u, _ = url.Parse(token.ApiURL + "/api/v2/routing/queues")

	var q url.Values
	q = u.Query()
	q.Set("active", strconv.FormatBool(params.Active))
	q.Set("name", params.Name)
	q.Set("sortBy", params.SortBy)
	q.Set("pageNumber", strconv.Itoa(params.PageNumber))
	q.Set("pageSize", strconv.Itoa(params.PageSize))
	u.RawQuery = q.Encode()

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
	if err = decoder.Decode(&list); err != nil {
		return
	}

	return
}
