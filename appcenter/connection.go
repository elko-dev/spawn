package appcenter

const (
	baseURL = "https://api.appcenter.ms/v0.1/"
)

// Connection struct
type Connection struct {
	AuthorizationToken string
	BaseUrl            string
}

// GetClientByAPIURL function to create client
func (connection *Connection) GetClientByAPIURL(pathURL string) *Client {
	url := connection.BaseUrl + pathURL
	return NewClient(connection, url)
}

// NewConnection Creates a new App Center connection instance using a personal access token.
func NewConnection(personalAccessToken string) *Connection {
	return &Connection{
		AuthorizationToken: personalAccessToken,
		BaseUrl:            baseURL,
	}
}
