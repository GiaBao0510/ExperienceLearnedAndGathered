package department

import (
	"fmt"
	"sort"
)

type Department struct {
	id           int
	name         string
	creationDate string
}

var Departments []Department

func (Obj Department) GetInfo() string {
	return fmt.Sprintf("id: %d; name: %s; creationDate: %s", Obj.id, Obj.name, Obj.creationDate)
}

// Hàm lấy bộ phận theo ID (theo dựa trên danh sách đã sắp xếp)
func FindByID_Orderly(id int) (*Department, bool) {
	left, right := 0, len(Departments)-1

	for left <= right {
		mid := left + (right-left)/2

		if Departments[mid].id == id {
			return &Departments[mid], true

		} else if Departments[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nil, false
}

// Hàm lấy bộ phận theo ID (theo dựa trên danh sách không theo thứ tự)
func FindByID_Unordered(id int) (*Department, int, bool) {
	left, right := 0, len(Departments)-1

	for left <= right {
		mid := (left + right) / 2

		if Departments[mid].id == id {
			return &Departments[mid], mid, true

		} else if Departments[left].id == id {
			return &Departments[left], left, true

		} else if Departments[right].id == id {
			return &Departments[right], right, true

		} else {
			left++
			right--
		}

	}
	return nil, -1, false
}

// Sắp xếp bộ phận theo ID
func SortDepartment() {
	sort.Slice(Departments, func(i, j int) bool {
		return Departments[i].id < Departments[j].id
	})
}
