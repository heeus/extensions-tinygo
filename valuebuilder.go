/*
* Copyright (c) 2021-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

type valueBuilder uint64
type TIntent = valueBuilder

func (i valueBuilder) PutInt32(name string, value int32) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	hostIntentPutInt32(uint64(i), uint32(nh.Data), uint32(nh.Len), uint32(value))
}

func (i valueBuilder) PutString(name string, value string) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	vh := (*reflect.StringHeader)(unsafe.Pointer(&value))
	hostIntentPutString(uint64(i), uint32(nh.Data), uint32(nh.Len), uint32(vh.Data), uint32(vh.Len))
}

//export HostIntentPutString
func hostIntentPutString(id uint64, namePtr, nameSize, valuePtr, valueSize uint32)

//export HostIntentPutInt32
func hostIntentPutInt32(id uint64, namePtr, nameSize, value uint32)
