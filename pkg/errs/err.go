package errs

import "errors"

var (
	ErrActIdInvalid      = errors.New("actId is invalid")
	ErrLivePlatIdInvalid = errors.New("livePlatId is invalid")
	ErrGameIdInvalid     = errors.New("gameId is invalid")
	ErrUserInvalid       = errors.New("user is invalid")
)
