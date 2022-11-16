// Package signature provides helper functions for handling the blockchain
// signature needs.
package signature

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ZeroHash represents a hash code of zeros.
const ZeroHash string = "0x0000000000000000000000000000000000000000000000000000000000000000"

// Hash returns a unique string for the value.
func Hash(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		return ZeroHash
	}

	hash := sha256.Sum256(data)
	return hexutil.Encode(hash[:])
}
