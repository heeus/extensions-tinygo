/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

type TKey uint64

func (v TKey) AsString(name string) string {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	ptr := hostKeyAsString(uint64(v), uint32(nh.Data), uint32(nh.Len))
	return decodeString(ptr)
}

func (v TKey) AsInt32(name string) int32 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return int32(hostKeyAsInt32(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

//export HostKeyAsString
func hostKeyAsString(id uint64, namePtr, nameSize uint32) uint64

//export HostKeyAsInt32
func hostKeyAsInt32(id uint64, namePtr, nameSize uint32) uint32
