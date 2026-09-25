package main

import (
	"fmt"
	"strings"
	"time"
)

// Hàm đọc input và trim space
func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Bảng menu chức năng
func Menu_STRING() {
	for {

		fmt.Print("Chọn chức năng: ")
		fmt.Println("1. Nhập key - value - exp")
		fmt.Println("2. Đọc giá trị key")
		fmt.Println("3. Xóa key")
		fmt.Println("4. Thoát")
		choice, _ := reader.ReadByte()
		reader.ReadString('\n') // Đọc bỏ ký tự newline sau khi đọc byte

		switch choice {
		case '1':
			InsertKeyAndValue()
		case '2':
			key := readInput("Nhập key cần đọc: ")
			GetKey(key)
		case '3':
			key := readInput("Nhập key cần xóa: ")
			GetKey(key)
			DeleteKey(key)
		default:
			fmt.Println("Thoát chương trình")
			return
		}
	}
}

// Hàm nhập key - value - exp
func InsertKeyAndValue() {
	key := readInput("Nhập key: ")
	value := readInput("Nhập value: ")
	expStr := readInput("Nhập thời hạn key (vd: 10s, 5m, 1h): ")

	exp, err := time.ParseDuration(expStr)
	if err != nil {
		fmt.Println("❌ Lỗi định dạng thời gian:", err)
		return
	}

	SetKey_Value(key, value, exp)
}

// Hàm set giá trị key & value
func SetKey_Value(key, value string, exp time.Duration) error {
	err := Rdb.Set(Ctx, key, value, exp).Err()
	if err != nil {
		fmt.Println("Error setting key-value:", err)
		return err
	}

	fmt.Printf("Key '%s' set with value '%s' and expiration %v\n", key, value, exp)
	return nil
}

// hàm đọc giá trị key
func GetKey(key string) (string, error) {
	val, err := Rdb.Get(Ctx, key).Result()
	if err != nil {
		fmt.Printf("Error getting key '%s': %v\n", key, err)
		return "", err
	}
	fmt.Printf("Value for key '%s': %s\n", key, val)
	return val, nil
}

// Hàm xóa key
func DeleteKey(key string) error {
	if err := Rdb.Del(Ctx, key).Err(); err != nil {
		fmt.Println("Error deleting key:", err)
		return err
	}
	fmt.Printf("Key '%s' deleted successfully\n", key)
	return nil
}
