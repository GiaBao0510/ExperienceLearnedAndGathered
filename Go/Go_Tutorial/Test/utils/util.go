package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

//hàm đọc thông tin đầu vào
func ReadInput(input string) string{

	fmt.Printf("%s ", input)
	reader := bufio.NewReader(os.Stdin) //Tạo một đối tượng reader để đọc dữ liệu từ bàn phím
	value, _ := reader.ReadString('\n') //Đọc một dòng dữ liệu từ bàn phím, kết thúc khi người dùng nhấn Enter
	return strings.TrimSpace(value)
}

//Hàm đọc chuỗi không được để trống
func ReadNonEmptyInput(input string) string{
	for{
		value := ReadInput(input)

		if !IsEmpty(value){
			return value
		}

		fmt.Println("❌Giá trị không được để trống")
	}
}

//Ham kiểm tra xem chuỗi có rỗng không
func IsEmpty(value string) bool{
	return value == "" || len(value) == 0
}

//Hàm chuyển số đầu vào thành số nguyên
func GetConvertedInt(promp string) int{
	for{

		input := ReadInput(promp)

		//Nếu có lỗi là rỗng thì yêu cầu nhập lại 
		if IsEmpty(input){
			fmt.Println("❌Giá trị không được để trống")
			continue
		}


		value, err := strconv.Atoi(input)

		//Nếu không có lỗi thì trả về
		if err == nil && value  > -1{
			return value
		}

		fmt.Println("❌Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}

//Hàm chuyển số đầu vào thành số thực
func GetConvertedFloat(prompt string) float64{
	for{

		input := ReadInput(prompt)

		//Nếu có lỗi là rỗng thì yêu cầu nhập lại 
		if IsEmpty(input){
			fmt.Println("❌Giá trị không được để trống")
			continue
		}

		value, err := strconv.ParseFloat(input, 64)

		//Nếu không có lỗi thì trả về
		if err == nil && value  > -1{
			return value
		}

		fmt.Println("❌Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}


//Hàm clear màn hình dựa trên hệ điều hành
func ClearScreen(){
	var cmd *exec.Cmd	//là một struct đại diện cho một lệnh sẽ được thực thi

	//Kiểm tra hệ điều hành hiện tại
	switch runtime.GOOS{
		case "windows":
			cmd = exec.Command("cmd", "/c", "cls")
		
		default: //Mặc định là hệ điều hành Unix/Linux/MacOS
			cmd = exec.Command("clear")
	}

	//Gắn output ra màn hình
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil{
		fmt.Println("Error clearing screen: ",err)
	}
}

//Hàm tạo mã ID ngẫu nhiên
func GenerateID() string{
	return uuid.New().String()
}
