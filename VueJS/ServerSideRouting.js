
const express = require('express');
const path = require("path");

const app = express();

//Hàm này được gọi khi có yêu cầu GET, nào đó đến server mà không khớp với các route cụ thể khác
async function onAllOtherPaths(req, res){
    //res.sendFile: Gửi file index.html nằm trong thư mục public về phía client.
    //path.resolve(__dirname, 'public', 'index.html'): Sử dụng hàm path.resolve để tạo đường dẫn tuyệt đối đến file index.html.
    //__dirname: Biến đặc biệt trong Node.js, chứa đường dẫn đến thư mục hiện tại của file đang chạy.
    res.sendFile(path.resolve(__dirname, 'public', 'index.html'));
}

/*
     Sử dụng app.get để đăng ký một route chung cho tất cả các yêu cầu GET (*), 
    chỉ định hàm onAllOtherPaths sẽ được gọi khi có yêu cầu khớp với route này.
*/
app.get('*', onAllOtherPaths);


//Lắng nghe trên cổng 3000
app.listen(3000, function(){
    console.log('Running...');
})
