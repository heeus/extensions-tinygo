/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

func keyBuilderImpl(storage, entity string) IKeyBuilder {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&storage))
	eh := (*reflect.StringHeader)(unsafe.Pointer(&entity))
	return keyBuilder(hostGetKey(uint32(sh.Data), uint32(sh.Len), uint32(eh.Data), uint32(eh.Len)))
}

type keyBuilder uint64

func (k keyBuilder) PutInt32(name string, value int32) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	hostKeyPutInt32(uint64(k), uint32(nh.Data), uint32(nh.Len), uint32(value))
}

func (k keyBuilder) PutString(name string, value string) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	vh := (*reflect.StringHeader)(unsafe.Pointer(&value))
	hostKeyPutString(uint64(k), uint32(nh.Data), uint32(nh.Len), uint32(vh.Data), uint32(vh.Len))
}

//export HostGetKey
func hostGetKey(storagePtr, storageSize, entityPtr, entitySize uint32) uint64

//export HostKeyPutString
func hostKeyPutString(id uint64, namePtr, nameSize, valuePtr, valueSize uint32)

//export HostKeyPutInt32
func hostKeyPutInt32(id uint64, namePtr, nameSize, value uint32)
