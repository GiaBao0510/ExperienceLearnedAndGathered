package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

// --- Tạo hàm kiểm tra đầu vào bắt buộc ---
func ValidationRequired(fieldname, value string) error{
	if value == ""{
		return fmt.Errorf("The %s field must not be blank.", fieldname)
	}
	return nil
}


// --- Tạo hàm kiểm tra độ dài tối đa & tối thiểu---
func ValidationLength(fieldname, value string, min, max int) error{
	if len(value) < min || len(value) > max{
		return fmt.Errorf("The %s field must be between %d and %d characters long.", fieldname, min, max)
	}
	return nil
}

// --- Tạo hàm kiểm tra regex---
func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error{
	if !reg.MatchString(value){
		return fmt.Errorf("The %s field is not in the correct format (%s).", fieldname, reg.String())
	}
	return nil
}

// --- Tạo hàm kiểm tra giá trị số nguyên dương ---
func ValidationPositiveInt(fieldname, value string) error{
	
	v, err := strconv.Atoi(value)
	
	if err != nil || v <= 0{
		return fmt.Errorf("The requirement is that %s must be a positive integer.", fieldname)
	}
	return nil
}

// --- Tạo hàm kiểm tra giá trị UUID ---
func ValidationUUID(fieldname, value string) (uuid.UUID, error){
	 uid , err := uuid.Parse(value)
	if err != nil{
		return uuid.Nil, fmt.Errorf("Field %s : %s is not a valid UUID.", fieldname, uid)
	}

	return uid, nil
}

// --- Tạo hàm kiểm tra giá trị có map hay không ---
func ValidationInList(fieldName, value string, list map[string]bool) error{
	if !list[value]{
		return fmt.Errorf("The value of %s, at this field %s, does not exist in the list. \n"+
		"These are the valid values: %v \n",value, fieldName, GetKeysFromMap(list))
	}
	return nil
}


// Lấy các keys trong map
func GetKeysFromMap(m map[string]bool) []string{
	
	keys := make([]string, 0, len(m))
	for key := range m{
		keys = append(keys, key)
	}

	return keys
}