package provider

import (
	"bytes"
	"encoding/json"
	"fast-updates-client/logger"
	"fmt"
	"io"
	"net/http"
)

type FeedValue struct {
	Feed  FeedId  `json:"feed"`
	Value float64 `json:"value"`
}

type HttpValuesProvider struct {
	baseUrl    string
	httpClient *http.Client
}

func (p *HttpValuesProvider) GetValues(feeds []FeedId) ([]float64, error) {
	feedValues, err := p.GetFeedValues(feeds)
	if err != nil {
		return nil, err
	}
	values, err := sortFeedValues(feeds, feedValues)
	if err != nil {
		return nil, err
	}

	return values, err
}

func NewHttpValueProvider(baseUrl string) *HttpValuesProvider {
	logger.Info("Creating new feed values provider")

	return &HttpValuesProvider{
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
	}
}

// Define the struct for the payload
type FeedValuesRequest struct {
	Feeds []FeedId `json:"feeds"`
}

type FeedId struct {
	Category byte   `json:"category"`
	Name     string `json:"name"`
}

type FeedValuesResponse struct {
	VotingRoundId int         `json:"votingRoundId"`
	Data          []FeedValue `json:"data"`
}

func (c *HttpValuesProvider) post(endpoint string, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.baseUrl+endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s (Code: %d)", resp.Status, resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func (c *HttpValuesProvider) GetFeedValues(feeds []FeedId) ([]FeedValue, error) {
	req, err := json.Marshal(
		FeedValuesRequest{
			Feeds: feeds,
		},
	)
	if err != nil {
		return nil, err
	}

	// TODO: Should we specify voting round id instead of 0?
	body, err := c.post("/feed-values/0", req)
	if err != nil {
		return nil, err
	}

	res := FeedValuesResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
