//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"github.com/google/wire"
)

// Naraku is a struct that contains the settings for the naraku.
type Naraku struct{}

// NewNaraku returns a new Naraku struct.
func NewNaraku() *Naraku {
	wire.Build(
		newNaraku,
	)
	return nil
}

// newNaraku returns a new naraku struct.
func newNaraku() *Naraku {
	return &Naraku{}
}
