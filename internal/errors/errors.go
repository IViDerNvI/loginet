package errors

import (
	err "golang.org/x/xerrors"
)

var (
	ErrSuccess   = err.New("log in successful")
	ErrDuplicate = err.New("you have logged in")
	ErrIpError   = err.New("ip address error")
	ErrUnknown   = err.New("unknown error")
	ErrMaxConn   = err.New("too much connections")
)
