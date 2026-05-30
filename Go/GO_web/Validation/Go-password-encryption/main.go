package main

import (
	"fmt"

	b "github.com/GiaBao0510/Go-password-encryption/GO-bcrypt"
	argon "github.com/GiaBao0510/Go-password-encryption/Go-argon2id"
)

func main(){ 
	
	Password := "mysecretpassword"	// Mật khẩu gốc cần băm
	
	// -------------------------- Bcrypt --------------------------
	// Khởi tạo
	bcrypt := b.NewBcrypt(b.Bcrypt{
		Password: Password,
		Cost: 50,
	})

	// Băm mật khẩu
	hashedPassword_bcrypt, err := bcrypt.HashPassword()


	if err != nil {
		panic(err)
	}
	println("Mật khẩu đã băm bằng thuật toán bcrypt:", hashedPassword_bcrypt)

	// Kiểm tra mật khẩu
	fmt.Println("Kiểm tra mật khẩu qua thuật toán bcrypt")
	isValid := bcrypt.CheckPasswordHash("mysecretpassword", hashedPassword_bcrypt)
	if isValid {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// -------------------------- Argon2id --------------------------
	// khởi tạo
	fmt.Println("\n\tKhởi tạo thuật toán Argon2id")
	argon2id := argon.NewHasher(argon.Argon2Params{
		TimeCost: 3,
		MemoryCost: 64 * 1024, // 64 MB
		Threads: 4,
		KeyLength: 32,
	})

	// Băm mật khẩu
	hashedPassword_argon2id, err := argon2id.Hash(Password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mật khẩu đã băm bằng thuật toán Argon2id:", hashedPassword_argon2id)

	// Kiểm tra mật khẩu
	fmt.Println("Kiểm tra mật khẩu qua thuật toán Argon2id")
	isValid, err = argon2id.Verify(hashedPassword_argon2id, "mysecretpasswordok")
	if err != nil {
		panic(err)
	}
	if isValid {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}