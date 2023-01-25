/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"unsafe"
)

func (v TValue) Length() uint32 {
	return hostValueLength(uint64(v))
}

func (v TValue) AsString(name string) string {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	ptr := hostValueAsString(uint64(v), uint32(nh.Data), uint32(nh.Len))
	return decodeString(ptr)
}

func (v TValue) AsBytes(name string) (ret []byte) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	ptr := hostValueAsBytes(uint64(v), uint32(nh.Data), uint32(nh.Len))

	strHdr := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	strHdr.Data = uintptr(uint32(ptr >> 32))
	strHdr.Len = extint(uint32(ptr))
	return
}

func (v TValue) AsInt32(name string) int32 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return int32(hostValueAsInt32(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v TValue) AsInt64(name string) int64 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return int64(hostValueAsInt64(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v TValue) AsFloat32(name string) float32 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return hostValueAsFloat32(uint64(v), uint32(nh.Data), uint32(nh.Len))
}

func (v TValue) AsFloat64(name string) float64 {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return hostValueAsFloat64(uint64(v), uint32(nh.Data), uint32(nh.Len))
}

func (v TValue) AsQName(name string) QName {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	pkgPtr := hostValueAsQNamePkg(uint64(v), uint32(nh.Data), uint32(nh.Len))
	entityPtr := hostValueAsQNameEntity(uint64(v), uint32(nh.Data), uint32(nh.Len))
	return QName{
		Pkg:    decodeString(pkgPtr),
		Entity: decodeString(entityPtr),
	}
}

func (v TValue) AsBool(name string) bool {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return hostValueAsBool(uint64(v), uint32(nh.Data), uint32(nh.Len)) > 0
}

func (v TValue) AsValue(name string) TValue {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&name))
	return TValue(hostValueAsValue(uint64(v), uint32(nh.Data), uint32(nh.Len)))
}

func (v TValue) GetAsBytes(index int) (ret []byte) {
	ptr := hostValueGetAsBytes(uint64(v), uint32(index))
	strHdr := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	strHdr.Data = uintptr(uint32(ptr >> 32))
	strHdr.Len = extint(uint32(ptr))
	return
}

func (v TValue) GetAsString(index int) string {
	ptr := hostValueGetAsString(uint64(v), uint32(index))
	return decodeString(ptr)
}

func (v TValue) GetAsInt32(index int) int32 {
	return int32(hostValueGetAsInt32(uint64(v), uint32(index)))
}

func (v TValue) GetAsInt64(index int) int64 {
	return int64(hostValueGetAsInt64(uint64(v), uint32(index)))
}

func (v TValue) GetAsFloat32(index int) float32 {
	return hostValueGetAsFloat32(uint64(v), uint32(index))
}

func (v TValue) GetAsFloat64(index int) float64 {
	return hostValueGetAsFloat64(uint64(v), uint32(index))
}

func (v TValue) GetAsValue(index int) TValue {
	return TValue(hostValueGetAsValue(uint64(v), uint32(index)))
}

func (v TValue) GetAsQName(index int) QName {
	pkgPtr := hostValueGetAsQNamePkg(uint64(v), uint32(index))
	entityPtr := hostValueGetAsQNameEntity(uint64(v), uint32(index))
	return QName{
		Pkg:    decodeString(pkgPtr),
		Entity: decodeString(entityPtr),
	}
}

func (v TValue) GetAsBool(index int) bool {
	return hostValueGetAsBool(uint64(v), uint32(index)) > 0
}

func decodeString(value uint64) (ret string) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&ret))
	strHdr.Data = uintptr(uint32(value >> 32))
	strHdr.Len = extint(uint32(value))
	return
}

//export HostValueLength
func hostValueLength(id uint64) uint32

//export HostValueAsBytes
func hostValueAsBytes(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsString
func hostValueAsString(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsInt32
func hostValueAsInt32(id uint64, namePtr, nameSize uint32) uint32

//export HostValueAsInt64
func hostValueAsInt64(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsFloat32
func hostValueAsFloat32(id uint64, namePtr, nameSize uint32) float32

//export HostValueAsFloat64
func hostValueAsFloat64(id uint64, namePtr, nameSize uint32) float64

//export HostValueAsValue
func hostValueAsValue(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsQNamePkg
func hostValueAsQNamePkg(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsQNameEntity
func hostValueAsQNameEntity(id uint64, namePtr, nameSize uint32) uint64

//export HostValueAsBool
func hostValueAsBool(id uint64, namePtr, nameSize uint32) uint64

//export HostValueGetAsBytes
func hostValueGetAsBytes(id uint64, index uint32) uint64

//export HostValueGetAsString
func hostValueGetAsString(id uint64, index uint32) uint64

//export HostValueGetAsInt32
func hostValueGetAsInt32(id uint64, index uint32) uint32

//export HostValueGetAsInt64
func hostValueGetAsInt64(id uint64, index uint32) uint64

//export HostValueGetAsFloat32
func hostValueGetAsFloat32(id uint64, index uint32) float32

//export HostValueGetAsFloat64
func hostValueGetAsFloat64(id uint64, index uint32) float64

//export HostValueGetAsValue
func hostValueGetAsValue(id uint64, index uint32) uint64

//export HostValueGetAsQNamePkg
func hostValueGetAsQNamePkg(id uint64, index uint32) uint64

//export HostValueGetAsQNameEntity
func hostValueGetAsQNameEntity(id uint64, index uint32) uint64

//export HostValueGetAsBool
func hostValueGetAsBool(id uint64, index uint32) uint64
