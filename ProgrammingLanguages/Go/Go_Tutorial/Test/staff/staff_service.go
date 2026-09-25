package staff

import (
	"fmt"
	"hello/department"
	"hello/utils"
)

func AddStaff() {
 
	fmt.Println("\n###1. Thêm nhân viên")
	id := utils.GetConvertedInt("Nhập id: ")
	name := utils.ReadNonEmptyInput("Nhập tên nhân viên: ")
	DayOfBirth := utils.ReadNonEmptyInput("Nhập ngày sinh (dd/mm/yyyy): ")
	Salary := utils.GetConvertedFloat("Nhập thu nhập của nhân viên: ")
	ID_department := utils.GetConvertedInt("Nhập id bộ phận: ")

	//Kiểm tra xem ID bộ phận có tồn tại hay không
	if department.CheckDuplicateID(ID_department) == false{
		fmt.Println("❌ID bộ phận không tồn tại")
		return
	}

	emp := Staff{
		id:           id,
		name:         name,
		DayOfBirth: DayOfBirth,
		salary: Salary,
		id_department: ID_department,
	}

	Staffs = append(Staffs, emp)
	fmt.Printf("✅Đã thêm nhân viên thành công: %v\n", emp.GetInfo())
}

func DeleteStaff() {
	fmt.Println("\n###2. Xóa nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên cần xóa: ")

	//Kiểm tra xem mã có tồn tại thì xóa
	_, idx,found := FindByID_Unordered(id)
	if found{
		Staffs[idx] = Staffs[len(Staffs) - 1]	//Hoán đổi vị trí muốn xóa với vị trí cuối
		Staffs = Staffs[:len(Staffs) - 1]	//Loại bỏ đi một phần tử ở cuối

		fmt.Println("✅Đã xóa thành công nhân viên có mã: ",id)
	}else{
		fmt.Println("❌Không tìm thấy nhân viên có mã: ",id)
	}
}

func EditStaff() {
	fmt.Println("\n###3. Chỉnh sửa nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên cần chỉnh sửa: ")

	obj, idx, found := FindByID_Unordered(id)
	
	if found{
		obj.name = utils.ReadInput("Nhập lại tên nhân viên: ")
		obj.DayOfBirth = utils.ReadInput("Nhập lại ngày sinh (dd/mm/yyyy): ")
		obj.salary = utils.GetConvertedFloat("Nhập lại thu nhập của nhân viên: ")
		obj.id_department = utils.GetConvertedInt("Nhập lại id bộ phận: ")

		Staffs[idx] = *obj //Cập nhật lại thông tin vào vị trí cũ

		fmt.Printf("✅Đã chỉnh sửa thành công nhân viên có mã %d: %s\n", id, obj.GetInfo())
	}else{
		fmt.Printf("❌Không tìm thấy nhân viên có mã %d\n", id)
	}
}

func SearchStaff() {
	fmt.Println("\n###5. Tìm kiếm nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên: ")
	department, _,found := FindByID_Unordered(id)
	
	if found{
		fmt.Printf("✅Nhân viên có id %d: %s\n", id, department.GetInfo())
	}else{
		fmt.Printf("❌Không tìm thấy nhân viên có id %d\n", id)
	}
}

func ListStaff() {
	fmt.Println("\n###4. Hiển thị danh sách nhân viên")

	if len(Staffs) == 0{
		fmt.Println("Không có nhân viên nào trong danh sách")
	}else{
		for _, staff := range Staffs{
			fmt.Println(staff.GetInfo())
		}
	}
}

//Kiểm tra ID trùng lặp
func CheckDuplicateID(id int) bool {
	_, _, found := FindByID_Unordered(id)
	return found
}

func StaffManagement() {
	for {

		utils.ClearScreen()

		// -- Giao diện quản lý nhân viên --
		fmt.Println("\tQuan ly nhan vien")
		fmt.Println("1. Thêm nhân viên")
		fmt.Println("2. Xóa nhân viên")
		fmt.Println("3. Sửa nhân viên")
		fmt.Println("4. Hiển thị nhân viên")
		fmt.Println("5. Tìm kiếm nhân viên")
		fmt.Println("6. Tính lương trung bình của nhân viên")
		fmt.Println("0. Quay lại menu chính")

		choice := utils.GetConvertedInt("Chon chức năng: ")

		switch choice {
		case 1:
			AddStaff()
		case 2:
			DeleteStaff()
		case 3:
			EditStaff()
		case 4:
			ListStaff()
		case 5:
			SearchStaff()
		case 6:
			fmt.Printf("Lương trung bình của nhân viên: %.2f\n", AverageEmployeeSalary())
		case 0:
			fmt.Println("Quay lại menu chính")
			return
		default:
			fmt.Println("==> Lua chon khong hop le")
		}

		utils.ReadInput("Nhấn Enter để tiếp tục...")

	}

}
