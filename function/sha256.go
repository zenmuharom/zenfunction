package function

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Sha256(arg string) (hashed string, err error) {
	// Data to be hashed
	data := []byte(arg)

	// Create an Sha256 hash object
	hasher := sha256.New()

	// Write the data to the hasher
	_, err = hasher.Write(data)
	if err != nil {
		assigner.Logger.Error("JsonDecode", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return
	}

	// Calculate the Sha256 hash and store it in a byte slice
	hash := hasher.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashed = hex.EncodeToString(hash)

	return
}
