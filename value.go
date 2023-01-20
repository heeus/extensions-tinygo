/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

type TValue uint64

func (v TValue) Length() uint32 {
	return hostValueLength(uint64(v))
}

func (v TValue) AsString(name string) string {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	ptr := hostValueAsString(uint64(v), uint32(nh.Data), uint32(nh.Len))
	return decodeString(ptr)
}

func (v TValue) AsInt32(name string) int32 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return int32(hostValueAsInt32(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v TValue) AsValue(name string) TValue {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return TValue(hostValueAsValue(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v TValue) GetAsString(index int) string {
	ptr := hostValueGetAsString(uint64(v), uint32(index))
	return decodeString(ptr)
}

func (v TValue) GetAsInt32(index int) int32 {
	return int32(hostValueGetAsInt32(uint64(v), uint32(index)))
}

func (v TValue) GetAsValue(index int) TValue {
	return TValue(hostValueGetAsValue(uint64(v), uint32(index)))
}

func decodeString(value uint64) (ret string) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&ret))
	strHdr.Data = uintptr(uint32(value >> 32))
	strHdr.Len = extint(uint32(value))

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
