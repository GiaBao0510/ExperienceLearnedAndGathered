/*
    Viết trình xử lý lỗi
        + Chức năng xử lý lỗi có 4 tham số: (err, req, res, next)
        + Xác định phần mềm trung gian xử lý lỗi cuối cùng, sau khi app.user() khác 
        và định tuyến cuộc gọi
*//*
    Để bỏ qua phần còn lại của các chức năng phần mền trung gian của bộ định tuyến,
    gọi next('router') để chuyển quyền điều khienerra khỏi bộ định tuyến
*/
var express = require('express');
var app = express();
var router = express.Router();
var port = 3000;

app.use('/data', router, function(req, res, next){
    try{
        console.log("Phần mềm trung gian này xử lý truyền dữ liệu")
    }catch(err){
        next(err);
    }
})

app.use(function(err, req, res, next){
    console.error('Error found!');
    res.status(500).send("Đã xảy ra gì đó")
})

app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
});