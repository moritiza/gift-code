package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Set custom validators
var (
	validateMobile validator.Func = func(fl validator.FieldLevel) bool {
		mobile, err := regexp.Compile(`^09[0-9]{9}$`)
		if err != nil {
			return false
		}
		if mobile.MatchString(fl.Field().String()) {
			return true
		}
		return false
	}

	validateCode validator.Func = func(fl validator.FieldLevel) bool {
		code, err := regexp.Compile(`^[a-zA-Z0-9]{6,}$`)
		if err != nil {
			return false
		}
		if code.MatchString(fl.Field().String()) {
			return true
		}
		return false
	}

	validateCodeCredit validator.Func = func(fl validator.FieldLevel) bool {
		codeCredit, err := regexp.Compile(`^[1-9]{1}[0-9]+$`)
		if err != nil {
			return false
		}
		if codeCredit.MatchString(fmt.Sprint(fl.Field())) {
			return true
		}
		return false
	}

	validateCount validator.Func = func(fl validator.FieldLevel) bool {
		count, err := regexp.Compile(`^[1-9]{1}[0-9]+$`)
		if err != nil {
			return false
		}
		if count.MatchString(fmt.Sprint(fl.Field())) {
			return true
		}
		return false
	}
)

// Validator create new validator with en language
func Validator() *validator.Validate {
	v := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	v.RegisterValidation("mobile", validateMobile)
	v.RegisterValidation("code", validateCode)
	v.RegisterValidation("code_credit", validateCodeCredit)
	v.RegisterValidation("count", validateCount)

	return v
}

// ValidatorErrors return validator errors and translate them to en language
func ValidatorErrors(v *validator.Validate, err error) string {
	var errString string

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validatorCustomErrors(v, trans)

	// Get all validator errors
	errs := err.(validator.ValidationErrors)

	// Translate each validator error
	for _, e := range errs {
		errString += e.Translate(trans) + "\n"
	}

	return strings.TrimSuffix(errString, "\n")
}

// validatorCustomErrors set custom error message for validate tags
func validatorCustomErrors(v *validator.Validate, trans ut.Translator) {
	v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	v.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0} not a valid mobile!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())

		return t
	})

	v.RegisterTranslation("code", trans, func(ut ut.Translator) error {
		return ut.Add("code", "{0} not a valid discount code!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("code", fe.Field())

		return t
	})

	v.RegisterTranslation("code_credit", trans, func(ut ut.Translator) error {
		return ut.Add("code_credit", "{0} not a valid code credit!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("code_credit", fe.Field())

		return t
	})

	v.RegisterTranslation("count", trans, func(ut ut.Translator) error {
		return ut.Add("count", "{0} not a valid count!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("count", fe.Field())

		return t
	})
}
