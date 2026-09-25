package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

/*
	Hàm này dùng để khởi tạo và kiểm tra kết nối đến Elasticsearch, trả về client nếu thành công hoặc lỗi nếu có vấn đề gì đó xảy ra
*/
func NewESClient(address string) (*elasticsearch.Client, error) {
	
	// elasticsearch.Config: cho phép cấu hình nhiều thứ: retry, timeout, auth,...
	cfg := elasticsearch.Config{
		Addresses: []string{address},
	}

	// elasticsearch.NewClient(): Tạo client mới nhưng chưa thực sự kết nối
	// Nó chỉ validate config và chuẩn bị HTTP transport, nhưng không ping ngay đến ES
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("Lỗi tạo ES client: %v", err)
	}

	// Kiểm tra kết nối bằng cách ping đến ES
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	res, err := client.Ping(client.Ping.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("Lỗi ping ES: %v", err)
	}
	defer res.Body.Close()

	if res.IsError(){
		return  nil, fmt.Errorf("Ping ES trả về lỗi: %s", res.String())
	}

	return client, nil
}