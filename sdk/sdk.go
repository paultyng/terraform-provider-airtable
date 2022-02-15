package sdk

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Client struct {
	httpClient *http.Client
	apiURL     string
}

type airtableRoundTripper struct {
	inner     http.RoundTripper
	userAgent string
	apiKey    string
}

func (rt *airtableRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", rt.userAgent)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rt.apiKey))

	return rt.inner.RoundTrip(req)
}

func NewClient(apiKey, userAgent string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &airtableRoundTripper{
				userAgent: userAgent,
				apiKey:    apiKey,
				inner: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
						DualStack: true,
					}).DialContext,
					MaxIdleConns:          100,
					IdleConnTimeout:       90 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				},
			},
		}
	}

	return &Client{
		httpClient: httpClient,
		apiURL:     "https://api.airtable.com/v0",
	}
}

type Record struct {
	ID          string
	CreatedTime time.Time
	Fields      map[string]string
}

type ListRecordsOptions struct {
	View string
}

type listRecordsResponseRecord struct {
	ID             string                     `json:"id"`
	RawCreatedTime string                     `json:"createdTime"`
	RawFields      map[string]json.RawMessage `json:"fields"`
}

type listRecordsResponse struct {
	Offset  string                      `json:"offset"`
	Records []listRecordsResponseRecord `json:"records"`
}

func (c *Client) ListRecords(workspaceID, table string, options *ListRecordsOptions) ([]Record, error) {
	u, _ := url.Parse(c.apiURL)
	u.Path = path.Join(u.Path, workspaceID, table)

	queryParams := make(url.Values)

	if options == nil {
		options = &ListRecordsOptions{}
	}

	if options.View != "" {
		queryParams.Add("view", options.View)
	}

	records := []Record{}

	for {
		u.RawQuery = queryParams.Encode()

		resp, err := c.httpClient.Get(u.String())
		if err != nil {
			return nil, err
		}

		body := &listRecordsResponse{}
		json.NewDecoder(resp.Body).Decode(body)

		for _, raw := range body.Records {
			created, err := time.Parse(time.RFC3339, raw.RawCreatedTime)
			if err != nil {
				return nil, err
			}
			fields := make(map[string]string)
			for key, field := range raw.RawFields {
				val, _ := field.MarshalJSON()
				fields[key] = string(val)
			}
			r := Record{
				Fields:      fields,
				ID:          raw.ID,
				CreatedTime: created,
			}
			records = append(records, r)
		}

		if body.Offset == "" {
			break
		}

		queryParams.Set("offset", body.Offset)
	}

	return records, nil
}
