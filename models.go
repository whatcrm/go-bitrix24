package b24

import "github.com/whatcrm/go-bitrix24/models"

// Handler with main data to work with amoCRM

type API struct {
	ClientID     string
	ClientSecret string
	Domain       string
	Auth         string
	Debug        bool
}

// Struct to send a data to amoCRM in order to get the tokens

// The tokens' struct

type UpdatedTokens struct {
	AccessToken    string `json:"access_token"`
	ClientEndpoint string `json:"client_endpoint"`
	Domain         string `json:"domain"`
	ExpiresIn      int    `json:"expires_in"`
	MemberID       string `json:"member_id"`
	RefreshToken   string `json:"refresh_token"`
	Scope          string `json:"scope"`
	ServerEndpoint string `json:"server_endpoint"`
	Status         string `json:"status"`
}

type RequestParams struct {
	// Update Token
	RefreshToken string `json:"refresh_token,omitempty"`

	// Entity ID
	ID string `json:"ID,omitempty"`

	// Regular Params
	Title       string `json:"TITLE,omitempty"`
	Description string `json:"DESCRIPTION,omitempty"`
	Placement   string `json:"PLACEMENT,omitempty"`
	Handler     string `json:"HANDLER,omitempty"`

	// Event
	Event         string `json:"EVENT,omitempty"`
	AuthType      int    `json:"auth_type,omitempty"`
	EventType     string `json:"event_type,omitempty"` // offline online
	AuthConnector string `json:"auth_connector,omitempty"`
	// TODO Options

	// Robot params + handler
	Code             string      `json:"CODE,omitempty"`
	AuthUserID       int         `json:"AUTH_USER_ID,omitempty"`
	Name             string      `json:"NAME,omitempty"`
	Properties       *Properties `json:"PROPERTIES,omitempty"`
	EventToken       string      `json:"EVENT_TOKEN,omitempty"`
	UseSubscription  string      `json:"USE_SUBSCRIPTION,omitempty"`
	ReturnProperties *Properties `json:"RETURN_PROPERTIES,omitempty"`
	ReturnValues     *Properties `json:"return_values,omitempty"`
}

type Properties struct {
	Select *Select `json:"select,omitempty"`
	String *String `json:"string,omitempty"`
}

type String struct {
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	Default     string `json:"Default,omitempty"`
	Description string `json:"Description,omitempty"`
	Required    string `json:"Required,omitempty"`
	Multiple    string `json:"Multiple,omitempty"`
}

type Bool struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default string `json:"default,omitempty"`
}

type Select struct {
	Default string      `json:"Default,omitempty"`
	Name    string      `json:"Name"`
	Type    string      `json:"Type"`
	Options interface{} `json:"Options,omitempty"`
}

type Options struct {
	Work  string `json:"work"`
	Home  string `json:"home"`
	Mob   string `json:"mob"`
	Other string `json:"other"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type MainResult struct {
	Result bool        `json:"result"`
	Time   models.Time `json:"time"`
}

type UnBind struct {
	Result struct {
		Count int `json:"count"`
	} `json:"result"`
	Time models.Time `json:"time"`
}
