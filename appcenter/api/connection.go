package api

const (
	baseURL = "https://api.appcenter.ms/v0.1/"
)

// Connection struct
type Connection struct {
	AuthorizationToken string
	BaseUrl            string
}

// GetClient function to create client
func (connection *Connection) GetClient() *Client {
	url := connection.BaseUrl
	return NewClient(connection, url)
}

// NewConnection Creates a new App Center connection instance using a personal access token.
func NewConnection(personalAccessToken string) *Connection {
	return &Connection{
		AuthorizationToken: personalAccessToken,
		BaseUrl:            baseURL,
	}
}
