package function

import (
	"fmt"
	"strings"

	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) Trim(arg0, arg1 string) (trimmed string, err error) {
	assigner.logger.Debug("Trim", zenlogger.ZenField{Key: "arg0", Value: fmt.Sprintf("%v", arg0)}, zenlogger.ZenField{Key: "arg1", Value: arg1})

	trimmed = strings.Trim(fmt.Sprintf("%v", arg0), arg1)
	return
}
