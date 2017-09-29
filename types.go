package slacktest

import (
	"log"
	"net/http"
	"net/http/httptest"
	"sync"

	slack "github.com/nlopes/slack"
)

type contextKey string

// ServerURLContextKey is the context key to store the server's url
var ServerURLContextKey contextKey = "__SERVER_URL__"

// ServerWSContextKey is the context key to store the server's ws url
var ServerWSContextKey contextKey = "__SERVER_WS_URL__"

// ServerBotNameContextKey is the bot name
var ServerBotNameContextKey contextKey = "__SERVER_BOTNAME__"

// ServerBotIDContextKey is the bot userid
var ServerBotIDContextKey contextKey = "__SERVER_BOTID__"

// ServerBotChannelsContextKey is the list of channels associated with the fake server
var ServerBotChannelsContextKey contextKey = "__SERVER_CHANNELS__"

// ServerBotGroupsContextKey is the list of channels associated with the fake server
var ServerBotGroupsContextKey contextKey = "__SERVER_GROUPS__"

// ServerBotHubNameContextKey is the context key for passing along the server name registered in the hub
var ServerBotHubNameContextKey contextKey = "__SERVER_HUBNAME__"

var seenInboundMessages *messageCollection
var seenOutboundMessages *messageCollection
var masterHub = newHub()

type hub struct {
	sync.RWMutex
	serverChannels map[string]*messageChannels
}

type messageChannels struct {
	seen chan (string)
	sent chan (string)
}
type messageCollection struct {
	sync.RWMutex
	messages []string
}

type serverChannels struct {
	sync.RWMutex
	channels []slack.Channel
}

type serverGroups struct {
	sync.RWMutex
	channels []slack.Group
}

// Server represents a Slack Test server
type Server struct {
	server     *httptest.Server
	mux        *http.ServeMux
	Logger     *log.Logger
	BotName    string
	BotID      string
	ServerAddr string
	SeenFeed   chan (string)
	channels   *serverChannels
	groups     *serverGroups
}

type fullInfoSlackResponse struct {
	slack.Info
	slack.WebResponse
}
