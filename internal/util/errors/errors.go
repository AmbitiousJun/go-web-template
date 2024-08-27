package errors

import "errors"

// Wrap 对 errors.Join 的简易封装
func Wrap(old error, msg string) error {
	if old == nil {
		return errors.New(msg)
	}
	return errors.Join(old, errors.New(msg))
}
