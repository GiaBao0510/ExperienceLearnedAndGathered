package department

import (
	"fmt"
	"hello/utils"
	"time"
)

func AddDepartment() {

	fmt.Println("\n###1. Thêm bộ phận")
	id := utils.GetConvertedInt("Nhập id: ")
	name := utils.ReadNonEmptyInput("Nhập tên bộ phận: ")
	creationDate := utils.ReadNonEmptyInput("Nhập ngày tạo (dd/mm/yyyy): ")

	department := Department{
		id:           id,
		name:         name,
		creationDate: creationDate,
	}

	Departments = append(Departments, department)
	fmt.Printf("Đã thêm bộ phận thành công: %v\n", department.GetInfo())
}

// Xóa dựa trên ID
func DeleteDepartment() {
	fmt.Println("\n###2. Xóa bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận cần xóa: ")

	//Kiểm tra xem mã có tồn tại thì xóa
	_, idx, found := FindByID_Unordered(id)
	if found {
		Departments[idx] = Departments[len(Departments)-1] //Hoán đổi vị trí muốn xóa với vị trí cuối
		Departments = Departments[:len(Departments)-1]     //Loại bỏ đi một phần tử ở cuối

		fmt.Println("✅Đã xóa thành công bộ phận có mã: ", id)
	} else {
		fmt.Println("❌Không tìm thấy bộ phận có mã: ", id)
	}
}

// hàm chỉnh sửa dựa trên ID
func EditDepartment() {
	fmt.Println("\n###3. Chỉnh sửa bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận cần chỉnh sửa: ")

	depart, idx, found := FindByID_Unordered(id)

	if found {
		depart.name = utils.ReadInput("Nhập lại tên bộ phận: ")
		depart.creationDate = time.Now().Local().Format("02/01/2006") //Tự động câp nhật ngày chỉnh sửa

		Departments[idx] = *depart //Cập nhật lại thông tin vào vị trí cũ

		fmt.Printf("✅Đã chỉnh sửa thành công bộ phận có mã %d: %s\n", id, depart.GetInfo())
	} else {
		fmt.Printf("❌Không tìm thấy bộ phận có mã %d\n", id)
	}
}

func SearchDepartment() {
	fmt.Println("\n###5. Tìm kiếm bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận: ")
	department, _, found := FindByID_Unordered(id)

	if found {
		fmt.Printf("✅Bộ phận có id %d: %s\n", id, department.GetInfo())
	} else {
		fmt.Printf("❌Không tìm thấy bộ phận có id %d\n", id)
	}
}

// Kiểm tra ID trùng lặp
func CheckDuplicateID(id int) bool {
	_, _, found := FindByID_Unordered(id)
	return found
}

func ListDepartment() {
	fmt.Println("\n###4. Hiển thị danh sách bộ phận")

	if len(Departments) == 0 {
		fmt.Println("Không có bộ phân nào trong danh sách")
	} else {
		for _, department := range Departments {
			fmt.Println(department.GetInfo())
		}
	}
}

func DepartmentManagement() {
	for {

		utils.ClearScreen()

		// -- Giao diện quản lý bộ phận --
		fmt.Println("\tQuan ly bo phan")
		fmt.Println("1. Thêm bộ phận")
		fmt.Println("2. Xóa bộ phận")
		fmt.Println("3. Sửa bộ phận")
		fmt.Println("4. Hiển thị bộ phận")
		fmt.Println("5. Tìm kiếm bộ phân")
		fmt.Println("0. Quay lại menu chính")

		choice := utils.GetConvertedInt("Chon chức năng: ")

		switch choice {
		case 1:
			AddDepartment()
		case 2:
			DeleteDepartment()
		case 3:
			EditDepartment()
		case 4:
			ListDepartment()
		case 5:
			SearchDepartment()
		case 0:
			fmt.Println("Quay lại menu chính")
			return
		default:
			fmt.Println("==> Lua chon khong hop le")
		}

		utils.ReadInput("Nhấn Enter để tiếp tục...")

	}

}
