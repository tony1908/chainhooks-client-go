package chainhooks

// ChainhooksBaseURLs contains the base URLs for different networks.
var ChainhooksBaseURLs = map[Network]string{
	NetworkMainnet: "https://api.mainnet.hiro.so",
	NetworkTestnet: "https://api.testnet.hiro.so",
}

// Default constants
const (
	DefaultAPIVersion = "1"
	DefaultChain      = ChainStacks
)

// HTTP methods
const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodPATCH  = "PATCH"
	MethodDELETE = "DELETE"
)

// API endpoints
const (
	EndpointChainhooks       = "/chainhooks/me"
	EndpointChainhook        = "/chainhooks/me/%s"
	EndpointChainhookEnabled = "/chainhooks/me/%s/enabled"
	EndpointBulkEnabled      = "/chainhooks/me/enabled"
	EndpointEvaluate         = "/chainhooks/me/%s/evaluate"
	EndpointConsumerSecret   = "/chainhooks/me/secret"
	EndpointStatus           = "/chainhooks"
)

// Header names
const (
	HeaderAccept        = "Accept"
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderAPIKey        = "x-api-key"
)

// Header values
const (
	ContentTypeJSON = "application/json"
)
