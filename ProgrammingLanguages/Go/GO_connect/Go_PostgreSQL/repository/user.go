package repository

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/lib/pq"
)

// Định nghĩa cấu trúc dữ liệu cho bảng "users"
type User struct {
    ID        int       `db:"id"`
    Username  string    `db:"username"`
    Email     string    `db:"email"`
    Password  string    `db:"password"`
    FullName  string    `db:"full_name"`
    CreatedAt time.Time `db:"created_at"`
}

// InsertUser thêm một người dùng mới vào cơ sở dữ liệu
func InsertUser(user User, db *sql.DB) error{
	query:= `Insert Into users (username, email, password, full_name) Values($1, $2, $3, $4)`

	_, err := db.Exec(query, user.Username, user.Email, user.Password, user.FullName)
	if err != nil{
		log.Println("Lỗi khi thêm người dùng: ", err)
		return err
	}

	log.Println("Thêm người dùng thành công!")
	return nil
}

// ListUsers lấy danh sách tất cả người dùng từ cơ sở dữ liệu
func ListUsers(db *sql.DB) ([]User, error){
	query := `Select id, username, email, password, full_name, created_at From users`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Lỗi khi truy vấn người dùng: ", err)
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName, &user.CreatedAt)

		if err != nil{
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

// Lấy người dùng theo ID
func GetUserByID(id string, db *sql.DB) (User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	
	var user User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName, &user.CreatedAt)

	if err != nil{
		log.Println("Lỗi khi lấy người dùng theo ID: ", err)
		return User{}, err
	}

	return user, nil
}

// Xóa người dùng theo ID
func DeleteUserByID(id string, db *sql.DB) error{
	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Lỗi khi xóa người dùng theo ID: ", err)
		return err
	}
	log.Println("Xóa người dùng thành công!")
	return nil
}

// Cập nhật thông tin người dùng theo ID
func UpdateUserByID(id string, user User, db *sql.DB) error{
	query := `UPDATE users SET username = $1, email = $2,` +
		`password = $3, full_name = $4 WHERE id = $5`;

	_, err := db.Exec(query, user.Username, user.Email, user.Password, user.FullName, id)
	if err != nil {
		log.Println("Lỗi khi cập nhật người dùng theo ID: ", err)
		return err
	}
	log.Println("Cập nhật người dùng thành công!")
	return nil
}