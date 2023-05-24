package function

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/zenmuharom/zenlogger"
)

func (assigner *DefaultAssigner) RandomInt(from, to int) (generated string, err error) {
	assigner.Logger.Debug("RandomInt", zenlogger.ZenField{Key: "from", Value: from}, zenlogger.ZenField{Key: "to", Value: to})

	if from < 0 {
		from = 0
	}
	re := regexp.MustCompile(`[a-z,-]`)
	id := uuid.New().String()
	if to < from {
		to = from
	}
	generated = re.ReplaceAllString(id, "")[from:to]
	return
}
