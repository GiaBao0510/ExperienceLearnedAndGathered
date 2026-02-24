package staff

import (
	"fmt"
	"sort"
)

type Staff struct {
	id         int
	name       string
	DayOfBirth string
	salary     float64
	id_department int
}

var Staffs []Staff

func (obj Staff) GetInfo() string {
	return fmt.Sprintf("id: %d; Name: %s; DayOfBirth: %s; Salary: %.2f; id_department: %d", obj.id, obj.name, obj.DayOfBirth, obj.salary, obj.id_department)
}

//Tìm nhân viên dụa trên ID
func FindByID_Orderly(id int) (*Staff, bool) {
	left, right := 0, len(Staffs)-1

	for left <= right {
		mid := left + (right-left)/2

		if Staffs[mid].id == id {
			return &Staffs[mid], true

		} else if Staffs[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return nil, false
}

//Hàm lấy nhân viên theo ID (theo dựa trên danh sách không theo thứ tự)
func FindByID_Unordered(id int) (*Staff, int, bool){
	left, right := 0, len(Staffs) -1

	for left <= right{
		mid := (left + right) /2

		if(Staffs[mid].id == id){
			return  &Staffs[mid], mid, true
		
		}else if(Staffs[left].id == id){
			return  &Staffs[left], left, true
		
		}else if(Staffs[right].id == id){
			return  &Staffs[right], right, true
		
		}else{
			left++
			right--
		}

	}
	return nil, -1, false
}

//Sắp xếp nhân viên theo ID
func SortStaff(){
	sort.Slice(Staffs, func(i, j int) bool{
		return Staffs[i].id < Staffs[j].id
	})
}

// Tính lương trung bình của nhân viên
func AverageEmployeeSalary() float64{
	if len(Staffs) == 0{
		fmt.Println("❌Hiện tại danh sách trống")
		return -1
	}
	var total float64 = 0
	for _, e := range Staffs{
		total += e.salary
	}

	return total/float64(len(Staffs))
}


