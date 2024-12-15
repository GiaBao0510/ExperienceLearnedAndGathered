/*
    Cú pháp: function(req, res, next){...};
    - mọi hàm trong ngăn xếp cần có 3 đối số:
        > 2 tham số đầu là: req (là đối tượng yêu cầu),
        res (là đối tượng phản hồi).
        > Tham số thứ 3:
            + Bản thân nó là 1 hàm
            + Được gọi, Express sẽ chuyển sang hàm tiếp theo trong ngăn xếp
*/
var express = require('express');
var app = express();
var port = 3000;

//1 hàm phần mềm trung gian ko có đường dẫn gắn kết: path
// Hàm được thực hiện mỗi khi app nhận được 1 yêu cầu
app.use((req, res, next) =>{
    console.log("Time: ",Date.now());
    next()
})

// 1 hàm phần mềm trung gian được gắn kết trên đường dẫn: /user/:id
// Hàm được thực hiện cho bất kỳ loại yêu cầu HTTP trên /user/:id
app.use('/user/:id',(req, res, next) =>{
    console.log("Request Type: ",req.method);
    next()
})

//1 tuyến đường và hàm xử lý của nó(hệ thống phần mềm trung gian)
//Hàm xử lý các yêu cầu GET đến đường dẫn: /user/:id
app.use('/user/:id',(req, res, next) =>{
    res.send('USER')
})

app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
})