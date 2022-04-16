package brp

import "errors"

var (
	ErrUnknownTag        = errors.New("unknown tag")
	ErrPacketTagMismatch = errors.New("packet tag mismatch")
)
