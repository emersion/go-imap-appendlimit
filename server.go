package appendlimit

import (
	"fmt"

	"github.com/emersion/go-imap/server"
)

// A backend that supports retrieving per-user message size limits.
type Backend interface {
	// Get the fixed maximum message size in octets that the backend will accept
	// when creating a new message.
	CreateMessageLimit() *uint32
}

// A user that supports retrieving per-user message size limits.
type User interface {
	// Get the fixed maximum message size in octets that the backend will accept
	// when creating a new message.
	//
	// This overrides the global backend limit.
	CreateMessageLimit() *uint32
}

type extension struct{}

func formatCapability(limit uint32) string {
	return fmt.Sprintf("%v=%v", Capability, limit)
}

func (ext *extension) Capabilities(c server.Conn) []string {
	// If the user is authenticated, maybe he has a limit set
	if u := c.Context().User; u != nil {
		if u, ok := u.(User); ok {
			if limit := u.CreateMessageLimit(); limit != nil {
				return []string{formatCapability(*limit)}
			}
		}
	}

	// User has no specific limit, let's check if the backend has one
	if be, ok := c.Server().Backend.(Backend); ok {
		if limit := be.CreateMessageLimit(); limit != nil {
			return []string{formatCapability(*limit)}
		}
	}

	// No limit available, just advertise for APPENDLIMIT support
	return []string{Capability}
}

func (ext *extension) Command(name string) server.HandlerFactory {
	return nil
}

func NewExtension() server.Extension {
	return &extension{}
}