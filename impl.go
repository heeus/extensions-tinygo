/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import "runtime"

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed")
	}
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

//export WasmInit
func init() {}

//export WasmFInit
func finit() {}

var ms runtime.MemStats

//export GetHeapInuse
func getHeapInuse() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapInuse
}

//export GetMallocs
func getMallocs() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Mallocs
}

//export GetFrees
func getFrees() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Frees
}

//export GetHeapSys
func getHeapSys() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapSys
}

//export GC
func gc() {
	runtime.GC()
}
