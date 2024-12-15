const app = require('express')();

/*
    Chúng ta cũng có thể sử dụng cú pháp: VariableName trong đường
    dẫn để chỉ định tham số tuyến đường (Tài liệu Express)
*/
//Ví dụ 1
app.get('/hello/:name',function(req, res){
    const routeParams = req.params;
    const name = routeParams.name;
    res.send("GET: hello, "+name);
    //Chúng ta cũng có thể truy cập các tham số tuyến đường thông qua req.params 
});

//Ví dụ 2: Bạn cũng có thể định nghĩa nhiều tham số tuyến đường trong URL
app.get('/ChuyenXe/:from-:to',function(req, res){
    const routeParams = req.params;
    const from = routeParams.from;
    const to = routeParams.to;
    res.send("GET: Chuyen xe tu "+from +' den '+to);
});

//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log('Running!');
})