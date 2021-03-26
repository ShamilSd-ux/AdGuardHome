package querylog

// Client is the information required by the query log to match against clients
// during searches.
type Client struct {
	Name  string            `json:"name"`
	Whois map[string]string `json:"whois,omitempty"`
	IDs   []string          `json:"ids"`
}

type clientCacheKey struct {
	clientID string
	ip       string
}

type clientCache map[clientCacheKey]*Client
