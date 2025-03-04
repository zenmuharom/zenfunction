package function

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"strings"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) EncryptWithPrivateKey(rawData, privateKey string) (string, error) {
	data := []byte(rawData)
	digest := sha256.Sum256(data)
	key, err := assigner.getPrivateKey(privateKey)
	if err != nil {
		assigner.Logger.Error("EncryptWithPrivateKey", zenlogger.ZenField{Key: "error", Value: err.Error()})
		return "", err
	}

	signature, signErr := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if signErr != nil {
		assigner.Logger.Error("EncryptWithPrivateKey", zenlogger.ZenField{Key: "error", Value: signErr.Error()})
		return "", signErr
	}

	b64sig := base64.StdEncoding.EncodeToString(signature)
	return b64sig, nil
}

func (assigner *DefaultAssigner) getPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	r := strings.NewReader(privateKey)
	pemBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.New("error reading private key")
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("error decoding private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("error parsing private key")
	}

	return key.(*rsa.PrivateKey), nil
}
