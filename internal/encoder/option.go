package encoder

import (
	"context"
	"io"

	"github.com/jslyzt/go-json/internal/option/encode"
)

type Option struct {
	Flag        encode.OptionFlag
	ColorScheme *ColorScheme
	Context     context.Context
	DebugOut    io.Writer
	DebugDOTOut io.WriteCloser
}

type EncodeFormat struct {
	Header string
	Footer string
}

type EncodeFormatScheme struct {
	Int       EncodeFormat
	Uint      EncodeFormat
	Float     EncodeFormat
	Bool      EncodeFormat
	String    EncodeFormat
	Binary    EncodeFormat
	ObjectKey EncodeFormat
	Null      EncodeFormat
}

type (
	ColorScheme = EncodeFormatScheme
	ColorFormat = EncodeFormat
)

func GetOptionFlag(xopt *Option) *encode.OptionFlag {
	if xopt != nil {
		return &xopt.Flag
	}
	return nil
}
