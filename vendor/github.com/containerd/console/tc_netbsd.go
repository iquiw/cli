/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package console

import (
	"bytes"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	cmdTcGet = unix.TIOCGETA
	cmdTcSet = unix.TIOCSETA
)

type ptmget struct {
	cfd uint16
	sfd uint16
	cn [1024]byte
	sn [1024]byte
}

// unlockpt unlocks the slave pseudoterminal device corresponding to the master pseudoterminal referred to by f.
// unlockpt should be called before opening the slave side of a pty.
func unlockpt(f *os.File) error {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(f.Fd()), unix.TIOCGRANTPT, 0)
	if err != 0 {
		return err
	}
	return nil
}

// ptsname retrieves the name of the first available pts for the given master.
func ptsname(f *os.File) (string, error) {
	ptmget := new(ptmget)
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(f.Fd()), unix.TIOCPTSNAME, uintptr(unsafe.Pointer(ptmget)))
	if  err != 0 {
		return "", err
	}
	n := bytes.IndexByte(ptmget.sn[:], 0)
	return string(ptmget.sn[:n]), nil
}
