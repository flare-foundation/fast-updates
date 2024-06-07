package provider

import (
	"bytes"
	"encoding/json"
	"fast-updates-client/logger"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type FeedId struct {
	Category byte   `json:"category"`
	Name     string `json:"name"`
}

type FeedValue struct {
	Feed  FeedId   `json:"feed"`
	Value *float64 `json:"value"`
}

type FeedValuesRequest struct {
	Feeds []FeedId `json:"feeds"`
}

type FeedValuesResponse struct {
	VotingRoundId int         `json:"votingRoundId"`
	Data          []FeedValue `json:"data"`
}

type HttpValuesProvider struct {
	Url        string
	httpClient *http.Client
}

func NewHttpValueProvider(url string) *HttpValuesProvider {
	logger.Info("Creating new feed values provider")

	return &HttpValuesProvider{
		Url:        url,
		httpClient: http.DefaultClient,
	}
}

func (p *HttpValuesProvider) GetValues(feeds []FeedId) ([]*float64, error) {
	feedValues, err := p.GetFeedValues(feeds)
	if err != nil {
		return nil, errors.Wrap(err, "error getting feed values")
	}
	values := sortFeedValues(feeds, feedValues)

	return values, err
}

func (p *HttpValuesProvider) GetFeedValues(feeds []FeedId) ([]FeedValue, error) {
	req, err := json.Marshal(
		FeedValuesRequest{
			Feeds: feeds,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling request")
	}

	body, err := p.post(req)
	if err != nil {
		return nil, errors.Wrap(err, "error posting feed value request")
	}

	res := FeedValuesResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}

	return res.Data, nil
}

func (p *HttpValuesProvider) post(requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", p.Url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error sending HTTP request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response status %s (Code: %d)", resp.Status, resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}
	return responseBody, nil
}
