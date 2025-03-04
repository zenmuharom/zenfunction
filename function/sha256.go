package function

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Sha256(args ...string) (hashed string, err error) {
	if len(args) == 0 || len(args) > 2 {
		err = errors.New("Invalid number of arguments")
		assigner.Logger.Error("Sha256", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return "", errors.New(err.Error())
	}

	// Data to be hashed
	data := []byte(args[0])

	var hash []byte

	if len(args) == 2 {
		// HMAC-SHA256 with private key
		key := []byte(args[1])
		h := hmac.New(sha256.New, key)
		h.Write(data)
		hash = h.Sum(nil)
	} else {
		// Standard SHA256 hash
		h := sha256.New()
		h.Write(data)
		hash = h.Sum(nil)
	}

	// Convert hash to hexadecimal string
	hashed = hex.EncodeToString(hash)
	return hashed, nil
}
