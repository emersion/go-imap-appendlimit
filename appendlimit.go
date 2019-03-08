// Implements the IMAP APPENDLIMIT Extension, as defined in RFC 7889.
package appendlimit

import (
	"github.com/emersion/go-imap"
)

const Capability = "APPENDLIMIT"

const StatusAppendLimit imap.StatusItem = "APPENDLIMIT"

const codeTooBig imap.StatusRespCode = "TOOBIG"
