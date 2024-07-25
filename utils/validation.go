package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

// Validator instance
var validate = validator.New()

// Register custom validations
func init() {
	validate.RegisterValidation("mobilephone", validateMobilePhone)
	validate.RegisterValidation("birthdate", validateBirthdate)
}

// ValidateStruct validates a struct based on tags
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// Custom validation for mobile phone format
func validateMobilePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	re := regexp.MustCompile(`^\+?(\d{1,3})?[-.\s]?(\d{10,13})$`)
	return re.MatchString(phone)
}

// Custom validation for birthdate (YYYY-MM-DD format)
func validateBirthdate(fl validator.FieldLevel) bool {
	birthdate := fl.Field().String()
	_, err := time.Parse("2006-01-02", birthdate)
	return err == nil
}
