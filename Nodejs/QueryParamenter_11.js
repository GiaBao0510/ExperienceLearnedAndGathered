const app = require('express')();

/*
    Chúng ta có thể truy cập đến tham số truy vấn thông qua req.query
*/
//Ví dụ 1
app.get('/hello/:name',function(req, res){
    const queryParams = req.query;
    const name = queryParams.name;
    res.send("GET: hello, "+name);
    //Chúng ta cũng có thể truy cập các tham số truy vấn thông qua req.query 
});



//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log('Running!');
})