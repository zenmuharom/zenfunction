package function

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Sha1(arg string) (hashed string, err error) {
	// Data to be hashed
	data := []byte(arg)

	// Create a SHA1 hash object
	hasher := sha1.New()

	// Write the data to the hasher
	_, err = hasher.Write(data)
	if err != nil {
		assigner.Logger.Error("Sha1", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return
	}

	// Calculate the SHA1 hash and store it in a byte slice
	hash := hasher.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashed = hex.EncodeToString(hash)

	return
}
