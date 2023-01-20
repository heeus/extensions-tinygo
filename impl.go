/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

import "runtime"

func canExistImpl(key TKeyBuilder) (bool, TValue) {
	id := hostCanExist(uint64(key))
	if id > 0 {
		return true, TValue(id)
	} else {
		return false, TValue(0)
	}
}

func mustExistImpl(key TKeyBuilder) TValue {
	return TValue(hostMustExist(uint64(key)))
}

func updateValueImpl(key TKeyBuilder, existingValue TValue) TIntent {
	return TIntent(hostUpdateValue(uint64(key), uint64(existingValue)))
}

func newValueImpl(key TKeyBuilder) TIntent {
	return TIntent(hostNewValue(uint64(key)))
}

//export HostMustExist
func hostMustExist(keyId uint64) (result uint64)

/*
	returns 0 when not exists
*/
//export HostCanExist
func hostCanExist(keyId uint64) (result uint64)

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
