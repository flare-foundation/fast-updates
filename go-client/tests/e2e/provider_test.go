package e2e

import (
	"fast-updates-client/provider"
	"strings"
	"sync"
)

// ScriptedProvider is an in-process provider.ValuesProvider whose values are
// fully controlled by the test scenario. Feeds without a target return nil,
// which the client treats as "feed not supported by the provider".
type ScriptedProvider struct {
	mu      sync.Mutex
	targets map[provider.FeedId]float64
}

func NewScriptedProvider() *ScriptedProvider {
	return &ScriptedProvider{targets: make(map[provider.FeedId]float64)}
}

// SetTarget sets the value returned for the feed with the given deployment
// index (see DeployedFeedId).
func (p *ScriptedProvider) SetTarget(feedIndex int, value float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.targets[DeployedFeedId(feedIndex)] = value
}

func (p *ScriptedProvider) GetValues(feeds []provider.FeedId) ([]*float64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	values := make([]*float64, len(feeds))
	for i, feed := range feeds {
		if value, ok := p.targets[feed]; ok {
			v := value
			values[i] = &v
		}
	}
	return values, nil
}

// DeployedFeedId reconstructs the FeedId the client derives for the i-th feed
// added by test_utils.Deploy: a 21-byte id with the last byte set to i+1,
// parsed as category = byte 0 and name = bytes 1..20 with trailing zero bytes
// trimmed.
func DeployedFeedId(i int) provider.FeedId {
	id := [21]byte{}
	id[20] = byte(i + 1)
	return provider.FeedId{
		Category: id[0],
		Name:     strings.TrimRight(string(id[1:]), "\x00"),
	}
}
