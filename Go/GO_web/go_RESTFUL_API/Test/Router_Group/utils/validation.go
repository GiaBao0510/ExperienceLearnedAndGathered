package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// --- Tạo hàm kiểm tra đầu vào bắt buộc ---
func ValidationRequired(fieldname, value string) error {
	if value == "" {
		return fmt.Errorf("The %s field must not be blank.", fieldname)
	}
	return nil
}

// --- Tạo hàm kiểm tra độ dài tối đa & tối thiểu---
func ValidationLength(fieldname, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return fmt.Errorf("The %s field must be between %d and %d characters long.", fieldname, min, max)
	}
	return nil
}

// --- Tạo hàm kiểm tra regex---
func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error {
	if !reg.MatchString(value) {
		return fmt.Errorf("The %s field is not in the correct format (%s).", fieldname, reg.String())
	}
	return nil
}

// --- Tạo hàm kiểm tra giá trị số nguyên dương ---
func ValidationPositiveInt(fieldname, value string) error {

	v, err := strconv.Atoi(value)

	if err != nil || v <= 0 {
		return fmt.Errorf("The requirement is that %s must be a positive integer.", fieldname)
	}
	return nil
}

// --- Tạo hàm kiểm tra giá trị UUID ---
func ValidationUUID(fieldname, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Field %s : %s is not a valid UUID.", fieldname, uid)
	}

	return uid, nil
}

// --- Tạo hàm kiểm tra giá trị có map hay không ---
func ValidationInList(fieldName, value string, list map[string]bool) error {
	if !list[value] {
		return fmt.Errorf("The value of %s, at this field %s, does not exist in the list. \n"+
			"These are the valid values: %v \n", value, fieldName, GetKeysFromMap(list))
	}
	return nil
}

// Lấy các keys trong map
func GetKeysFromMap(m map[string]bool) []string {

	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

// Viết hàm xử lý lỗi
func HandleValidationErrors(Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr, ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr {

			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ."
			case "slug":
				errors[e.Field()] = e.Field() + " Chỉ có thể chứa: chữ thường, số, dấu gạch ngang (-) hoặc dấu chấm (.)"
			case "max":
				errors[e.Field()] = e.Field() + " Độ dài tối đa là " + e.Param() + " ký tự."
			case "min":
				errors[e.Field()] = e.Field() + " Độ dài tối thiểu là " + e.Param() + " ký tự."
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = e.Field() + " phải là một trong các giá trị sau: " + allowedValues + "."
			}
			log.Printf("Validation error on field '%s': %s, Tag: %+v", e.Field(), e.Error(), e.Tag())
		}

		return gin.H{"error": errors}
	}

	return gin.H{"error": "Yêu cầu không hợp lệ: " + Err.Error()}
}

// Tự tạo một hàm để kiểm tra lỗi và trả về lỗi nếu có
func RegisterValidationError() error {

	//Kiểm tra xem kiểu đầu vào có thuộc kiểu validator không
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Failed to register validation: could not get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)

	//Đăng ký hàm kiểm tra regex cho tag "slug", và kiểm tra field có phải là string hay không, nếu không thì trả về lỗi
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String()) // Kiểm tra xem field này có khớp với regex hay không
	})

	return nil
}
