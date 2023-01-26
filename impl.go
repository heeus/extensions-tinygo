/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import (
	"reflect"
	"runtime"
	"unsafe"
)

func Assert(condition bool, msg string) {
	if !condition {
		Panic("assertion failed: " + msg)
	}
}

func Panic(msg string) {
	nh := (*reflect.StringHeader)(unsafe.Pointer(&msg))
	hostPanic(uint32(nh.Data), uint32(nh.Len))
}

func queryValueImpl(key TKeyBuilder) (bool, TValue) {
	id := hostQueryValue(uint64(key))
	if id > 0 {
		return true, TValue(id)
	} else {
		return false, TValue(0)
	}
}

func getValueImpl(key TKeyBuilder) TValue {
	return TValue(hostGetValue(uint64(key)))
}

func updateValueImpl(key TKeyBuilder, existingValue TValue) TIntent {
	return TIntent(hostUpdateValue(uint64(key), uint64(existingValue)))
}

func newValueImpl(key TKeyBuilder) TIntent {
	return TIntent(hostNewValue(uint64(key)))
}

func readValuesImpl(key TKeyBuilder, callback func(key TKey, value TValue)) {
	currentReadCallback = callback
	hostReadValues(uint64(key))
}

var currentReadCallback func(key TKey, value TValue)

//export WasmOnReadValue
func onReadValue(key, value uint64) {
	currentReadCallback(TKey(key), TValue(value))
}

//export HostReadValues
func hostReadValues(keyId uint64)

//export HostGetValue
func hostGetValue(keyId uint64) (result uint64)

/*
	returns 0 when not exists
*/
//export HostQueryValue
func hostQueryValue(keyId uint64) (result uint64)

//export HostNewValue
func hostNewValue(keyId uint64) uint64

//export HostUpdateValue
func hostUpdateValue(keyId uint64, existingValueId uint64) uint64

//export WasmAbiVersion_0_0_1
func proxyABIVersion() {
}

var ms runtime.MemStats

//export WasmGetHeapInuse
func getHeapInuse() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapInuse
}

//export WasmGetMallocs
func getMallocs() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Mallocs
}

//export WasmGetFrees
func getFrees() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Frees
}

//export WasmGetHeapSys
func getHeapSys() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapSys
}

//export WasmGC
func gc() {
	runtime.GC()
}

//export HostPanic
func hostPanic(msgPtr, msgSize uint32)

//export HostRowWriterPutString
func hostRowWriterPutString(id uint64, typ uint32, namePtr, nameSize, valuePtr, valueSize uint32)

//export HostRowWriterPutBytes
func hostRowWriterPutBytes(id uint64, typ uint32, namePtr, nameSize, valuePtr, valueSize uint32)

//export HostRowWriterPutQName
func hostRowWriterPutQName(id uint64, typ uint32, namePtr, nameSize, pkgPtr, pkgSize, entityPtr, entitySize uint32)

//export HostRowWriterPutIntBool
func hostRowWriterPutBool(id uint64, typ uint32, namePtr, nameSize, value uint32)

//export HostRowWriterPutInt32
func hostRowWriterPutInt32(id uint64, typ uint32, namePtr, nameSize, value uint32)

//export HostRowWriterPutInt64
func hostRowWriterPutInt64(id uint64, typ uint32, namePtr, nameSize uint32, value uint64)

//export HostRowWriterPutFloat32
func hostRowWriterPutFloat32(id uint64, typ uint32, namePtr, nameSize uint32, value float32)

//export HostRowWriterPutFloat64
func hostRowWriterPutFloat64(id uint64, typ uint32, namePtr, nameSize uint32, value float64)
