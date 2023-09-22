package function

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) MD5(arg string) (hashed string, err error) {
	// Data to be hashed
	data := []byte(arg)

	// Create an MD5 hash object
	hasher := md5.New()

	// Write the data to the hasher
	_, err = hasher.Write(data)
	if err != nil {
		assigner.Logger.Error("JsonDecode", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return
	}

	// Calculate the MD5 hash and store it in a byte slice
	hash := hasher.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashed = hex.EncodeToString(hash)

	return
}
