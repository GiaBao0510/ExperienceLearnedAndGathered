const {MongoClient} = require('mongodb');
require('mongodb').MongoClient;

//ĐƯờng dẫn để liên kết đến cơ sở dữ liệu trên đám mây. Gồm tên tài khoản và mật khẩu
const pathCloud = "mongodb+srv://giabaodev:<password>@pgiabaodev.vtpstl5.mongodb.net/?retryWrites=true&w=majority&appName=pgiabaodev";

let client =null,
    db = null,
    collection = null,
    collection2 = null;

async function main(){
    client = await MongoClient.connect(path);       //Thực hiện kết nối đến cơ sở dữ liệu
    db = await client.db();                         //Lấy cơ sở dữ liệu
    collection = await db.collection('Bảng1');      //Kết nối đến bảng 1
    collection2 = await db.collection('Bảng2');
   
}

main();