package validator

import "errors"

func (v *Validator) StringMax(value string, max int) error {
	if len(value) > max {
		return errors.New("The character limit has been exceeded.")
	}
	return nil
}
