package decoder

import (
	"context"

	"github.com/jslyzt/go-json/internal/option/decode"
)

type Option struct {
	Flags   decode.OptionFlag
	Context context.Context
	Path    *Path
}

func GetOptionFlag(xopt *Option) *decode.OptionFlag {
	if xopt != nil {
		return &xopt.Flags
	}
	return nil
}
