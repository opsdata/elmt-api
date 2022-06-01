package v1

import (
	"github.com/opsdata/common-base/pkg/validation"
	"github.com/opsdata/common-base/pkg/validation/field"
)

// Validate validates that a user object is valid.
func (u *User) Validate() field.ErrorList {
	val := validation.NewValidator(u)
	allErrs := val.Validate()

	if err := validation.IsValidPassword(u.Password); err != nil {
		allErrs = append(allErrs, field.Invalid(field.NewPath("password"), err.Error(), ""))
	}

	return allErrs
}

// ValidateUpdate validates that a user object is valid when update.
// Like User.Validate but not validate password.
func (u *User) ValidateUpdate() field.ErrorList {
	val := validation.NewValidator(u)
	return val.Validate()
}

// Validate validates that a secret object is valid.
func (s *Secret) Validate() field.ErrorList {
	val := validation.NewValidator(s)
	return val.Validate()
}
