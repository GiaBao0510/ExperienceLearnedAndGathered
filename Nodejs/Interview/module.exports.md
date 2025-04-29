 
### **`module.exports` là gì?**

`module.exports` là cách để export (xuất) dữ liệu (hàm, class, object, biến,...) từ một **file .js/module** có thể **import (nhập)** và sử dụng ở file khác trong Node.js

---
###  **Tại sao lại dùng `module.exports` ?**

Node.js sử dụng hệ thống module gọi là **CommonJS** và mỗi file được coi là **module độc lập**.

Nếu bạn muốn tái sử dụng code ở file khác (ví dụ: function, class, middleware, config,...) bạn **phải xuất nó ra ngoài** bằng `module.exports`

***Ví dụ***
📁 File: `math.js`
```js
function add(a, b) {
    return a + b;
}

module.exports = { add }; // xuất function
```

📁 File: `app.js`
```js
const math = require('./math');

console.log(math.add(2, 3)); // Kết quả: 5
```

---
### **Phân biệt giữa `module.exports` và `exports`** 

Chỉ có `module.exports` là **thật sự được export khi được required** đến nó, Còn `exports` thì không phải, nó chỉ là một **reference** của `module.exports`

```js
//Cài này đúng
exports.SayHi = () => console.log('hi');

//Cái này ghi đè
module.exports = () => console.log('hi');
```

>**⚠️ Lưu ý:** Nếu dùng `module.exports = ...` thì `exports` sẽ không còn hiệu lực nữa, vì `export` chỉ là tham chiếu tắt của `module.exports` ban đầu

---
### **`module.exports` có ảnh hưởng đến hiệu suất không?**

**Không ảnh hưởng đáng kể đến hiệu suất**.

Lý do:

- Node.js **cache các module** sau lần `require()` đầu tiên, vì vậy module chỉ được load và thực thi **một lần duy nhất**.
- Các lần gọi `require(...)` sau đó chỉ **truy cập bản cache**, không đọc lại file từ disk.

> 👉 Vì vậy `module.exports` là cách hiệu quả và thiết kế chuẩn của Node.js để chia nhỏ code mà **không gây tốn hiệu năng**.

---
#### **`module.exports` trong Nodejs thực hiện chức năng gì, cho một ví dụ đơn giản?**

Trong Node JS, `module.exports` là một đối tượng đặt biệt được bao gồm trong mỗi file JavaScript mặc định của ứng dụng NodeJS. `module` là một biến đại diện cho `module` hiện tại, và `exports` là một đối tượng sẽ được tiết lộ như một module. Vì vậy bất cứ thứ gì bạn gán cho `module.exports` sẽ được tiết lộ như là một module.
Mục đính chính của `module.export` là để thực hiện lập trình theo mô-đun. Lập trình theo mô-đun giúp tách biệt các chức năng của chương trình thành các mô-đun độc lập, có thể thay thế, sao cho mỗi mô-đun chứa mọi thứ cần thiết để thực hiện một khía cạnh của chức năng mong muốn. 
Sử dụng `module.export` giúp chúng ta tách biệt logic kinh doanh ra khỏi các mô-đun khác, đồng thời giúp dễ bảo trì và quản lý mã nguồn trong các mô-đun kinh doanh khác nhau.

***Ví dụ đơn giản về cách sử dụng `module.exports`***
- Giả sử bạn có 2 file là `app.js` và `calculator.js` trong thư mục dự án của bạn.
- Trong `calculator.js`, bạn tạo ra một class Calculator với một số phương thức như cộng, trừ, nhân,...
