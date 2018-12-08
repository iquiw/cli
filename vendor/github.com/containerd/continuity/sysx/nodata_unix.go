// +build darwin freebsd netbsd

package sysx

import (
	"syscall"
)

const ENODATA = syscall.ENOATTR
