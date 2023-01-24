/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

var KeyBuilder func(storage, entity string) (b TKeyBuilder) = keyBuilderImpl

// QueryValue queries value. When not exists it returns exists=false and value=nil.
var QueryValue func(key TKeyBuilder) (exists bool, value TValue) = queryValueImpl

// GetValue gets value. Panics when value is not exist
var GetValue func(key TKeyBuilder) TValue = getValueImpl

// ReadValues reads using partial key and returns values in callback.
//
// Important: key and value are not kept after callback!
var ReadValues func(key TKeyBuilder, callback func(key TKey, value TValue)) = readValuesImpl

// UpdateValue creates intent to update a value
var UpdateValue func(key TKeyBuilder, existingValue TValue) TIntent = updateValueImpl

// NewValue creates intent for new value
var NewValue func(key TKeyBuilder) TIntent = newValueImpl

type QName struct {
	Pkg    string
	Entity string
}

type IKey interface {
	AsString(name string) string
	AsInt32(name string) int32
	AsInt64(name string) int64
	AsFloat32(name string) float32
	AsFloat64(name string) float64
	AsBytes(name string) []byte
	AsQName(name string) QName
	AsBool(name string) bool
}

type IValue interface {
	Length() uint32

	AsString(name string) string
	AsInt32(name string) int32
	AsValue(name string) IValue // throws panic if field is not an object or array

	GetAsString(index int) string
	GetAsInt32(index int) int32
	GetAsValue(index int) IValue // throws panic if field is not an object or array
}

type IKeyBuilder interface {
	PutInt32(name string, value int32)
	PutString(name string, value string)
}

type IValueBuilder interface {
	PutInt32(name string, value int32)
	PutString(name string, value string)
}

type IIntent interface {
	IValueBuilder
}
