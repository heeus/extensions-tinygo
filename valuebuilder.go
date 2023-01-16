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
type intent = valueBuilder

func (i valueBuilder) PutInt32(name string, value int32) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	hostIntentPutInt32(uint64(i), uint32(nh.Data), uint32(nh.Len), uint32(value))
}

func (i valueBuilder) PutString(name string, value string) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	vh := (*reflect.StringHeader)(unsafe.Pointer(&value))
	hostIntentPutString(uint64(i), uint32(nh.Data), uint32(nh.Len), uint32(vh.Data), uint32(vh.Len))
}

func (i valueBuilder) PutValue(name string) IValueBuilder {
	return nil
	// TODO
}

func (i valueBuilder) SetInt32(index int, value int32) {
	// TODO
}

func (i valueBuilder) SetString(index int, value string) {
	// TODO
}

func (i valueBuilder) SetValue(index int) IValueBuilder {
	return nil
	// TODO
}

//export HostIntentPutString
func hostIntentPutString(id uint64, namePtr, nameSize, valuePtr, valueSize uint32)

//export HostIntentPutInt32
func hostIntentPutInt32(id uint64, namePtr, nameSize, value uint32)
