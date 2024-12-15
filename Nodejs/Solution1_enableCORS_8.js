const app= require('express')();

/*RECALL nghĩa là cách Trang web tự động tạo ra 1 phản hồi để yêu cầu. 
Đây là cách NodeJS Server xử lý các tuyến đường được xác định như sau: */
app.get('/hello', function(req, res){
    res.send("GET hello!")
})

//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log("running");
});