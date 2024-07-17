package anydict

import "fmt"

// some util constants to avoid repetition
const (
	_PROP_NOT_PRESENT      = "prop %s is not present in the dict"
	_PROP_NOT_OF_TYPE      = "prop %s is not of type %T"
	_CONVERT_FROM_TO       = "cannot be safely converted from '%T' to '%T'"
	_INVALID_TYPE_AT_INDEX = "invalid type %T at index %d"
)

// PropNotPresentError is used when the prop is not present in the Dict
type PropNotPresentError struct {
	msg string
}

var _ error = (*PropNotPresentError)(nil)

func (e *PropNotPresentError) Error() string {
	return e.msg
}

func newPropNotPresentError(prop string) error {
	return &PropNotPresentError{
		msg: fmt.Sprintf(_PROP_NOT_PRESENT, prop),
	}
}

// PropNotOfTypeError is used when the prop is not of the specified type
// (or cannot be casted to that type)
type PropNotOfTypeError struct {
	msg string
}

var _ error = (*PropNotOfTypeError)(nil)

func (e *PropNotOfTypeError) Error() string {
	return e.msg
}

func newPropNotOfTypeError[T any](prop string, val T) error {
	return &PropNotOfTypeError{
		msg: fmt.Sprintf(_PROP_NOT_OF_TYPE, prop, val),
	}
}

// InvalidConversionError is used when the prop cannot be safely converted
// from his original type to the requested one, like:
//
//	original := int(2345)
//	requested := int8(original)  <- this will fail, cause the value 2345 cannot be represented in int8
type InvalidConversionError struct {
	msg string
}

var _ error = (*InvalidConversionError)(nil)

func (e *InvalidConversionError) Error() string {
	return e.msg
}

func newInvalidConversionError[F, T any](from F, to T) error {
	return &InvalidConversionError{
		msg: fmt.Sprintf(_CONVERT_FROM_TO, from, to),
	}
}

type InvalidTypeAtIndexError struct {
	msg string
}

var _ error = (*InvalidTypeAtIndexError)(nil)

func (e *InvalidTypeAtIndexError) Error() string {
	return e.msg
}

func newInvalidTypeAtIndex[T any](index int, val T) error {
	return &InvalidTypeAtIndexError{
		msg: fmt.Sprintf(_INVALID_TYPE_AT_INDEX, val, index),
	}
}
