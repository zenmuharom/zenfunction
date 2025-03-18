package function

import (
	"github.com/google/uuid"
)

func (assigner *DefaultAssigner) Uuid() (generated string, err error) {
	id := uuid.New()
	generated = id.String()
	return
}
