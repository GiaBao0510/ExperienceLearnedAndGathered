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

// Phần mềm trung gian cũng có thể được khai báo trong 1 mảng để tái sử dụng

function logOriginalUrl(req, res, next){
    console.log('Request URL: ', req.originalUrl);
    next();
}

function logMethod(req, res, next){
    console.log('Request Type: ', req.method);
    next();
}

const logStuff = [logOriginalUrl, logMethod]
app.get('/user/:id', logStuff,(req, res, next) =>{
    res.send('User Info')
})

app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
});