package zenfunction

import (
	"github.com/zenmuharom/zenfunction/function"
	"github.com/zenmuharom/zenlogger"
)

func New(logger zenlogger.Zenlogger) function.Assigner {
	return &function.DefaultAssigner{
		Logger: logger,
	}
}
