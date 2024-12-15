/*
     Để bỏ qua phần còn lại của các chức năng phần mền trung gian của bộ định tuyến,
    gọi next('router') để chuyển quyền điều khiển ra khỏi bộ định tuyến
*/
var express = require('express');
var app = express();
var router = express.Router();
var port = 3000;

//Xác định bộ định tuyến bằng cách kiểm tra và bảo lãnh khi cần
router.use(function(req, res, next){
    if(!req.header["x-auth"]) return next("router");
    next();
})

router.get('/user/:id',function(req, res){
    res.send("hello, user!");
})

//dùng bộ định tuyến và 401 bất cứ thứ gì rơi vào
app.use('/admin', router, function(req, res){
    res.sendStatus(401);
})

app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
});