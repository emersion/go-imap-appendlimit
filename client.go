package appendlimit

import (
	"github.com/emersion/go-imap"
)

func MailboxStatusAppendLimit(status *imap.MailboxStatus) *uint32 {
	val := status.Items[StatusAppendLimit]
	if val == nil {
		return nil
	}

	res, err := imap.ParseNumber(val)
	if err != nil {
		return nil
	}

	return &res
}
