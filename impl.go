/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

func canExistImpl(key IKeyBuilder) (bool, Value) {
	id := hostCanExist(uint64(key.(keyBuilder)))
	if id > 0 {
		return true, Value(id)
	} else {
		return false, Value(0)
	}
}

func mustExistImpl(key IKeyBuilder) Value {
	return Value(hostMustExist(uint64(key.(keyBuilder))))
}

func updateValueImpl(key IKeyBuilder, existingValue IValue) IIntent {
	return intent(hostUpdateValue(uint64(key.(keyBuilder)), uint64(existingValue.(Value))))
}

func newValueImpl(key IKeyBuilder) IIntent {
	return intent(hostNewValue(uint64(key.(keyBuilder))))
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
