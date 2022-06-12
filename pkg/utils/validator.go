package utils

import (
	"context"
	"vatansoft-sms-service/pkg/validation"
)

func ValidateRequestWithCtx(c context.Context, r interface{}) map[string]string {
	v := c.Value(Validator).(validation.IValidator)
	if errs := v.Validate(r); len(errs) > 0 {
		return errs
	}

	return nil
}
