
var express = require('express');
var app = express();
var port = 3000;

// Phần mềm trung gian với bộ định tuyến

app.get('/user/:id', 
    function(req, res, next){
        //Nếu user ID là 0 thì đến đường tiếp theo
        if(req.params.id ==="0") next("route");
        //Nếu ko thì vượt qua điều khiển
        //Hàm phần mềm trung gian tiếp theo trong ngăn xếp
        else next(); 
    },
    function(req, res, next){
        //Gửi phản hồi thông thường
        res.send('regular');
    }
)

//Trình xử lý cho đường đẫn: /user/:id, gửi phản hồi đặc biệt
app.get('/user/:id', function(req, res, next){
    res.send('special');
})

app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
});