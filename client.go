package chainhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client represents a Chainhooks API client.
type Client struct {
	baseURL     string
	apiKey      *string
	jwt         *string
	httpClient  *http.Client
	userAgent   string
	timeout     time.Duration
	headers     map[string]string
}

// ClientConfig represents the configuration for creating a new client.
type ClientConfig struct {
	BaseURL   string
	APIKey    *string
	JWT       *string
	HTTPClient *http.Client
	Timeout   time.Duration
	UserAgent string
}

// NewClient creates a new Chainhooks API client.
//
// The baseURL is required and should point to the Chainhooks API endpoint.
// You can use ChainhooksBaseURLs[NetworkMainnet] or ChainhooksBaseURLs[NetworkTestnet]
// for the standard Hiro-hosted endpoints.
func NewClient(baseURL string) *Client {
	return NewClientWithConfig(&ClientConfig{
		BaseURL: baseURL,
	})
}

// NewClientWithConfig creates a new client with custom configuration.
func NewClientWithConfig(cfg *ClientConfig) *Client {
	if cfg == nil {
		cfg = &ClientConfig{}
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = ChainhooksBaseURLs[NetworkMainnet]
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	if cfg.Timeout > 0 {
		cfg.HTTPClient.Timeout = cfg.Timeout
	}

	if cfg.UserAgent == "" {
		cfg.UserAgent = "chainhooks-client-go/1.0.0"
	}

	// Ensure baseURL doesn't have trailing slash
	cfg.BaseURL = strings.TrimSuffix(cfg.BaseURL, "/")

	client := &Client{
		baseURL:    cfg.BaseURL,
		apiKey:     cfg.APIKey,
		jwt:        cfg.JWT,
		httpClient: cfg.HTTPClient,
		userAgent:  cfg.UserAgent,
		timeout:    cfg.Timeout,
		headers:    make(map[string]string),
	}

	// Set default headers
	client.headers[HeaderAccept] = ContentTypeJSON
	client.headers[HeaderContentType] = ContentTypeJSON

	return client
}

// SetAPIKey sets the API key for authentication.
func (c *Client) SetAPIKey(apiKey string) {
	c.apiKey = &apiKey
}

// SetJWT sets the JWT for authentication.
func (c *Client) SetJWT(jwt string) {
	c.jwt = &jwt
}

// SetHeader sets a custom header.
func (c *Client) SetHeader(key, value string) {
	c.headers[key] = value
}

// request performs an HTTP request to the Chainhooks API.
func (c *Client) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	fullURL := fmt.Sprintf("%s%s", c.baseURL, path)

	// Encode request body
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Set authentication headers
	if c.jwt != nil {
		req.Header.Set(HeaderAuthorization, fmt.Sprintf("Bearer %s", *c.jwt))
	}
	if c.apiKey != nil {
		req.Header.Set(HeaderAPIKey, *c.apiKey)
	}

	req.Header.Set("User-Agent", c.userAgent)

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}

	// Handle response
	if resp.StatusCode >= 400 {
		return newHttpError(resp, req)
	}

	// For 204 No Content, don't try to unmarshal
	if resp.StatusCode == http.StatusNoContent {
		resp.Body.Close()
		return nil
	}

	// Read and unmarshal response body
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response body: %w", err)
		}
	}

	return nil
}

// ============================================================================
// Chainhook Management Methods
// ============================================================================

// RegisterChainhook registers a new chainhook.
func (c *Client) RegisterChainhook(ctx context.Context, definition *ChainhookDefinition) (*Chainhook, error) {
	if definition == nil {
		return nil, &ValidationError{
			Field:  "definition",
			Reason: "definition cannot be nil",
		}
	}

	var result Chainhook
	err := c.request(ctx, MethodPOST, EndpointChainhooks, definition, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateChainhook updates an existing chainhook.
func (c *Client) UpdateChainhook(ctx context.Context, uuid UUID, definition *ChainhookDefinition) (*Chainhook, error) {
	if uuid == "" {
		return nil, &ValidationError{
			Field:  "uuid",
			Reason: "uuid cannot be empty",
		}
	}

	if definition == nil {
		return nil, &ValidationError{
			Field:  "definition",
			Reason: "definition cannot be nil",
		}
	}

	path := fmt.Sprintf(EndpointChainhook, uuid)
	var result Chainhook
	err := c.request(ctx, MethodPATCH, path, definition, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetChainhooks retrieves all chainhooks with pagination support.
func (c *Client) GetChainhooks(ctx context.Context, opts *PaginationOptions) (*PaginatedChainhookResponse, error) {
	path := EndpointChainhooks

	// Add query parameters
	if opts != nil {
		params := url.Values{}
		params.Set("offset", fmt.Sprintf("%d", opts.Offset))
		params.Set("limit", fmt.Sprintf("%d", opts.Limit))
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}

	var result PaginatedChainhookResponse
	err := c.request(ctx, MethodGET, path, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetChainhook retrieves a specific chainhook by UUID.
func (c *Client) GetChainhook(ctx context.Context, uuid UUID) (*Chainhook, error) {
	if uuid == "" {
		return nil, &ValidationError{
			Field:  "uuid",
			Reason: "uuid cannot be empty",
		}
	}

	path := fmt.Sprintf(EndpointChainhook, uuid)
	var result Chainhook
	err := c.request(ctx, MethodGET, path, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// EnableChainhook enables or disables a chainhook.
func (c *Client) EnableChainhook(ctx context.Context, uuid UUID, enabled bool) error {
	if uuid == "" {
		return &ValidationError{
			Field:  "uuid",
			Reason: "uuid cannot be empty",
		}
	}

	path := fmt.Sprintf(EndpointChainhookEnabled, uuid)
	body := map[string]bool{"enabled": enabled}

	return c.request(ctx, MethodPATCH, path, body, nil)
}

// BulkEnableChainhooks enables or disables multiple chainhooks based on filters.
func (c *Client) BulkEnableChainhooks(ctx context.Context, request *BulkEnableChainhooksRequest) (*BulkEnableChainhooksResponse, error) {
	if request == nil {
		return nil, &ValidationError{
			Field:  "request",
			Reason: "request cannot be nil",
		}
	}

	// Validate that at least one filter is provided
	if len(request.UUIDs) == 0 && request.WebhookURL == nil && len(request.Statuses) == 0 {
		return nil, &ValidationError{
			Field:  "request",
			Reason: "at least one filter (uuids, webhook_url, or statuses) must be provided",
		}
	}

	var result BulkEnableChainhooksResponse
	err := c.request(ctx, MethodPATCH, EndpointBulkEnabled, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteChainhook deletes a chainhook.
func (c *Client) DeleteChainhook(ctx context.Context, uuid UUID) error {
	if uuid == "" {
		return &ValidationError{
			Field:  "uuid",
			Reason: "uuid cannot be empty",
		}
	}

	path := fmt.Sprintf(EndpointChainhook, uuid)
	return c.request(ctx, MethodDELETE, path, nil, nil)
}

// ============================================================================
// Consumer Secret Methods
// ============================================================================

// RotateConsumerSecret generates or rotates the consumer secret.
func (c *Client) RotateConsumerSecret(ctx context.Context) (*ConsumerSecretResponse, error) {
	var result ConsumerSecretResponse
	err := c.request(ctx, MethodPOST, EndpointConsumerSecret, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetConsumerSecret retrieves the current consumer secret.
func (c *Client) GetConsumerSecret(ctx context.Context) (*ConsumerSecretResponse, error) {
	var result ConsumerSecretResponse
	err := c.request(ctx, MethodGET, EndpointConsumerSecret, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteConsumerSecret deletes the consumer secret.
func (c *Client) DeleteConsumerSecret(ctx context.Context) error {
	return c.request(ctx, MethodDELETE, EndpointConsumerSecret, nil, nil)
}

// ============================================================================
// Evaluation Methods
// ============================================================================

// EvaluateChainhook triggers an on-demand evaluation of a chainhook.
func (c *Client) EvaluateChainhook(ctx context.Context, uuid UUID, blockHeight uint64) error {
	if uuid == "" {
		return &ValidationError{
			Field:  "uuid",
			Reason: "uuid cannot be empty",
		}
	}

	path := fmt.Sprintf(EndpointEvaluate, uuid)
	body := &EvaluateChainhookRequest{
		BlockHeight: blockHeight,
	}

	return c.request(ctx, MethodPOST, path, body, nil)
}

// ============================================================================
// Status Methods
// ============================================================================

// GetStatus retrieves the API status.
func (c *Client) GetStatus(ctx context.Context) (*ApiStatusResponse, error) {
	var result ApiStatusResponse
	err := c.request(ctx, MethodGET, EndpointStatus, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
