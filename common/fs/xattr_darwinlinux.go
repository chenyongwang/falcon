// Copyright (c) 2015 Rackspace
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// +build linux darwin

package fs

import (
	"syscall"
	"unsafe"
)

func fsetxattr(fd uintptr, attr string, value []byte) (int, error) {
	attrp, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return 0, err
	}
	valuep := unsafe.Pointer(&value[0])
	if r0, _, e1 := syscall.Syscall6(syscall.SYS_FSETXATTR, fd, uintptr(unsafe.Pointer(attrp)), uintptr(valuep), uintptr(len(value)), 0, 0); e1 == 0 {
		return int(r0), nil
	} else {
		return 0, e1
	}
}

func setxattr(path string, attr string, value []byte) (int, error) {
	attrp, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return 0, err
	}
	pathp, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	valuep := unsafe.Pointer(&value[0])
	if r0, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR, uintptr(unsafe.Pointer(pathp)), uintptr(unsafe.Pointer(attrp)), uintptr(valuep), uintptr(len(value)), 0, 0); e1 == 0 {
		return int(r0), nil
	} else {
		return 0, e1
	}
}

func fgetxattr(fd uintptr, attr string, value []byte) (int, error) {
	attrp, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return 0, err
	}
	var r0 uintptr
	var e1 syscall.Errno
	if len(value) == 0 {
		r0, _, e1 = syscall.Syscall6(syscall.SYS_FGETXATTR, fd, uintptr(unsafe.Pointer(attrp)), 0, 0, 0, 0)
	} else {
		valuep := unsafe.Pointer(&value[0])
		r0, _, e1 = syscall.Syscall6(syscall.SYS_FGETXATTR, fd, uintptr(unsafe.Pointer(attrp)), uintptr(valuep), uintptr(len(value)), 0, 0)
	}
	if e1 == 0 {
		return int(r0), nil
	} else {
		return 0, e1
	}
}

func getxattr(path string, attr string, value []byte) (int, error) {
	attrp, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return 0, err
	}
	pathp, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	var r0 uintptr
	var e1 syscall.Errno
	if len(value) == 0 {
		r0, _, e1 = syscall.Syscall6(syscall.SYS_GETXATTR, uintptr(unsafe.Pointer(pathp)), uintptr(unsafe.Pointer(attrp)), 0, 0, 0, 0)
	} else {
		valuep := unsafe.Pointer(&value[0])
		r0, _, e1 = syscall.Syscall6(syscall.SYS_GETXATTR, uintptr(unsafe.Pointer(pathp)), uintptr(unsafe.Pointer(attrp)), uintptr(valuep), uintptr(len(value)), 0, 0)
	}
	if e1 == 0 {
		return int(r0), nil
	} else {
		return 0, e1
	}
}
