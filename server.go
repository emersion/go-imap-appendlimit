package appendlimit

import (
	"fmt"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/server"
)

// An error that should be returned by User.CreateMessage when the message size
// is too big.
var ErrTooBig = server.ErrStatusResp(&imap.StatusResp{
	Type: imap.StatusRespNo,
	Code: codeTooBig,
	Info: "Message size exceeding limit",
})

// A backend that supports retrieving per-user message size limits.
type Backend interface {
	// Get the fixed maximum message size in octets that the backend will accept
	// when creating a new message. If there is no limit, return nil.
	CreateMessageLimit() *uint32
}

// A user that supports retrieving per-user message size limits.
type User interface {
	// Get the fixed maximum message size in octets that the backend will accept
	// when creating a new message. If there is no limit, return nil.
	//
	// This overrides the global backend limit.
	CreateMessageLimit() *uint32
}

// StatusSetAppendLimit sets limit value in MailboxStatus object,
// nil pointer value will remove limit.
func StatusSetAppendLimit(status *imap.MailboxStatus, value *uint32) {
	if value != nil {
		status.Items[StatusAppendLimit] = *value
	} else {
		delete(status.Items, StatusAppendLimit)
	}
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

// Create a new server extension.
func NewExtension() server.Extension {
	return &extension{}
}
