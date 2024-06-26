package utils

import "fmt"

// InvalidRpcError represents an error for an invalid Arbitrum RPC URL.
type InvalidRpcError struct {
	URL     string
	Message string
}

// Error implements the error interface for InvalidRpcError.
func (e *InvalidRpcError) Error() string {
	return fmt.Sprintf("Invalid Arbitrum RPC URL: %s - %s", e.URL, e.Message)
}

// NewInvalidRpcError creates a new InvalidRpcError.
func NewInvalidRpcError(url, message string) error {
	return &InvalidRpcError{
		URL:     url,
		Message: message,
	}
}

// FailedFetchPositionsError represents an error for failing to fetch futures positions from the blockchain.
type FailedFetchPositionsError struct {
	Account string
	Message string
}

// Error implements the error interface for FailedFetchPositionsError.
func (e *FailedFetchPositionsError) Error() string {
	return fmt.Sprintf("Failed to fetch positions for account %s: %s", e.Account, e.Message)
}

// NewFailedFetchPositionsError creates a new FailedFetchPositionsError.
func NewFailedFetchPositionsError(account, message string) error {
	return &FailedFetchPositionsError{
		Account: account,
		Message: message,
	}
}

func NewFailedFetchOptionPositionsError(message string) error {
	return &FailedFetchPositionsError{
		Account: "",
		Message: message,
	}
}

// InvalidAddressError is a custom error type for invalid Ethereum addresses
type InvalidAddressError struct {
	Address string
}

// Error implements the error interface for InvalidAddressError
func (e *InvalidAddressError) Error() string {
	return fmt.Sprintf("Invalid Ethereum address: %s", e.Address)
}

func NewInvalidAddressError(address string) *InvalidAddressError {
	return &InvalidAddressError{Address: address}
}

type InvalidContractAddressError struct {
	Address string
	Message string
}

// Implement the Error() method for InvalidContractAddressError
func (e *InvalidContractAddressError) Error() string {
	return fmt.Sprintf("Invalid contract address: %s - %s", e.Address, e.Message)
}

// NewInvalidContractAddressError is the constructor for InvalidContractAddressError
func NewInvalidContractAddressError(address, message string) *InvalidContractAddressError {
	return &InvalidContractAddressError{
		Address: address,
		Message: message,
	}
}

// FailedContractCallError represents an error for a failed contract call
type FailedContractCallError struct {
	Address string
	Message string
}

// Implement the Error() method for FailedContractCallError
func (e *FailedContractCallError) Error() string {
	return fmt.Sprintf("Failed contract call to address: %s - %s", e.Address, e.Message)
}

// NewFailedContractCallError is the constructor for FailedContractCallError
func NewFailedContractCallError(address, message string) *FailedContractCallError {
	return &FailedContractCallError{
		Address: address,
		Message: message,
	}
}

// PriceCacheMissingError represents an error when a specific price cache entry is missing.
type PriceCacheMissingError struct {
}

// Error implements the error interface.
func (e *PriceCacheMissingError) Error() string {
	return fmt.Sprintf("price cache missing")
}

// NewPriceCacheMissingError creates a new instance of PriceCacheMissingError.
func NewPriceCacheMissingError() *PriceCacheMissingError {
	return &PriceCacheMissingError{}
}

type StreamFailedToStartError struct{}

// Error implements the error interface for StreamFailedToStartError.
func (e *StreamFailedToStartError) Error() string {
	return "stream failed to start"
}

// NewStreamFailedToStartError creates a new instance of StreamFailedToStartError.
func NewStreamFailedToStartError() *StreamFailedToStartError {
	return &StreamFailedToStartError{}
}

type UserAlreadyInCacheError struct {
	Username string
}

// Implement the Error method for the custom error type
func (e *UserAlreadyInCacheError) Error() string {
	return fmt.Sprintf("User '%s' is already in the cache", e.Username)
}

// Constructor function for creating a new UserAlreadyInCacheError
func NewUserAlreadyInCacheError(username string) error {
	return &UserAlreadyInCacheError{Username: username}
}

// Define the custom error type with an exportable name
type CacheStreamNotRunningError struct {
}

// Implement the Error method for the custom error type
func (e *CacheStreamNotRunningError) Error() string {
	return "Cache is not running"
}

// Constructor function for creating a new CacheStreamNotRunningError
func NewCacheStreamNotRunningError() error {
	return &CacheStreamNotRunningError{}
}

// positionsDifferentTypeError is a custom error type.
type PositionsDifferentTypeError struct {
	message string
}

// NewPositionsDifferentTypeError is a constructor function for positionsDifferentTypeError.
func NewPositionsDifferentTypeError(message string) *PositionsDifferentTypeError {
	return &PositionsDifferentTypeError{message: message}
}

// Error implements the error interface for positionsDifferentTypeError.
func (e *PositionsDifferentTypeError) Error() string {
	return e.message
}
