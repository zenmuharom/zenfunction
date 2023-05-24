package zenfunction

import (
	"github.com/zenmuharom/zenfunction/function"
	"github.com/zenmuharom/zenlogger"
)

type Zenfunction interface {
	function.Assigner
}

func New(logger zenlogger.Zenlogger) Zenfunction {
	return &function.DefaultAssigner{
		Logger: logger,
	}
}
