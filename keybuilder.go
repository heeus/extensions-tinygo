/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

func keyBuilderImpl(storage, entity string) TKeyBuilder {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&storage))
	eh := (*reflect.StringHeader)(unsafe.Pointer(&entity))
	return TKeyBuilder(hostGetKey(uint32(sh.Data), uint32(sh.Len), uint32(eh.Data), uint32(eh.Len)))
}

func (k TKeyBuilder) PutInt32(name string, value int32) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	hostRowWriterPutInt32(uint64(k), 0, uint32(nh.Data), uint32(nh.Len), uint32(value))
}

func (k TKeyBuilder) PutString(name string, value string) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	vh := (*reflect.StringHeader)(unsafe.Pointer(&value))
	hostRowWriterPutString(uint64(k), 0, uint32(nh.Data), uint32(nh.Len), uint32(vh.Data), uint32(vh.Len))
}

//export hostGetKey
func hostGetKey(storagePtr, storageSize, entityPtr, entitySize uint32) uint64
