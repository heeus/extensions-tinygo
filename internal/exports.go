package extengine

import "runtime"

var ms runtime.MemStats

//export WasmGetHeapInuse
func GetHeapInuse() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapInuse
}

//export WasmGetMallocs
func GetMallocs() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Mallocs
}

//export WasmGetFrees
func GetFrees() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.Frees
}

//export WasmGetHeapSys
func GetHeapSys() uint64 {
	runtime.ReadMemStats(&ms)
	return ms.HeapSys
}

//export WasmGC
func Gc() {
	runtime.GC()
}
