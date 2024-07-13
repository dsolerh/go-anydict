package anydict

import "fmt"

// some util constants to avoid repetition
const (
	_PROP_NOT_PRESENT = "prop %s is not present in the dict"
	_PROP_NOT_OF_TYPE = "prop %s is not of type %T"
	_CONVERT_FROM_TO  = "cannot be safely converted from '%T' to '%T'"
)

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
