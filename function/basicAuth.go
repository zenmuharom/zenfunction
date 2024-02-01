package function

import (
	"encoding/base64"
	"fmt"
)

func (assigner *DefaultAssigner) BasicAuth(arg string) (basicAuth string, err error) {
	// Set your credentials here
	base64encoded := base64.StdEncoding.EncodeToString([]byte(arg))
	basicAuth = fmt.Sprintf("Basic %v", base64encoded)

	return
}
