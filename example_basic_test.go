package chainhooks

import (
	"context"
	"testing"
)

// ExampleNewClient demonstrates basic client creation.
func ExampleNewClient() {
	client := NewClient(ChainhooksBaseURLs[NetworkMainnet])
	client.SetAPIKey("your-api-key")
	_ = client
	// Output:
}

// ExampleNewChainhookBuilder demonstrates building a chainhook definition.
func ExampleNewChainhookBuilder() {
	definition, err := NewChainhookBuilder("stx-transfer-hook", NetworkMainnet).
		WithWebhookURL("https://example.com/webhook").
		AddSTXTransfer(
			PrincipalStandard("SP..."),
			nil,
			nil,
		).
		WithEnableOnRegistration(true).
		Build()

	if err != nil {
		panic(err)
	}

	_ = definition
	// Output:
}

// ExampleClient_RegisterChainhook demonstrates registering a chainhook.
func ExampleClient_RegisterChainhook(t *testing.T) {
	client := NewClient(ChainhooksBaseURLs[NetworkMainnet])
	client.SetAPIKey("test-api-key")

	definition := &ChainhookDefinition{
		Name:    "test-hook",
		Version: "1",
		Chain:   ChainStacks,
		Network: NetworkMainnet,
		Filters: ChainhookFilters{
			Events: []interface{}{
				&STXTransferFilter{
					Type: EventTypeSTXTransfer,
				},
			},
		},
		Action: ChainhookAction{
			Type: "http_post",
			URL:  "https://example.com/webhook",
		},
	}

	// This would normally make an HTTP request
	_, _ = client.RegisterChainhook(context.Background(), definition)
	// Output:
}

// ExampleBulkEnableUUIDs demonstrates creating a bulk enable request.
func ExampleBulkEnableUUIDs() {
	request := BulkEnableUUIDs(true, UUID("uuid-1"), UUID("uuid-2"))
	_ = request
	// Output:
}

// ExampleBulkEnableByWebhook demonstrates bulk enable by webhook URL.
func ExampleBulkEnableByWebhook() {
	request := BulkEnableByWebhook(false, "https://example.com/webhook")
	_ = request
	// Output:
}

// ExampleBulkEnableByStatus demonstrates bulk enable by status.
func ExampleBulkEnableByStatus() {
	request := BulkEnableByStatus(true, ChainhookStatusStreaming)
	_ = request
	// Output:
}

// ExampleIsHttpError demonstrates error type checking.
func ExampleIsHttpError() {
	var err error

	if IsHttpError(err) {
		httpErr, _ := AsHttpError(err)
		_ = httpErr.StatusCode
	}
	// Output:
}

// ExampleNewPaginationOptions demonstrates creating pagination options.
func ExampleNewPaginationOptions() {
	opts := NewPaginationOptions(0, 10)
	_ = opts
	// Output:
}
