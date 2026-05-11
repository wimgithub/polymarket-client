package polyhttp

import (
	"errors"
	"net"
)

// IsTimeout reports whether err is or wraps a network timeout error.
func IsTimeout(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}
