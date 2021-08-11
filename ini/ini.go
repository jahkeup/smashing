package ini

import (
	"bufio"
	"context"
	"io"

	"gopkg.in/ini.v1"
)

type LoadOptions = ini.LoadOptions

var defaultLoadOptions = ini.LoadOptions{
	AllowPythonMultilineValues: true,
	AllowNestedValues:          true,
}

type ReadWriter struct {
	tree *ini.File
}

func NewReadWriter(parserOptions *ini.LoadOptions) *ReadWriter {
	var tree *ini.File

	if parserOptions == nil {
		tree = ini.Empty(defaultLoadOptions)
	} else {
		tree = ini.Empty(defaultLoadOptions, *parserOptions)
	}

	return &ReadWriter{tree}
}

func (r *ReadWriter) Read(ctx context.Context, rd io.Reader) error {
	return r.tree.Append(bufio.NewReader(rd))
}

func (r *ReadWriter) Write(ctx context.Context, wr io.Writer) error {
	_, err := r.tree.WriteTo(wr)
	return err
}
