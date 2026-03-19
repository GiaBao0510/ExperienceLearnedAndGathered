package main

import (
	"bufio"
	"fmt"
	"go-postgresql/database"
	"go-postgresql/repository"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	
	// Kết nối DB
	db := database.InitDB()
	defer db.Close()

	//Vòng lặp chọn chức năng
	for {
		fmt.Println("\t\tMenu:")
		fmt.Println("1. Hiển thị danh sách người dùng")
		fmt.Println("2. Thêm người dùng")
		fmt.Println("3. Sửa thông tin người dùng")
		fmt.Println("4. Xóa người dùng")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn chức năng: ")
		choice := bufio.NewScanner(os.Stdin)
		choice.Scan()

		switch choice.Text(){
			case "1":
				fmt.Println("Hiển thị danh sách người dùng")
				users, err:= repository.ListUsers(db)
				if err != nil {
					log.Println("Lỗi khi lấy danh sách người dùng: ", err)
					continue
				}
				for _, user := range users {
					fmt.Printf("ID: %d, Username: %s, Email: %s, Full Name: %s, Created At: %s\n",
						user.ID, user.Username, user.Email, user.FullName, user.CreatedAt.Format("2006-01-02 15:04:05"))
				}
			case "2":
				fmt.Println("Thêm người dùng")
				var user repository.User
				fmt.Print("Nhập username: ")
				fmt.Scanln(&user.Username)
				fmt.Print("Nhập email: ")
				fmt.Scanln(&user.Email)
				fmt.Print("Nhập password: ")
				fmt.Scanln(&user.Password)
				fmt.Print("Nhập full name: ")
				fmt.Scanln(&user.FullName)
				err := repository.InsertUser(user, db)
				if err != nil {
					log.Println("Lỗi khi thêm người dùng: ", err)
				}else{
					fmt.Println("Thêm người dùng thành công!")
				}
			case "3":
				fmt.Println("Sửa thông tin người dùng")
				fmt.Print("Nhập ID người dùng cần sửa: ")
				var id string
				fmt.Scanln(&id)
				user, err := repository.GetUserByID(id, db)
				if err != nil {
					log.Println("Lỗi khi lấy thông tin người dùng: ", err)
					continue
				}
				fmt.Printf("Thông tin hiện tại:\nID: %d, Username: %s, Email: %s, Full Name: %s\n", user.ID, user.Username, user.Email, user.FullName)
				fmt.Print("Nhập username mới: ")
				fmt.Scanln(&user.Username)
				fmt.Print("Nhập email mới: ")
				fmt.Scanln(&user.Email)
				fmt.Print("Nhập password mới: ")
				fmt.Scanln(&user.Password)
				fmt.Print("Nhập full name mới: ")
				fmt.Scanln(&user.FullName)
				err = repository.UpdateUserByID(id, user, db)
				if err != nil {
					log.Println("Lỗi khi cập nhật người dùng: ", err)
				} else {
					fmt.Println("Cập nhật người dùng thành công!")
				}
			case "4":
				fmt.Println("Xóa người dùng")
				fmt.Print("Nhập ID người dùng cần xóa: ")
				var id string
				fmt.Scanln(&id)
				err := repository.DeleteUserByID(id, db)
				if err != nil {
					log.Println("Lỗi khi xóa người dùng: ", err)
				}else{
					fmt.Println("Xóa người dùng thành công!")
				}
			case "0":
				fmt.Println("Thoát")
				return
			default:
				fmt.Println("Lựa chọn không hợp lệ")
		}
	}
}