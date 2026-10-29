package sub2api

import (
	"strings"

	"github.com/QuantumNous/new-api/relay/channel/newapi"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/relaykit/types"
)

type Adaptor struct {
	newapi.Adaptor
}

func (a *Adaptor) GetRequestURL(info *relaycommon.RelayInfo) (string, error) {
	if !usesAntigravityGateway(info) || info.ChannelMeta == nil {
		return a.Adaptor.GetRequestURL(info)
	}

	// Sub2API exposes Antigravity-backed models under a dedicated route prefix.
	// Keep channel configuration compatible with both the gateway root and a
	// base URL that already includes /antigravity.
	routedInfo := *info
	routedMeta := *info.ChannelMeta
	routedInfo.ChannelMeta = &routedMeta
	baseURL := strings.TrimRight(routedInfo.ChannelBaseUrl, "/")
	if !strings.HasSuffix(baseURL, "/antigravity") {
		baseURL += "/antigravity"
	}
	routedInfo.ChannelBaseUrl = baseURL
	return a.Adaptor.GetRequestURL(&routedInfo)
}

func usesAntigravityGateway(info *relaycommon.RelayInfo) bool {
	if info == nil {
		return false
	}
	model := strings.ToLower(strings.TrimSpace(info.GetUpstreamModelName()))
	if model == "" {
		model = strings.ToLower(strings.TrimSpace(info.GetOriginModelName()))
	}
	if info.RelayFormat == types.RelayFormatClaude {
		return strings.HasPrefix(model, "claude-opus-5-5") ||
			strings.HasPrefix(model, "claude-sonnet-5-5") ||
			strings.HasPrefix(model, "claude-fable-5")
	}
	if info.RelayFormat == types.RelayFormatGemini {
		return strings.HasPrefix(model, "gemini-4-argon")
	}
	return false
}

func (a *Adaptor) GetModelList() []string {
	return ModelList
}

func (a *Adaptor) GetChannelName() string {
	return ChannelName
}
