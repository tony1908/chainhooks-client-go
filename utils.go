package chainhooks

// ============================================================================
// Helper Functions for Building Filters
// ============================================================================

// StringPtr returns a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

// BoolPtr returns a pointer to the given bool.
func BoolPtr(b bool) *bool {
	return &b
}

// Uint64Ptr returns a pointer to the given uint64.
func Uint64Ptr(u uint64) *uint64 {
	return &u
}

// PrincipalStandard creates a standard principal.
func PrincipalStandard(address string) *Principal {
	return &Principal{
		Standard: StringPtr(address),
	}
}

// PrincipalContract creates a contract principal.
func PrincipalContract(address string) *Principal {
	return &Principal{
		Contract: StringPtr(address),
	}
}

// ============================================================================
// Builder for ChainhookDefinition
// ============================================================================

// ChainhookBuilder is a builder for constructing ChainhookDefinition objects.
type ChainhookBuilder struct {
	definition *ChainhookDefinition
	filters    []interface{}
	err        error
}

// NewChainhookBuilder creates a new ChainhookBuilder.
func NewChainhookBuilder(name string, network Network) *ChainhookBuilder {
	return &ChainhookBuilder{
		definition: &ChainhookDefinition{
			Name:    name,
			Version: DefaultAPIVersion,
			Chain:   DefaultChain,
			Network: network,
		},
		filters: []interface{}{},
	}
}

// WithName sets the chainhook name.
func (b *ChainhookBuilder) WithName(name string) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	b.definition.Name = name
	return b
}

// WithNetwork sets the network.
func (b *ChainhookBuilder) WithNetwork(network Network) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	b.definition.Network = network
	return b
}

// WithWebhookURL sets the webhook URL for the action.
func (b *ChainhookBuilder) WithWebhookURL(url string) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	b.definition.Action = ChainhookAction{
		Type: "http_post",
		URL:  url,
	}
	return b
}

// AddFilter adds an event filter to the chainhook.
func (b *ChainhookBuilder) AddFilter(filter EventFilter) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	b.filters = append(b.filters, filter)
	return b
}

// AddFTTransfer adds a fungible token transfer filter.
func (b *ChainhookBuilder) AddFTTransfer(asset string, sender, receiver *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&FTTransferFilter{
		Type:      EventTypeFTTransfer,
		Asset:     asset,
		Sender:    sender,
		Recipient: receiver,
		Amount:    amount,
	})
}

// AddFTMint adds a fungible token mint filter.
func (b *ChainhookBuilder) AddFTMint(asset string, recipient *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&FTMintFilter{
		Type:      EventTypeFTMint,
		Asset:     asset,
		Recipient: recipient,
		Amount:    amount,
	})
}

// AddFTBurn adds a fungible token burn filter.
func (b *ChainhookBuilder) AddFTBurn(asset string, sender *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&FTBurnFilter{
		Type:   EventTypeFTBurn,
		Asset:  asset,
		Sender: sender,
		Amount: amount,
	})
}

// AddNFTTransfer adds an NFT transfer filter.
func (b *ChainhookBuilder) AddNFTTransfer(asset string, sender, receiver *Principal) *ChainhookBuilder {
	return b.AddFilter(&NFTTransferFilter{
		Type:      EventTypeNFTTransfer,
		Asset:     asset,
		Sender:    sender,
		Recipient: receiver,
	})
}

// AddNFTMint adds an NFT mint filter.
func (b *ChainhookBuilder) AddNFTMint(asset string, recipient *Principal) *ChainhookBuilder {
	return b.AddFilter(&NFTMintFilter{
		Type:      EventTypeNFTMint,
		Asset:     asset,
		Recipient: recipient,
	})
}

// AddNFTBurn adds an NFT burn filter.
func (b *ChainhookBuilder) AddNFTBurn(asset string, sender *Principal) *ChainhookBuilder {
	return b.AddFilter(&NFTBurnFilter{
		Type:   EventTypeNFTBurn,
		Asset:  asset,
		Sender: sender,
	})
}

// AddSTXTransfer adds an STX transfer filter.
func (b *ChainhookBuilder) AddSTXTransfer(sender, receiver *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&STXTransferFilter{
		Type:      EventTypeSTXTransfer,
		Sender:    sender,
		Recipient: receiver,
		Amount:    amount,
	})
}

// AddSTXMint adds an STX mint filter.
func (b *ChainhookBuilder) AddSTXMint(recipient *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&STXMintFilter{
		Type:      EventTypeSTXMint,
		Recipient: recipient,
		Amount:    amount,
	})
}

// AddSTXBurn adds an STX burn filter.
func (b *ChainhookBuilder) AddSTXBurn(sender *Principal, amount *string) *ChainhookBuilder {
	return b.AddFilter(&STXBurnFilter{
		Type:   EventTypeSTXBurn,
		Sender: sender,
		Amount: amount,
	})
}

// AddContractDeploy adds a contract deployment filter.
func (b *ChainhookBuilder) AddContractDeploy(deployerPrincipal *Principal) *ChainhookBuilder {
	return b.AddFilter(&ContractDeployFilter{
		Type:              EventTypeContractDeploy,
		DeployerPrincipal: deployerPrincipal,
	})
}

// AddContractCall adds a contract call filter.
func (b *ChainhookBuilder) AddContractCall(contractID *string, method *string, sender *Principal) *ChainhookBuilder {
	return b.AddFilter(&ContractCallFilter{
		Type:               EventTypeContractCall,
		ContractIdentifier: contractID,
		Method:             method,
		Sender:             sender,
	})
}

// AddContractLog adds a contract log filter.
func (b *ChainhookBuilder) AddContractLog(contractID *string) *ChainhookBuilder {
	return b.AddFilter(&ContractLogFilter{
		Type:               EventTypeContractLog,
		ContractIdentifier: contractID,
	})
}

// AddBalanceChange adds a balance change filter.
func (b *ChainhookBuilder) AddBalanceChange(principal *Principal) *ChainhookBuilder {
	return b.AddFilter(&BalanceChangeFilter{
		Type:      EventTypeBalanceChange,
		Principal: principal,
	})
}

// AddCoinbase adds a coinbase filter.
func (b *ChainhookBuilder) AddCoinbase(recipient *Principal) *ChainhookBuilder {
	return b.AddFilter(&CoinbaseFilter{
		Type:      EventTypeCoinbase,
		Recipient: recipient,
	})
}

// AddTenureChange adds a tenure change filter.
func (b *ChainhookBuilder) AddTenureChange() *ChainhookBuilder {
	return b.AddFilter(&TenureChangeFilter{
		Type: EventTypeTenureChange,
	})
}

// WithOptions sets the options for the chainhook.
func (b *ChainhookBuilder) WithOptions(opts *ChainhookOptions) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	b.definition.Options = opts
	return b
}

// WithEnableOnRegistration sets whether to enable the chainhook on registration.
func (b *ChainhookBuilder) WithEnableOnRegistration(enable bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.EnableOnRegistration = BoolPtr(enable)
	return b
}

// WithExpireAfterEvaluations sets the expiration after evaluations.
func (b *ChainhookBuilder) WithExpireAfterEvaluations(count uint64) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.ExpireAfterEvaluations = Uint64Ptr(count)
	return b
}

// WithExpireAfterOccurrences sets the expiration after occurrences.
func (b *ChainhookBuilder) WithExpireAfterOccurrences(count uint64) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.ExpireAfterOccurrences = Uint64Ptr(count)
	return b
}

// WithDecodeClarityValues sets whether to decode Clarity values.
func (b *ChainhookBuilder) WithDecodeClarityValues(decode bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.DecodeClarityValues = BoolPtr(decode)
	return b
}

// WithIncludeContractABI sets whether to include contract ABI.
func (b *ChainhookBuilder) WithIncludeContractABI(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludeContractABI = BoolPtr(include)
	return b
}

// WithIncludeContractSourceCode sets whether to include contract source code.
func (b *ChainhookBuilder) WithIncludeContractSourceCode(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludeContractSourceCode = BoolPtr(include)
	return b
}

// WithIncludePostConditions sets whether to include post conditions.
func (b *ChainhookBuilder) WithIncludePostConditions(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludePostConditions = BoolPtr(include)
	return b
}

// WithIncludeRawTransactions sets whether to include raw transactions.
func (b *ChainhookBuilder) WithIncludeRawTransactions(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludeRawTransactions = BoolPtr(include)
	return b
}

// WithIncludeBlockSignatures sets whether to include block signatures.
func (b *ChainhookBuilder) WithIncludeBlockSignatures(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludeBlockSignatures = BoolPtr(include)
	return b
}

// WithIncludeBlockMetadata sets whether to include block metadata.
func (b *ChainhookBuilder) WithIncludeBlockMetadata(include bool) *ChainhookBuilder {
	if b.err != nil {
		return b
	}
	if b.definition.Options == nil {
		b.definition.Options = &ChainhookOptions{}
	}
	b.definition.Options.IncludeBlockMetadata = BoolPtr(include)
	return b
}

// Build validates and returns the ChainhookDefinition.
func (b *ChainhookBuilder) Build() (*ChainhookDefinition, error) {
	if b.err != nil {
		return nil, b.err
	}

	// Validate required fields
	if b.definition.Name == "" {
		return nil, &ValidationError{
			Field:  "name",
			Reason: "name is required",
		}
	}

	if b.definition.Action.URL == "" {
		return nil, &ValidationError{
			Field:  "action.url",
			Reason: "webhook URL is required",
		}
	}

	if len(b.filters) == 0 {
		return nil, &ValidationError{
			Field:  "filters",
			Reason: "at least one filter is required",
		}
	}

	b.definition.Filters = ChainhookFilters{
		Events: b.filters,
	}

	return b.definition, nil
}

// ============================================================================
// Helper Functions for Building Options
// ============================================================================

// NewChainhookOptions creates a new ChainhookOptions with all fields set to nil.
func NewChainhookOptions() *ChainhookOptions {
	return &ChainhookOptions{}
}

// ChainhookOptionsBuilder is a builder for constructing ChainhookOptions objects.
type ChainhookOptionsBuilder struct {
	opts *ChainhookOptions
}

// NewChainhookOptionsBuilder creates a new ChainhookOptionsBuilder.
func NewChainhookOptionsBuilder() *ChainhookOptionsBuilder {
	return &ChainhookOptionsBuilder{
		opts: &ChainhookOptions{},
	}
}

// EnableOnRegistration sets whether to enable the chainhook on registration.
func (b *ChainhookOptionsBuilder) EnableOnRegistration(enable bool) *ChainhookOptionsBuilder {
	b.opts.EnableOnRegistration = BoolPtr(enable)
	return b
}

// ExpireAfterEvaluations sets the expiration after evaluations.
func (b *ChainhookOptionsBuilder) ExpireAfterEvaluations(count uint64) *ChainhookOptionsBuilder {
	b.opts.ExpireAfterEvaluations = Uint64Ptr(count)
	return b
}

// ExpireAfterOccurrences sets the expiration after occurrences.
func (b *ChainhookOptionsBuilder) ExpireAfterOccurrences(count uint64) *ChainhookOptionsBuilder {
	b.opts.ExpireAfterOccurrences = Uint64Ptr(count)
	return b
}

// DecodeClarityValues sets whether to decode Clarity values.
func (b *ChainhookOptionsBuilder) DecodeClarityValues(decode bool) *ChainhookOptionsBuilder {
	b.opts.DecodeClarityValues = BoolPtr(decode)
	return b
}

// IncludeContractABI sets whether to include contract ABI.
func (b *ChainhookOptionsBuilder) IncludeContractABI(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludeContractABI = BoolPtr(include)
	return b
}

// IncludeContractSourceCode sets whether to include contract source code.
func (b *ChainhookOptionsBuilder) IncludeContractSourceCode(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludeContractSourceCode = BoolPtr(include)
	return b
}

// IncludePostConditions sets whether to include post conditions.
func (b *ChainhookOptionsBuilder) IncludePostConditions(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludePostConditions = BoolPtr(include)
	return b
}

// IncludeRawTransactions sets whether to include raw transactions.
func (b *ChainhookOptionsBuilder) IncludeRawTransactions(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludeRawTransactions = BoolPtr(include)
	return b
}

// IncludeBlockSignatures sets whether to include block signatures.
func (b *ChainhookOptionsBuilder) IncludeBlockSignatures(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludeBlockSignatures = BoolPtr(include)
	return b
}

// IncludeBlockMetadata sets whether to include block metadata.
func (b *ChainhookOptionsBuilder) IncludeBlockMetadata(include bool) *ChainhookOptionsBuilder {
	b.opts.IncludeBlockMetadata = BoolPtr(include)
	return b
}

// Build returns the ChainhookOptions.
func (b *ChainhookOptionsBuilder) Build() *ChainhookOptions {
	return b.opts
}

// ============================================================================
// Query Builder Helpers
// ============================================================================

// NewPaginationOptions creates a new PaginationOptions with the given offset and limit.
func NewPaginationOptions(offset, limit uint64) *PaginationOptions {
	return &PaginationOptions{
		Offset: offset,
		Limit:  limit,
	}
}

// BulkEnableUUIDs creates a bulk enable request for specific UUIDs.
func BulkEnableUUIDs(enabled bool, uuids ...UUID) *BulkEnableChainhooksRequest {
	return &BulkEnableChainhooksRequest{
		Enabled: enabled,
		UUIDs:   uuids,
	}
}

// BulkEnableByWebhook creates a bulk enable request for chainhooks with a specific webhook URL.
func BulkEnableByWebhook(enabled bool, webhookURL string) *BulkEnableChainhooksRequest {
	return &BulkEnableChainhooksRequest{
		Enabled:    enabled,
		WebhookURL: StringPtr(webhookURL),
	}
}

// BulkEnableByStatus creates a bulk enable request for chainhooks with specific statuses.
func BulkEnableByStatus(enabled bool, statuses ...ChainhookStatus) *BulkEnableChainhooksRequest {
	return &BulkEnableChainhooksRequest{
		Enabled:  enabled,
		Statuses: statuses,
	}
}

// ============================================================================
// Error Helpers
// ============================================================================

// IsHttpError checks if an error is an HttpError.
func IsHttpError(err error) bool {
	_, ok := err.(*HttpError)
	return ok
}

// AsHttpError converts an error to HttpError if possible.
func AsHttpError(err error) (*HttpError, bool) {
	httpErr, ok := err.(*HttpError)
	return httpErr, ok
}

// GetHttpStatusCode extracts the HTTP status code from an error if it's an HttpError.
func GetHttpStatusCode(err error) (int, bool) {
	httpErr, ok := AsHttpError(err)
	if !ok {
		return 0, false
	}
	return httpErr.StatusCode, true
}

// IsNotFound checks if an error is a 404 Not Found error.
func IsNotFound(err error) bool {
	statusCode, ok := GetHttpStatusCode(err)
	return ok && statusCode == 404
}

// IsUnauthorized checks if an error is a 401 Unauthorized error.
func IsUnauthorized(err error) bool {
	statusCode, ok := GetHttpStatusCode(err)
	return ok && statusCode == 401
}

// IsForbidden checks if an error is a 403 Forbidden error.
func IsForbidden(err error) bool {
	statusCode, ok := GetHttpStatusCode(err)
	return ok && statusCode == 403
}

// IsServerError checks if an error is a 5xx server error.
func IsServerError(err error) bool {
	statusCode, ok := GetHttpStatusCode(err)
	return ok && statusCode >= 500
}

// IsClientError checks if an error is a 4xx client error.
func IsClientError(err error) bool {
	statusCode, ok := GetHttpStatusCode(err)
	return ok && statusCode >= 400 && statusCode < 500
}
