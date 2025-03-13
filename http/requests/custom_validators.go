package requests

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/OdairPianta/julia/config"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("exists", CustomExistsValidator)
		v.RegisterValidation("not_exists", CustomNotExistsValidator)
		v.RegisterValidation("exists_or_null", CustomExistsOrNullValidator)
		v.RegisterValidation("phone_number", CustomPhoneNumberValidator)
		v.RegisterValidation("required_without", CustomRequiredWithoutValidator)
	} else {
		panic("Error getting validator CustomExistsValidator")
	}
}

func CustomExistsValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	params := strings.Split(param, ".")
	if len(params) != 2 {
		return false
	}

	table := params[0]
	column := params[1]
	var value string

	field := fl.Field()
	switch field.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = fmt.Sprintf("%d", field.Uint())
	case reflect.String:
		value = field.String()
	default:
		return false
	}

	var count int64
	err := config.DB.Table(table).Where(fmt.Sprintf("%s = ?", column), value).Count(&count).Error
	if err != nil {
		fmt.Println("Error when run custom exists or null validator: ", err)
		return false
	}

	return count > 0
}

func CustomNotExistsValidator(fl validator.FieldLevel) bool {
	return !CustomExistsValidator(fl)
}

func CustomExistsOrNullValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	params := strings.Split(param, ".")
	if len(params) != 2 {
		return false
	}

	table := params[0]
	column := params[1]
	var value string

	field := fl.Field()
	if field.String() == "" || field.IsZero() {
		return true
	}

	switch field.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = fmt.Sprintf("%d", field.Uint())
	case reflect.String:
		value = field.String()
	default:
		return false
	}

	var count int64
	err := config.DB.Table(table).Where(fmt.Sprintf("%s = ?", column), value).Count(&count).Error
	if err != nil {
		fmt.Println("Error when run custom exists or null validator: ", err)
		return false
	}

	return count > 0
}

func CustomPhoneNumberValidator(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.String() == "" || field.IsZero() {
		return true
	}

	phone := field.String()
	regex := `^\+\d{2}\d{2}\d{4,5}\d{4}$`
	matched := regexp.MustCompile(regex).MatchString(phone)
	return matched
}

func CustomRequiredWithoutValidator(fl validator.FieldLevel) bool {
	param := fl.Param()
	otherField := fl.Parent().FieldByName(param)
	if otherField.String() != "" {
		return true
	}

	field := fl.Field()
	if field.String() == "" {
		return false
	}

	return true
}
