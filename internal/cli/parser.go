package cli

import (
	"context"
	"io"
)

// processor is the parser's implementation for handling inputs and generating
// outputs.
type processor interface {
	Read(context.Context, io.Reader) error
	Write(context.Context, io.Writer) error
}
