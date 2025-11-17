# Chainhooks Client - Go

A comprehensive Go client library for interacting with the [Hiro Chainhooks API](https://docs.hiro.so/chainhooks). Chainhooks is a blockchain monitoring and event triggering system for the Stacks blockchain that enables developers to register webhooks that detect and react to specific blockchain events.

## Features

- ✅ Full support for all Chainhooks API endpoints
- ✅ Complete type safety with Go structs
- ✅ Support for 16 different blockchain event types
- ✅ Flexible authentication (API Key and JWT)
- ✅ Pagination support for list operations
- ✅ Builder pattern for easy chainhook definition construction
- ✅ Comprehensive error handling with HTTP context
- ✅ Context support for cancellation and timeouts
- ✅ Minimal dependencies (uses only Go standard library)

## Installation

```bash
go get github.com/hirosystems/chainhooks-client-go
```

## Quick Start

### Basic Usage

```go
package main

import (
	"context"
	"log"

	"github.com/hirosystems/chainhooks-client-go"
)

func main() {
	// Create a client
	client := chainhooks.NewClient(chainhooks.ChainhooksBaseURLs[chainhooks.NetworkMainnet])
	client.SetAPIKey("your-api-key")

	// Create a simple chainhook definition
	definition := &chainhooks.ChainhookDefinition{
		Name:    "my-stx-hook",
		Version: "1",
		Chain:   chainhooks.ChainStacks,
		Network: chainhooks.NetworkMainnet,
		Filters: chainhooks.ChainhookFilters{
			Events: []interface{}{
				&chainhooks.STXTransferFilter{
					Type: chainhooks.EventTypeSTXTransfer,
				},
			},
		},
		Action: chainhooks.ChainhookAction{
			Type: "http_post",
			URL:  "https://example.com/webhook",
		},
	}

	// Register the chainhook
	ctx := context.Background()
	hook, err := client.RegisterChainhook(ctx, definition)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Chainhook registered: %s", hook.UUID)
}
```

### Using the Builder Pattern

```go
// Create a chainhook using the fluent builder API
definition, err := chainhooks.NewChainhookBuilder("stx-transfer-hook", chainhooks.NetworkMainnet).
	WithWebhookURL("https://example.com/webhook").
	AddSTXTransfer(
		chainhooks.PrincipalStandard("SP..."),
		chainhooks.PrincipalStandard("SP..."),
		chainhooks.StringPtr("1000000"),
	).
	WithEnableOnRegistration(true).
	Build()

if err != nil {
	log.Fatal(err)
}

hook, err := client.RegisterChainhook(context.Background(), definition)
if err != nil {
	log.Fatal(err)
}
```

## Client Configuration

### Basic Client

```go
client := chainhooks.NewClient(
	chainhooks.ChainhooksBaseURLs[chainhooks.NetworkMainnet],
)
```

### Client with Configuration

```go
client := chainhooks.NewClientWithConfig(&chainhooks.ClientConfig{
	BaseURL:   "https://api.mainnet.hiro.so",
	APIKey:    chainhooks.StringPtr("your-api-key"),
	JWT:       chainhooks.StringPtr("your-jwt-token"),
	Timeout:   30 * time.Second,
	UserAgent: "my-app/1.0.0",
})
```

### Setting Authentication

```go
client.SetAPIKey("your-api-key")
// OR
client.SetJWT("your-jwt-token")
// OR both
client.SetAPIKey("your-api-key")
client.SetJWT("your-jwt-token")
```

## API Methods

### Chainhook Management

#### Register a Chainhook

```go
definition := &chainhooks.ChainhookDefinition{
	Name:    "my-hook",
	Version: "1",
	Chain:   chainhooks.ChainStacks,
	Network: chainhooks.NetworkMainnet,
	Filters: chainhooks.ChainhookFilters{
		Events: []interface{}{
			&chainhooks.STXTransferFilter{
				Type: chainhooks.EventTypeSTXTransfer,
			},
		},
	},
	Action: chainhooks.ChainhookAction{
		Type: "http_post",
		URL:  "https://example.com/webhook",
	},
}

hook, err := client.RegisterChainhook(context.Background(), definition)
```

#### Get Chainhook

```go
hook, err := client.GetChainhook(context.Background(), "uuid-string")
```

#### List Chainhooks (with pagination)

```go
opts := &chainhooks.PaginationOptions{
	Offset: 0,
	Limit:  10,
}

response, err := client.GetChainhooks(context.Background(), opts)
if err != nil {
	log.Fatal(err)
}

for _, hook := range response.Chainhooks {
	log.Printf("Hook: %s (Status: %s)", hook.UUID, hook.Status)
}
```

#### Update Chainhook

```go
updated := &chainhooks.ChainhookDefinition{
	// ... updated definition
}

hook, err := client.UpdateChainhook(context.Background(), "uuid-string", updated)
```

#### Enable/Disable Chainhook

```go
// Enable
err := client.EnableChainhook(context.Background(), "uuid-string", true)

// Disable
err := client.EnableChainhook(context.Background(), "uuid-string", false)
```

#### Bulk Enable/Disable

```go
// By UUIDs
request := chainhooks.BulkEnableUUIDs(
	true,
	chainhooks.UUID("uuid-1"),
	chainhooks.UUID("uuid-2"),
)

response, err := client.BulkEnableChainhooks(context.Background(), request)

// By Webhook URL
request := chainhooks.BulkEnableByWebhook(
	false,
	"https://example.com/webhook",
)

// By Status
request := chainhooks.BulkEnableByStatus(
	true,
	chainhooks.ChainhookStatusStreaming,
)
```

#### Delete Chainhook

```go
err := client.DeleteChainhook(context.Background(), "uuid-string")
```

### Consumer Secrets

#### Get Consumer Secret

```go
secret, err := client.GetConsumerSecret(context.Background())
log.Println(secret.Secret)
```

#### Rotate Consumer Secret

```go
secret, err := client.RotateConsumerSecret(context.Background())
log.Println(secret.Secret)
```

#### Delete Consumer Secret

```go
err := client.DeleteConsumerSecret(context.Background())
```

### Evaluation

#### Evaluate Chainhook

```go
err := client.EvaluateChainhook(context.Background(), "uuid-string", 100000)
```

### API Status

#### Get Status

```go
status, err := client.GetStatus(context.Background())
log.Printf("Status: %s, Version: %s", status.Status, status.Version)
```

## Event Types

The client supports 16 different blockchain event types:

### Token Events

```go
// Fungible Token Transfer
&chainhooks.FTTransferFilter{
	Type:      chainhooks.EventTypeFTTransfer,
	Asset:     "USDA",
	Sender:    chainhooks.PrincipalStandard("SP..."),
	Recipient: chainhooks.PrincipalStandard("SP..."),
	Amount:    chainhooks.StringPtr("1000000"),
}

// Fungible Token Mint
&chainhooks.FTMintFilter{
	Type:      chainhooks.EventTypeFTMint,
	Asset:     "USDA",
	Recipient: chainhooks.PrincipalStandard("SP..."),
	Amount:    chainhooks.StringPtr("1000000"),
}

// Fungible Token Burn
&chainhooks.FTBurnFilter{
	Type:   chainhooks.EventTypeFTBurn,
	Asset:  "USDA",
	Sender: chainhooks.PrincipalStandard("SP..."),
	Amount: chainhooks.StringPtr("1000000"),
}
```

### NFT Events

```go
// NFT Transfer
&chainhooks.NFTTransferFilter{
	Type:      chainhooks.EventTypeNFTTransfer,
	Asset:     "nft-collection",
	Sender:    chainhooks.PrincipalStandard("SP..."),
	Recipient: chainhooks.PrincipalStandard("SP..."),
}

// NFT Mint
&chainhooks.NFTMintFilter{
	Type:      chainhooks.EventTypeNFTMint,
	Asset:     "nft-collection",
	Recipient: chainhooks.PrincipalStandard("SP..."),
}

// NFT Burn
&chainhooks.NFTBurnFilter{
	Type:   chainhooks.EventTypeNFTBurn,
	Asset:  "nft-collection",
	Sender: chainhooks.PrincipalStandard("SP..."),
}
```

### STX (Native Token) Events

```go
// STX Transfer
&chainhooks.STXTransferFilter{
	Type:      chainhooks.EventTypeSTXTransfer,
	Sender:    chainhooks.PrincipalStandard("SP..."),
	Recipient: chainhooks.PrincipalStandard("SP..."),
	Amount:    chainhooks.StringPtr("1000000"),
}

// STX Mint
&chainhooks.STXMintFilter{
	Type:      chainhooks.EventTypeSTXMint,
	Recipient: chainhooks.PrincipalStandard("SP..."),
	Amount:    chainhooks.StringPtr("1000000"),
}

// STX Burn
&chainhooks.STXBurnFilter{
	Type:   chainhooks.EventTypeSTXBurn,
	Sender: chainhooks.PrincipalStandard("SP..."),
	Amount: chainhooks.StringPtr("1000000"),
}
```

### Smart Contract Events

```go
// Contract Deployment
&chainhooks.ContractDeployFilter{
	Type:              chainhooks.EventTypeContractDeploy,
	DeployerPrincipal: chainhooks.PrincipalStandard("SP..."),
}

// Contract Call
&chainhooks.ContractCallFilter{
	Type:               chainhooks.EventTypeContractCall,
	ContractIdentifier: chainhooks.StringPtr("SP...contract_name"),
	Method:             chainhooks.StringPtr("method_name"),
	Sender:             chainhooks.PrincipalStandard("SP..."),
}

// Contract Log (emit events)
&chainhooks.ContractLogFilter{
	Type:               chainhooks.EventTypeContractLog,
	ContractIdentifier: chainhooks.StringPtr("SP...contract_name"),
}
```

### System Events

```go
// Balance Change
&chainhooks.BalanceChangeFilter{
	Type:      chainhooks.EventTypeBalanceChange,
	Principal: chainhooks.PrincipalStandard("SP..."),
}

// Coinbase (mining rewards)
&chainhooks.CoinbaseFilter{
	Type:      chainhooks.EventTypeCoinbase,
	Recipient: chainhooks.PrincipalStandard("SP..."),
}

// Tenure Change
&chainhooks.TenureChangeFilter{
	Type: chainhooks.EventTypeTenureChange,
}
```

## Builder Pattern Examples

### Complex Chainhook with Multiple Filters

```go
definition, err := chainhooks.NewChainhookBuilder(
	"multi-filter-hook",
	chainhooks.NetworkMainnet,
).
	WithWebhookURL("https://example.com/webhook").
	AddSTXTransfer(
		chainhooks.PrincipalStandard("SP..."),
		nil,
		chainhooks.StringPtr("1000000"),
	).
	AddFTTransfer(
		"USDA",
		chainhooks.PrincipalStandard("SP..."),
		nil,
		nil,
	).
	AddContractCall(
		chainhooks.StringPtr("SP...my_contract"),
		chainhooks.StringPtr("swap"),
		nil,
	).
	WithEnableOnRegistration(true).
	WithExpireAfterOccurrences(100).
	WithIncludeContractABI(true).
	WithDecodeClarityValues(true).
	Build()

if err != nil {
	log.Fatal(err)
}
```

### Using ChainhookOptionsBuilder

```go
options := chainhooks.NewChainhookOptionsBuilder().
	EnableOnRegistration(true).
	ExpireAfterEvaluations(10000).
	DecodeClarityValues(true).
	IncludeContractABI(true).
	Build()

definition := &chainhooks.ChainhookDefinition{
	Name:    "my-hook",
	Version: "1",
	Chain:   chainhooks.ChainStacks,
	Network: chainhooks.NetworkMainnet,
	Filters: chainhooks.ChainhookFilters{
		Events: []interface{}{
			&chainhooks.STXTransferFilter{
				Type: chainhooks.EventTypeSTXTransfer,
			},
		},
	},
	Action: chainhooks.ChainhookAction{
		Type: "http_post",
		URL:  "https://example.com/webhook",
	},
	Options: options,
}
```

## Error Handling

The client provides robust error handling with helpful utilities:

```go
hook, err := client.GetChainhook(ctx, "uuid")
if err != nil {
	// Check if it's an HTTP error
	if chainhooks.IsHttpError(err) {
		httpErr, _ := chainhooks.AsHttpError(err)
		log.Printf("HTTP Error: %d %s", httpErr.StatusCode, httpErr.Body)
	}

	// Check specific error conditions
	if chainhooks.IsNotFound(err) {
		log.Println("Chainhook not found")
	} else if chainhooks.IsUnauthorized(err) {
		log.Println("Unauthorized - check API key")
	} else if chainhooks.IsServerError(err) {
		log.Println("Server error - retry later")
	}
}
```

### Error Types

- `HttpError` - HTTP request/response errors with full context
- `ValidationError` - Validation errors when building requests
- `ConfigError` - Configuration errors

## Helper Functions

The client includes several utility functions for common tasks:

```go
// Create pointers to basic types
name := chainhooks.StringPtr("my-hook")
enabled := chainhooks.BoolPtr(true)
count := chainhooks.Uint64Ptr(100)

// Create principals
standardAddr := chainhooks.PrincipalStandard("SP...")
contractAddr := chainhooks.PrincipalContract("SP...contract_name")

// Create pagination options
opts := chainhooks.NewPaginationOptions(0, 10)

// Bulk operation helpers
bulkByUUIDs := chainhooks.BulkEnableUUIDs(true, uuid1, uuid2)
bulkByWebhook := chainhooks.BulkEnableByWebhook(false, "https://example.com/webhook")
bulkByStatus := chainhooks.BulkEnableByStatus(true, chainhooks.ChainhookStatusStreaming)

// Error checking helpers
isNotFound := chainhooks.IsNotFound(err)
isUnauth := chainhooks.IsUnauthorized(err)
isServer := chainhooks.IsServerError(err)
isClient := chainhooks.IsClientError(err)
```

## Type Reference

### Networks

```go
chainhooks.NetworkMainnet // "mainnet"
chainhooks.NetworkTestnet  // "testnet"
```

### Chainhook Status

```go
chainhooks.ChainhookStatusNew
chainhooks.ChainhookStatusStreaming
chainhooks.ChainhookStatusExpired
chainhooks.ChainhookStatusInterrupted
```

### Event Types

See the [Event Types](#event-types) section above for all 16 supported event types.

## Context Support

All API methods support context for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

hook, err := client.GetChainhook(ctx, "uuid")
```

## Thread Safety

The client is safe for concurrent use across multiple goroutines, as each request creates its own HTTP request objects.

## Base URLs

```go
chainhooks.ChainhooksBaseURLs[chainhooks.NetworkMainnet]
// Returns: "https://api.mainnet.hiro.so"

chainhooks.ChainhooksBaseURLs[chainhooks.NetworkTestnet]
// Returns: "https://api.testnet.hiro.so"
```

## Examples

### Complete Example: Create and Monitor a Chainhook

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/hirosystems/chainhooks-client-go"
)

func main() {
	// Create client
	client := chainhooks.NewClientWithConfig(&chainhooks.ClientConfig{
		BaseURL: chainhooks.ChainhooksBaseURLs[chainhooks.NetworkTestnet],
	})
	client.SetAPIKey("your-api-key")

	ctx := context.Background()

	// Define chainhook
	definition, err := chainhooks.NewChainhookBuilder(
		"testnet-stx-monitor",
		chainhooks.NetworkTestnet,
	).
		WithWebhookURL("https://example.com/webhook").
		AddSTXTransfer(nil, nil, nil).
		WithEnableOnRegistration(true).
		WithDecodeClarityValues(true).
		Build()

	if err != nil {
		log.Fatal(err)
	}

	// Register chainhook
	hook, err := client.RegisterChainhook(ctx, definition)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Registered chainhook: %s", hook.UUID)
	log.Printf("Status: %s, Enabled: %v", hook.Status, hook.Enabled)

	// List chainhooks
	response, err := client.GetChainhooks(ctx, &chainhooks.PaginationOptions{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Total chainhooks: %d", response.Total)

	// Get specific chainhook
	time.Sleep(1 * time.Second)
	hook, err = client.GetChainhook(ctx, hook.UUID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Latest status: %s", hook.Status)

	// Disable chainhook
	err = client.EnableChainhook(ctx, hook.UUID, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Chainhook disabled")

	// Clean up
	err = client.DeleteChainhook(ctx, hook.UUID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Chainhook deleted")
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

Apache-2.0

## Resources

- [Chainhooks Documentation](https://docs.hiro.so/chainhooks)
- [Hiro Systems](https://www.hiro.so/)
- [Stacks Blockchain](https://www.stacks.co/)
