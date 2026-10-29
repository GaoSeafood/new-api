package sub2api

import (
	"encoding/json"
	"testing"

	"github.com/QuantumNous/new-api/constant"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/relaykit/dto"
	"github.com/QuantumNous/new-api/relaykit/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequestURLAlphaSearch(t *testing.T) {
	adaptor := &Adaptor{}
	info := &relaycommon.RelayInfo{
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelType:    constant.ChannelTypeSub2API,
			ChannelBaseUrl: "https://sub2api.example",
		},
		RequestURLPath: "/v1/alpha/search",
		RelayMode:      relayconstant.RelayModeAlphaSearch,
	}

	url, err := adaptor.GetRequestURL(info)
	require.NoError(t, err)
	assert.Equal(t, "https://sub2api.example/v1/alpha/search", url)
}

func TestGetRequestURLRoutesAntigravityModels(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		model       string
		relayFormat types.RelayFormat
		requestPath string
		want        string
	}{
		{
			name:        "Claude Opus 5.5",
			baseURL:     "https://sub2api.example",
			model:       "claude-opus-5-5-high",
			relayFormat: types.RelayFormatClaude,
			requestPath: "/v1/messages",
			want:        "https://sub2api.example/antigravity/v1/messages",
		},
		{
			name:        "Claude Sonnet 5.5 with configured prefix",
			baseURL:     "https://sub2api.example/antigravity/",
			model:       "claude-sonnet-5-5-medium",
			relayFormat: types.RelayFormatClaude,
			requestPath: "/v1/messages",
			want:        "https://sub2api.example/antigravity/v1/messages",
		},
		{
			name:        "Gemini 4 Argon",
			baseURL:     "https://sub2api.example",
			model:       "gemini-4-argon-high",
			relayFormat: types.RelayFormatGemini,
			requestPath: "/v1beta/models/gemini-4-argon-high:generateContent",
			want:        "https://sub2api.example/antigravity/v1beta/models/gemini-4-argon-high:generateContent",
		},
		{
			name:        "ordinary Claude stays on the gateway root",
			baseURL:     "https://sub2api.example",
			model:       "claude-sonnet-4-5",
			relayFormat: types.RelayFormatClaude,
			requestPath: "/v1/messages",
			want:        "https://sub2api.example/v1/messages",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &relaycommon.RelayInfo{
				RelayFormat:    tt.relayFormat,
				RequestURLPath: tt.requestPath,
				ChannelMeta: &relaycommon.ChannelMeta{
					ChannelType:       constant.ChannelTypeSub2API,
					ChannelBaseUrl:    tt.baseURL,
					UpstreamModelName: tt.model,
				},
			}

			url, err := (&Adaptor{}).GetRequestURL(info)

			require.NoError(t, err)
			assert.Equal(t, tt.want, url)
			assert.Equal(t, tt.baseURL, info.ChannelBaseUrl)
		})
	}
}

func TestAdaptorInheritsNewAPIResponsesCompactSupport(t *testing.T) {
	adaptor := &Adaptor{}
	info := &relaycommon.RelayInfo{
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelType:    constant.ChannelTypeSub2API,
			ChannelBaseUrl: "https://sub2api.example",
		},
		RequestURLPath: "/v1/responses/compact",
		RelayMode:      relayconstant.RelayModeResponsesCompact,
	}

	url, err := adaptor.GetRequestURL(info)

	require.NoError(t, err)
	assert.Equal(t, "https://sub2api.example/v1/responses/compact", url)
	assert.Equal(t, "sub2api", adaptor.GetChannelName())
	assert.Empty(t, adaptor.GetModelList())
}

func TestConvertClaudeRequestPreservesAdaptiveThinkingForCompatibleModel(t *testing.T) {
	adaptor := &Adaptor{}
	maxTokens := uint(8192)
	temperature := 0.2
	topP := 0.99
	request := &dto.ClaudeRequest{
		Model:        "gpt-5.6-sol",
		MaxTokens:    &maxTokens,
		Temperature:  &temperature,
		TopP:         &topP,
		Thinking:     &dto.Thinking{Type: "adaptive", Display: "summarized"},
		OutputConfig: json.RawMessage(`{"effort":"xhigh","provider_option":true}`),
		Messages: []dto.ClaudeMessage{
			{Role: "user", Content: "hello"},
		},
	}
	info := &relaycommon.RelayInfo{
		OriginModelName: "gpt-5.6-sol",
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelType: constant.ChannelTypeSub2API,
		},
	}

	converted, err := adaptor.ConvertClaudeRequest(nil, info, request)

	require.NoError(t, err)
	assert.Same(t, request, converted)
	require.NotNil(t, request.Thinking)
	assert.Equal(t, "adaptive", request.Thinking.Type)
	assert.Equal(t, "summarized", request.Thinking.Display)
	assert.JSONEq(t, `{"effort":"xhigh","provider_option":true}`, string(request.OutputConfig))
	assert.Same(t, &temperature, request.Temperature)
	assert.Same(t, &topP, request.TopP)
	assert.Equal(t, "xhigh", info.ReasoningEffort)
	assert.Equal(t, "gpt-5.6-sol", info.UpstreamModelName)
}
