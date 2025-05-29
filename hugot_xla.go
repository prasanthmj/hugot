//go:build XLA || ALL

package hugot

import (
	"github.com/prasanthmj/hugot/options"
)

func NewXLASession(opts ...options.WithOption) (*Session, error) {
	return newSession("XLA", opts...)
}
