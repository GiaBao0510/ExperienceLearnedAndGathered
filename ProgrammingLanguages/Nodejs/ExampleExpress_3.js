const express = require('express'); //Them thu vien express
const app = express();              //Tao doi tuonng App

/*
    Express có các đối tượng Request và Respone:
        + req là đối tượng Request.
        + res là đối tượng Respone.
        + req.send() gửi 1 phản hồi HTTP với giá trị đã cho.
    > Nếu bạn chuyển 1 chuỗi,theo mặc định nó đặt tiêu để Content-Type đến text/html
    > Nếu bạn chuyển 1 đối tượng hoặc mảng, nó đặt tiêu đề kiểu nội dung ứng dụng/json
    và phân tích tham số đó thành JSON.
        * send() tự động đặt tiêu đề phản hồi HTTP có độ dài nội dung.
        * send() cũng tự động ngắt kết nối
    Cú pháp: app.methob(path, handler);
*/

app.get('/',function(req, res){
    res.send('Hello World!');
});

//Lang nghe tren cong 3000
app.listen(3000, function(){
    console.log('Example app listening on post 3000!');
})

