package respositories

import (
	"context"
	"errors"
	model "golang-mongodb-crud/Model"
	"golang-mongodb-crud/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection name
const CollectionName = "students"

// StudentRepository chứa các hàm thao tac với collection students
type StudentRepository struct {
	collection *mongo.Collection
}

// NewStudentRepository khởi tạo Instance của StudentRepository
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		collection: config.GetCollection(CollectionName),
	}
}

// CREATE
func (r *StudentRepository) Create(input model.CreateStudentInput) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Tạo đối tương nhận đầu vào
	now := time.Now()
	student := model.Student{
		ID:        primitive.NewObjectID(), // MongoDB tự tạo ID dạng ObjectID
		Name:      input.Name,
		Email:     input.Email,
		PhoneNum:  input.PhoneNum,
		GPA:       input.GPA,
		CreatedAt: now,
		UpdatedAt: now,
	}

	//Thực hiện Insert vào MongoDB
	_, err := r.collection.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// READ ALL
func (r *StudentRepository) GetAll() ([]model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var students []model.Student

	// bson.M{} là filter rỗng, tức là lấy tất cả document trong collection
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // Đảm bảo cursor được đóng sau khi sử dụng

	//Duyệt qua từng kết quả trả về và giải mã vào struct Student
	if err = cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	//Trả về slice rỗng thay vì nil
	if students == nil {
		students = []model.Student{}
	}

	return students, nil
}

// READ BY ID
func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Chuyển ID từ string sang ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID không hợp lệ")
	}

	//Tìm document dựa trên ID
	var student model.Student

	filter := bson.M{"_id": objectID} // MongoDB lưu ID dưới dạng _id, nên filter phải dùng _id thay vì id
	err = r.collection.FindOne(ctx, filter).Decode(&student)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Không tìm thấy sinh viên với ID đã cho")
		}
		return nil, err
	}

	return &student, nil
}

// UPDATE
func (r *StudentRepository) Update(id string, input model.UpdateStudentInput) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID không hợp lệ")
	}

	filter := bson.M{"_id": objectID}

	// $set: để cập nhật các trường mới, $currentDate để cập nhật trường UpdatedAt với thời gian hiện tại
	update := bson.M{
		"$set": bson.M{
			"name":      input.Name,
			"email":     input.Email,
			"phoneNum":  input.PhoneNum,
			"gpa":       input.GPA,
			"updated_at": time.Now(),
		},
	}

	// Cập nhật và trả về document sau khi cập nhật
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After) // Tìm document dựa trên ID và cập nhật

	var updateStudent model.Student
	err = r.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updateStudent) // Nếu không tìm thấy document nào để cập nhật, trả về lỗi
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Không tìm thấy sinh viên với ID đã cho")
		}
		return nil, err
	}

	return &updateStudent, nil
}

// DELETE
func (r *StudentRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID không hợp lệ")
	}

	filter := bson.M{"_id": objectID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Nếu DeletedCount là 0 nghĩa là không tìm thấy document nào để xóa, trả về lỗi
	if result.DeletedCount == 0 {
		return errors.New("Không tìm thấy sinh viên với ID đã cho")
	}

	return nil
}
