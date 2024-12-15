const express = require('express');
const app = express();

/*
    Tao ra 1 router de anh xa truy cap den ten mien goc URL 
    su dung phuong thuc the HTTP GET de phan hoi cai ma chung toi 
    muon cung cap .
    Cu phap: app.methob(path, handler);

*/
app.get('/', (req, res)=>{
    res.send('Main page!');
});

//Su dung mot phuong thuc cho moi dong tu HTTP
//1.Phuong thuc GET
app.get('/helloEveryBody', (req, res)=>{
    res.send('Get hello all');
});
//2.Phuong thuc POST
app.post('/helloEveryBody', (req, res)=>{
    res.send('Post hello all');
});
//3.Phuong thuc PUT
app.put('/', (req, res)=>{});
//4.Phuong thuc DELETE
app.delete('/', (req, res)=>{});
//5.Phuong thuc PATCH
app.patch('/', (req, res)=>{});


//Lang nghe tren cong 3000
app.listen(3000, function(){
    console.log('Example app listening on post 3000!');
});