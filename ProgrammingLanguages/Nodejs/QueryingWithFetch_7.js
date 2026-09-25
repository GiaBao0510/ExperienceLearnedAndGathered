const app= require('express')();

//Máy chủ gửi phản hồi lên trình duyệt cho phép các trang web thuộc bất kỳ nguồn gốc nào yêu cầu trong tệp này
app.get('/',function(req, res){
    res.header("Access-Control-Allow-Origin", "*");
    res.send('Main page');
});


function onTextReady(text){
    console.log(text);
}

function onResponse(response){
    return response.text();
}

fetch('http://localhost:3000/hello')
    .then(onResponse)
    .then(onTextReady);
    
//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log("running");
});