/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

var KeyBuilder func(storage, entity string) (b TKeyBuilder) = keyBuilderImpl
var CanExist func(key TKeyBuilder) (exists bool, value TValue) = canExistImpl
var MustExist func(key TKeyBuilder) TValue = mustExistImpl
var UpdateValue func(key TKeyBuilder, existingValue TValue) TIntent = updateValueImpl
var NewValue func(key TKeyBuilder) TIntent = newValueImpl

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
