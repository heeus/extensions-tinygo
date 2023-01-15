/*
* Copyright (c) 2023-present unTill Pro, Ltd.
*  @author Michael Saigachenko
 */

package extensions

var KeyBuilder func(storage, entity string) (b IKeyBuilder) = keyBuilderImpl
var CanExist func(key IKeyBuilder) (exists bool, value IValue) = canExistImpl
var MustExist func(key IKeyBuilder) IValue = mustExistImpl
var UpdateValue func(key IKeyBuilder, existingValue IValue) IIntent = updateValueImpl
var NewValue func(key IKeyBuilder) IIntent = newValueImpl

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
	PutValue(name string) IValueBuilder // throws panic if field is not an object or array

	SetInt32(index int, value int32)
	SetString(index int, value string)
	SetValue(index int) IValueBuilder // throws panic if field is not an object or array
}

type IIntent interface {
	IValueBuilder
}
