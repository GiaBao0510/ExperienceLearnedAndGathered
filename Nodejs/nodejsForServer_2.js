//Tải 1 mô-đun, tương tụ có thẻ dùng để import thư viện
const http = require('http');  //Biến http được trả về bởi require('http') có thể được sử dụng để thực hiện gọi tới APIHTTP  

const { ppid } = require('process');
const server = http.createServer(); //Tạo ra một đối tượng Server

//Hàm on() tương đương với hàm addEventListener()
//Sự kiện yêu cầu được phát ra mỗi khi có một yêu cầu HTTP cho chương trình Nodejs để xử lý
// + req là tham số dùng để cung cấp thông tin về yêu cầu
// * res là tham số phản hồi mà chúng ta ghi vào
server.on('request', function(req, res){    
    res.statusCode = 200;                           //Đặt mã trạng thái http
    res.setHeader('Content-Type', 'text/plain');    //Đạt tiêu đề HTTP
    res.end('Hello World\n');                       // Viết tin nhắn để phản hồi sau đó báo hiện cho máy chủ là tin nhắn đã hoàn thành
});


server.on('listening', function(){
    console.log('Server running!');
});

//Hàm sẽ khiến chương trình bắt đầu chấp nhận tin nhắn được gửi đến số cổng nhất định
server.listen(3000);

