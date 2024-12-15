const express = require('express');
const app = express();

//Các yêu cầu của kiểu nội dung application
app.use(express.json());

//Phân tích các yêu cầu của ứng dụng loại nội dung/x-www-form-ủlencoded
app.use(express.urlencoded({extended: true}));

app.get('/hello',function(req, res){
    const body = req.body;
    const name = body.name;
    const email = body.email;
    res.send("POST: Name: "+name+" ,Email: "+email);
})

//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log('Running!'); 
})