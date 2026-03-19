Đầu tiên chúng ta nên nhớ là hàm main() là hàm khởi tạo của hàm Go

Ví dụ trong một dự án chúng ta thường viết hàm main() như sau:
```go
func main() {
	r := routers.NewRouter()

	InitMysql()
	InitRedis()
	InitKafka()
	InitElasticSearch()

	r.Run(":8080")
}
```

Nếu như trong dự án của một công ty không nhỏ, thì chúng ta không nên thiết kế như này trong hàm main. Vì hàm main là một phương thức đặc biệt. Nhiệm vụ của hàm này chủ yếu là Run mà thôi. Còn những init như thế này thì phải rút đi ở chỗ khác hoặc nhóm lại ở func khác trong một package khác