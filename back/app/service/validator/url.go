package validator

import (
	"errors"
	"regexp"
)

func (v *Validator) Url(value string) error {
	pattern := `^https?://[-_.!~*'()a-zA-Z0-9;/?:@&=+$,%#]+`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if !regex.MatchString(value) {
		return errors.New("The URL format is incorrect.")
	}

	return nil
}
