const express = require('express');
const app= express();

/*
 Dòng mã này làm cho máy chủ của chúng tôi bắt đầu phân phát 
 trực tiếp các tệp tin trong thư mục 'public' 
*/
app.use(express.static('public'));

app.get('/', function(req, res){
    res.send("Main page!")
})

//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log("running");
});