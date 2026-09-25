const { default: axios } = require('axios');
const express = require('express');
const app = express();

/*
    Trong Express:
        * Các tham số Router về cơ bản là các biến bắt nguồn từ URL của các phần được đặt tên
        * Epress nắm bắt giá trị trong phần được đặt tên và lưu nó trong thuộc tính req.params
*/

app.get('/user/:userId/books/:bookId', (req, res) => {
    req.params; // userId: 42, bookId: 101
    res.json(req.params); // Gửi response bằng res.json()
})


app.listen(3000, async ()=>{
    //Khởi động server trước 
    console.log("Server listening in port 3000");

    // Thực hiện request axios sau khi server đã khởi động
    const res = await axios.get('http://localhost:3000/user/42/books/101');
    console.log(res.data);
});
