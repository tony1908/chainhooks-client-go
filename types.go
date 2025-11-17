package chainhooks

// ============================================================================
// Common Types
// ============================================================================

// Network represents the blockchain network.
type Network string

const (
	NetworkMainnet Network = "mainnet"
	NetworkTestnet Network = "testnet"
)

// Chain represents the blockchain.
type Chain string

const (
	ChainStacks Chain = "stacks"
)

// UUID is a unique identifier for a chainhook.
type UUID string

// Principal represents a Stacks principal (address).
type Principal struct {
	// Standard address format (e.g., "SP...")
	Standard *string `json:"standard,omitempty"`
	// Contract address format (e.g., "SP...contract_name")
	Contract *string `json:"contract,omitempty"`
}

// ============================================================================
// Chainhook Status Types
// ============================================================================

// ChainhookStatus represents the status of a chainhook.
type ChainhookStatus string

const (
	ChainhookStatusNew         ChainhookStatus = "new"
	ChainhookStatusStreaming   ChainhookStatus = "streaming"
	ChainhookStatusExpired     ChainhookStatus = "expired"
	ChainhookStatusInterrupted ChainhookStatus = "interrupted"
)

// ============================================================================
// Event Filter Types
// ============================================================================

// EventType represents different types of blockchain events.
type EventType string

const (
	// Token events
	EventTypeFTEvent       EventType = "ft_event"
	EventTypeFTMint        EventType = "ft_mint"
	EventTypeFTBurn        EventType = "ft_burn"
	EventTypeFTTransfer    EventType = "ft_transfer"
	EventTypeNFTEvent      EventType = "nft_event"
	EventTypeNFTMint       EventType = "nft_mint"
	EventTypeNFTBurn       EventType = "nft_burn"
	EventTypeNFTTransfer   EventType = "nft_transfer"
	EventTypeSTXEvent      EventType = "stx_event"
	EventTypeSTXMint       EventType = "stx_mint"
	EventTypeSTXBurn       EventType = "stx_burn"
	EventTypeSTXTransfer   EventType = "stx_transfer"
	// Contract events
	EventTypeContractDeploy EventType = "contract_deploy"
	EventTypeContractCall   EventType = "contract_call"
	EventTypeContractLog    EventType = "contract_log"
	// System events
	EventTypeBalanceChange EventType = "balance_change"
	EventTypeCoinbase      EventType = "coinbase"
	EventTypeTenureChange  EventType = "tenure_change"
)

// ============================================================================
// Event Filter Definitions
// ============================================================================

// EventFilter is an interface for all event filter types.
type EventFilter interface {
	eventFilterMarker()
}

// FTEventFilter represents a fungible token event filter.
type FTEventFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Action          string    `json:"action"`
	Sender          *Principal `json:"sender,omitempty"`
	Receiver        *Principal `json:"receiver,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *FTEventFilter) eventFilterMarker() {}

// FTMintFilter represents a fungible token mint event filter.
type FTMintFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Recipient       *Principal `json:"recipient,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *FTMintFilter) eventFilterMarker() {}

// FTBurnFilter represents a fungible token burn event filter.
type FTBurnFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Sender          *Principal `json:"sender,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *FTBurnFilter) eventFilterMarker() {}

// FTTransferFilter represents a fungible token transfer event filter.
type FTTransferFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Sender          *Principal `json:"sender,omitempty"`
	Recipient       *Principal `json:"recipient,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *FTTransferFilter) eventFilterMarker() {}

// NFTEventFilter represents an NFT event filter.
type NFTEventFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Action          string    `json:"action"`
	Sender          *Principal `json:"sender,omitempty"`
	Receiver        *Principal `json:"receiver,omitempty"`
}

func (f *NFTEventFilter) eventFilterMarker() {}

// NFTMintFilter represents an NFT mint event filter.
type NFTMintFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Recipient       *Principal `json:"recipient,omitempty"`
}

func (f *NFTMintFilter) eventFilterMarker() {}

// NFTBurnFilter represents an NFT burn event filter.
type NFTBurnFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Sender          *Principal `json:"sender,omitempty"`
}

func (f *NFTBurnFilter) eventFilterMarker() {}

// NFTTransferFilter represents an NFT transfer event filter.
type NFTTransferFilter struct {
	Type            EventType `json:"type"`
	Asset           string    `json:"asset"`
	Sender          *Principal `json:"sender,omitempty"`
	Recipient       *Principal `json:"recipient,omitempty"`
}

func (f *NFTTransferFilter) eventFilterMarker() {}

// STXEventFilter represents an STX (Stacks native token) event filter.
type STXEventFilter struct {
	Type            EventType `json:"type"`
	Action          string    `json:"action"`
	Sender          *Principal `json:"sender,omitempty"`
	Receiver        *Principal `json:"receiver,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *STXEventFilter) eventFilterMarker() {}

// STXMintFilter represents an STX mint event filter.
type STXMintFilter struct {
	Type            EventType `json:"type"`
	Recipient       *Principal `json:"recipient,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *STXMintFilter) eventFilterMarker() {}

// STXBurnFilter represents an STX burn event filter.
type STXBurnFilter struct {
	Type            EventType `json:"type"`
	Sender          *Principal `json:"sender,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *STXBurnFilter) eventFilterMarker() {}

// STXTransferFilter represents an STX transfer event filter.
type STXTransferFilter struct {
	Type            EventType `json:"type"`
	Sender          *Principal `json:"sender,omitempty"`
	Recipient       *Principal `json:"recipient,omitempty"`
	Amount          *string   `json:"amount,omitempty"`
}

func (f *STXTransferFilter) eventFilterMarker() {}

// ContractDeployFilter represents a contract deployment event filter.
type ContractDeployFilter struct {
	Type            EventType `json:"type"`
	DeployerPrincipal *Principal `json:"deployer_principal,omitempty"`
}

func (f *ContractDeployFilter) eventFilterMarker() {}

// ContractCallFilter represents a contract call event filter.
type ContractCallFilter struct {
	Type            EventType `json:"type"`
	ContractIdentifier *string `json:"contract_identifier,omitempty"`
	Method          *string   `json:"method,omitempty"`
	Sender          *Principal `json:"sender,omitempty"`
}

func (f *ContractCallFilter) eventFilterMarker() {}

// ContractLogFilter represents a contract log event filter.
type ContractLogFilter struct {
	Type            EventType `json:"type"`
	ContractIdentifier *string `json:"contract_identifier,omitempty"`
}

func (f *ContractLogFilter) eventFilterMarker() {}

// BalanceChangeFilter represents a balance change event filter.
type BalanceChangeFilter struct {
	Type            EventType `json:"type"`
	Principal       *Principal `json:"principal,omitempty"`
}

func (f *BalanceChangeFilter) eventFilterMarker() {}

// CoinbaseFilter represents a coinbase event filter.
type CoinbaseFilter struct {
	Type            EventType `json:"type"`
	Recipient       *Principal `json:"recipient,omitempty"`
}

func (f *CoinbaseFilter) eventFilterMarker() {}

// TenureChangeFilter represents a tenure change event filter.
type TenureChangeFilter struct {
	Type            EventType `json:"type"`
}

func (f *TenureChangeFilter) eventFilterMarker() {}

// ============================================================================
// Chainhook Definition
// ============================================================================

// ChainhookOptions represents optional configuration for a chainhook.
type ChainhookOptions struct {
	EnableOnRegistration       *bool   `json:"enable_on_registration,omitempty"`
	ExpireAfterEvaluations     *uint64 `json:"expire_after_evaluations,omitempty"`
	ExpireAfterOccurrences     *uint64 `json:"expire_after_occurrences,omitempty"`
	DecodeClarityValues        *bool   `json:"decode_clarity_values,omitempty"`
	IncludeContractABI         *bool   `json:"include_contract_abi,omitempty"`
	IncludeContractSourceCode  *bool   `json:"include_contract_source_code,omitempty"`
	IncludePostConditions      *bool   `json:"include_post_conditions,omitempty"`
	IncludeRawTransactions     *bool   `json:"include_raw_transactions,omitempty"`
	IncludeBlockSignatures     *bool   `json:"include_block_signatures,omitempty"`
	IncludeBlockMetadata       *bool   `json:"include_block_metadata,omitempty"`
}

// ChainhookFilters represents the event filters for a chainhook.
type ChainhookFilters struct {
	Events []interface{} `json:"events"`
}

// ChainhookAction represents the action to take when a chainhook is triggered.
type ChainhookAction struct {
	Type string `json:"type"` // "http_post"
	URL  string `json:"url"`
}

// ChainhookDefinition represents the definition of a chainhook.
type ChainhookDefinition struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"` // "1"
	Chain   Chain                  `json:"chain"`   // "stacks"
	Network Network                `json:"network"`
	Filters ChainhookFilters       `json:"filters"`
	Options *ChainhookOptions      `json:"options,omitempty"`
	Action  ChainhookAction        `json:"action"`
}

// ============================================================================
// Chainhook Response
// ============================================================================

// ChainhookStatusInfo represents the detailed status information of a chainhook.
type ChainhookStatusInfo struct {
	Status                    ChainhookStatus `json:"status"`
	Enabled                   bool            `json:"enabled"`
	CreatedAt                 int64           `json:"created_at"`
	LastEvaluatedAt           *int64          `json:"last_evaluated_at"`
	LastEvaluatedBlockHeight  *uint64         `json:"last_evaluated_block_height"`
	LastOccurrenceAt          *int64          `json:"last_occurrence_at"`
	LastOccurrenceBlockHeight *uint64         `json:"last_occurrence_block_height"`
	EvaluatedBlockCount       uint64          `json:"evaluated_block_count"`
	OccurrenceCount           uint64          `json:"occurrence_count"`
}

// Chainhook represents a registered chainhook with its metadata.
type Chainhook struct {
	UUID       UUID                 `json:"uuid"`
	Definition *ChainhookDefinition `json:"definition"`
	Status     ChainhookStatusInfo  `json:"status"`
}

// ============================================================================
// Pagination
// ============================================================================

// PaginationOptions represents pagination parameters.
type PaginationOptions struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

// PaginatedChainhookResponse represents a paginated response of chainhooks.
type PaginatedChainhookResponse struct {
	Total      uint64        `json:"total"`
	Offset     uint64        `json:"offset"`
	Limit      uint64        `json:"limit"`
	Chainhooks []Chainhook   `json:"chainhooks"`
}

// ============================================================================
// Bulk Operations
// ============================================================================

// BulkEnableChainhooksRequest represents a request to bulk enable/disable chainhooks.
type BulkEnableChainhooksRequest struct {
	Enabled       bool     `json:"enabled"`
	UUIDs         []UUID   `json:"uuids,omitempty"`
	WebhookURL    *string  `json:"webhook_url,omitempty"`
	Statuses      []ChainhookStatus `json:"statuses,omitempty"`
}

// BulkEnableChainhooksResponse represents the response from a bulk enable/disable operation.
type BulkEnableChainhooksResponse struct {
	UpdatedCount uint64 `json:"updated_count"`
}

// ============================================================================
// Consumer Secret
// ============================================================================

// ConsumerSecretResponse represents the consumer secret response.
type ConsumerSecretResponse struct {
	Secret string `json:"secret"`
}

// ============================================================================
// Evaluation
// ============================================================================

// EvaluateChainhookRequest represents a request to evaluate a chainhook.
type EvaluateChainhookRequest struct {
	BlockHeight uint64 `json:"block_height"`
}

// ============================================================================
// API Status
// ============================================================================

// ApiStatusResponse represents the API status response.
type ApiStatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
