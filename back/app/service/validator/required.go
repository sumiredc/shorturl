package validator

import "errors"

func (v *Validator) Required(value string) error {
	if value == "" {
		return errors.New("This field is required.")
	}

	return nil
}
