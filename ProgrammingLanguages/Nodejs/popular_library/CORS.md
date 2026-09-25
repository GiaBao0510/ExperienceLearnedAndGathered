## 🔹 1. **Thư viện `cors` là gì?**

`cors` là viết tắt của **Cross-Origin Resource Sharing** – chia sẻ tài nguyên giữa các nguồn gốc khác nhau.

> 🌐 Nói đơn giản: Nếu frontend và backend **không cùng domain**, thì trình duyệt **chặn** request gọi đến backend.  
> Thư viện `cors` sẽ giúp bạn **cho phép** những request như vậy hoạt động.

---

## 🔹 2. **Tại sao cần `cors` trong dự án?**

### 🔸 Ví dụ tình huống phổ biến:

- Frontend: [http://localhost:3000](http://localhost:3000) (React/Vue/Angular)
- Backend API: [http://localhost:5000](http://localhost:5000) (Express)

➡️ Khi frontend gọi API từ backend khác domain/port, **trình duyệt sẽ chặn** vì lý do bảo mật CORS.

✅ Thư viện `cors` giúp bạn:

- Cho phép client từ domain khác truy cập API
- Có thể cấu hình cho phép cụ thể origin nào, method nào, header nào

## 🔹 3. **Cách áp dụng thư viện `cors` vào dự án Express**

### ✅ Cài đặt:
```bash
npm install cors
```

✅ Tích hợp đơn giản vào `app.js`:
```js
const express = require('express');
const cors = require('cors');

const app = express();

// Cho phép mọi origin (mở hoàn toàn)
app.use(cors());

app.use(express.json());

module.exports = app;
```

>☑️ Tuy nhiên trong môi trường thực tế (production), nên **chỉ cho phép những origin đáng tin cậy** để tránh rủi ro bảo mật.

## 🔹 4. **Ví dụ cấu hình nâng cao (có thể mở rộng & bảo trì)**

### 🎯 Mục tiêu:

- Chỉ cho phép từ domain nhất định
- Cho phép gửi cookie (credentials)
- Tách riêng file cấu hình CORS

##### **📁 Cấu trúc:**
```arduino
src/
├── config/
│   └── cors.config.js      <- Cấu hình cors riêng
├── app.js
```

🔸 `src/config/cors.config.js`
```js
const allowedOrigins = ['http://localhost:3000', 'https://your-frontend.com'];

const corsOptions = {
  origin: function (origin, callback) {
    // Cho phép nếu origin nằm trong danh sách
    if (!origin || allowedOrigins.includes(origin)) {
      callback(null, true);
    } else {
      callback(new Error('Not allowed by CORS'));
    }
  },
  credentials: true, // Cho phép gửi cookie
  optionsSuccessStatus: 200
};

module.exports = corsOptions;
```