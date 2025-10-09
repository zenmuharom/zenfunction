package function

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) hmacSha256(data, key string) (hashed string, err error) {
	if data == "" || key == "" {
		err = errors.New("Invalid number of arguments")
		assigner.Logger.Error("hmacSha256", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return "", err
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	hash := h.Sum(nil)

	hashed = hex.EncodeToString(hash)
	return hashed, nil
}
