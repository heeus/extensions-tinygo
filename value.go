/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

type value uint64

func (v value) Length() uint32 {
	return hostValueLength(uint64(v))
}

func (v value) AsString(name string) string {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	ptr := hostValueAsString(uint64(v), uint32(nh.Data), uint32(nh.Len))
	return decodeString(ptr)
}

func (v value) AsInt32(name string) int32 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return int32(hostValueAsInt32(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v value) AsValue(name string) IValue {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return value(hostValueAsValue(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v value) GetAsString(index int) string {
	ptr := hostValueGetAsString(uint64(v), uint32(index))
	return decodeString(ptr)
}

func (v value) GetAsInt32(index int) int32 {
	return int32(hostValueGetAsInt32(uint64(v), uint32(index)))
}

func (v value) GetAsValue(index int) IValue {
	return value(hostValueGetAsValue(uint64(v), uint32(index)))
}

func decodeString(value uint64) (ret string) {
	ptr := uint32(value >> 32)
	size := uint32(value)

	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&ret))
	strHdr.Data = uintptr(ptr)
	strHdr.Len = uintptr(size)
	return
}

type emptyValue struct{}

func (v emptyValue) Length() uint32               { return 0 }
func (v emptyValue) AsString(name string) string  { return "" }
func (v emptyValue) AsInt32(name string) int32    { return 0 }
func (v emptyValue) AsValue(name string) IValue   { return &emptyValue{} }
func (v emptyValue) GetAsInt32(index int) int32   { return 0 }
func (v emptyValue) GetAsString(index int) string { return "" }
func (v emptyValue) GetAsValue(index int) IValue  { return &emptyValue{} }

//export HostValueLength
func hostValueLength(id uint64) uint32

//export HostValueExists
func hostValueExists(id uint64) uint32

//export HostValueAsString
func hostValueAsString(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsInt32
func hostValueAsInt32(id uint64, namePtr, nameSize uint32) uint32

//export HostValueAsValue
func hostValueAsValue(id uint64, namePtr, nameSize uint32) uint64

// dddexport HostValueGetAsString
func hostValueGetAsString(id uint64, index uint32) uint64

// dddexport HostValueGetAsInt32
func hostValueGetAsInt32(id uint64, index uint32) uint32

// dddexport HostValueGetAsValue
func hostValueGetAsValue(id uint64, index uint32) uint64
