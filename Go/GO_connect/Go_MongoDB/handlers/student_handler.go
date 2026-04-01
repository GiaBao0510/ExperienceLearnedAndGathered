package handlers

import (
	model "golang-mongodb-crud/Model"
	"golang-mongodb-crud/respositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	repo *respositories.StudentRepository
}

// Tạo một instance mới của StudentHandler
func NewStudentHandler() *StudentHandler {
	return &StudentHandler {
		repo: respositories.NewStudentRepository(),
	}
}

//POST để tạo mới một student
func (h *StudentHandler) CreateStudent(c *gin.Context) {

	var input model.CreateStudentInput

	// Đọc dữ liệu và validation dữ liệu từ body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"message": "Dữ liệu không hợp lệ" + err.Error(),
		})
		return
	}

	// Gọi hàm Create từ repository để tạo mới student
	student, err := h.repo.Create(input)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "false",
			"message": "Lỗi khi tạo student: " + err.Error(),
		})
		return
	}

	// Trả về response thành công
	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"message": "Student created successfully",
		"data": student,
	})
}

//GET để lấy tất cả student
func (h *StudentHandler) GetAll(c *gin.Context){
	students, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "false",
			"message": "Lỗi khi lấy danh sách student: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"message": "Danh sách student retrieved successfully",
		"data": students,
	})
}

// GET để lấy một student theo ID
func (h *StudentHandler) GetByID(c *gin.Context){

	id := c.Param("id")
	student, err := h.repo.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "false",
			"message": "Lỗi khi tìm kiếm student: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"message": "Student retrieved successfully",
		"data": student,
	})
}

// PUT để cập nhật một student theo ID
func (h *StudentHandler) UpdateStudent(c *gin.Context) {

	var input model.UpdateStudentInput
	id := c.Param("id")

	// Đọc dữ liệu và validation dữ liệu từ body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"message": "Dữ liệu không hợp lệ" + err.Error(),
		})
		return
	}

	// Gọi hàm Update từ repository để cập nhật student
	student, err := h.repo.Update(id, input)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "false",
			"message": "Lỗi khi cập nhật student: " + err.Error(),
		})
		return
	}

	// Trả về response thành công
	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"message": "Student updated successfully",
		"data": student,
	})
}

// DELETE để xóa một student theo ID
func (h *StudentHandler) DeleteStudent(c *gin.Context) {

	id := c.Param("id")
	err := h.repo.Delete(id)

	if err != nil {
		status := http.StatusInternalServerError
		
		if err.Error() == "Không tìm thấy sinh viên với ID đã cho" || err.Error() == "ID không hợp lệ" {
			status = http.StatusNotFound
		}

		c.JSON(status, gin.H{
			"status": "false",
			"message": "Lỗi khi tìm kiếm student: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "true",
		"message": "Student deleted successfully",
	})
}