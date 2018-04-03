package proio

import (
	"io"

	"go-hep.org/x/hep/fwk"
)

// InputStream implements the go-hep.org/x/hep/fwk Task interface
type InputStream struct {
	R io.Reader

	outputs []string
	rdr     *Reader
}

// Connect establishes output ports (returning an error for the wrong type),
// and sets up a Reader with R (io.Reader) as the raw input.
func (stream *InputStream) Connect(ports []fwk.Port) error {
	for _, port := range ports {
		stream.outputs = append(stream.outputs, port.Name)
	}

	stream.rdr = NewReader(stream.R)

	return nil
}

// Read grabs the next Event from the underlying Reader.
func (stream *InputStream) Read(ctx fwk.Context) error {
	event, err := stream.rdr.Next()
	if err != nil {
		return err
	}

	for _, output := range stream.outputs {
		if err = ctx.Store().Put(output, event); err != nil {
			return err
		}
	}

	return nil
}

// Disconnect closes the underlying Reader, but leaves R open.
func (stream *InputStream) Disconnect() error {
	stream.rdr.Close()

	return nil
}
