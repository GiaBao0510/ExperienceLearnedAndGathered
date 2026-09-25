## 🔹 1. **Thư viện `express` là gì?**

`express` là một **framework web nhẹ** dành cho Node.js. Nó giúp bạn dễ dàng xây dựng các ứng dụng web hoặc API một cách nhanh chóng, sạch sẽ, và có tổ chức hơn.

> ✅ Có thể hiểu đơn giản: Nếu Node.js là động cơ thì Express chính là khung sườn của chiếc xe web app.

---

## 🔹 2. **Tại sao cần `express` trong dự án?**

Mặc dù bạn có thể dùng Node.js thuần (native) để tạo server, nhưng:

- Quá trình xử lý request/response rất **thủ công và cồng kềnh**.
- Không có sẵn những thứ như routing, middleware, hay parsing JSON.

**`Express` giúp bạn:**

| Tính năng                | Lợi ích                                                         |
| ------------------------ | --------------------------------------------------------------- |
| **Routing**              | Dễ dàng định nghĩa các API endpoint                             |
| **Middleware**           | Gắn thêm các bước xử lý cho request (kiểm tra token, log, v.v.) |
| **Body Parser**          | Tự động parse JSON/body từ client gửi lên                       |
| **Cấu trúc mở rộng tốt** | Dễ tách nhỏ thành nhiều module                                  |
| **Hệ sinh thái mạnh**    | Dễ tích hợp với MongoDB, Redis, Auth0, v.v.                     |

## 🔹 3. **Cách áp dụng thư viện `express` vào dự án (theo hướng có thể mở rộng & bảo trì)**

### ✅ Cấu trúc dự án gợi ý (theo hướng có tổ chức):
```pgsql
project/
├── src/
│   ├── app.js              <- Khởi tạo app express
│   ├── server.js           <- File chạy server
│   ├── routes/
│   │   └── user.route.js   <- Các route cho user
│   ├── controllers/
│   │   └── user.controller.js
│   ├── middlewares/
│   │   └── auth.middleware.js
│   └── services/
│       └── user.service.js
└── package.json
```

### ✅ Ví dụ mã code áp dụng `express`:

#### 🔸 `src/app.js` – Tạo app Express
```js
const express = require('express');
const app = express();
const userRoutes = require('./routes/user.route');

app.use(express.json()); // Middleware để parse JSON
app.use('/api/users', userRoutes); // Mount route

module.exports = app;
```

#### 🔸 `src/server.js` – Chạy server
```js
const app = require('./app');

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
```

🔸 `src/routes/user.route.js` – Định nghĩa các route
```js
const express = require('express');
const router = express.Router();
const userController = require('../controllers/user.controller');

router.get('/', userController.getAllUsers);
router.post('/', userController.createUser);

module.exports = router;
```

🔸 `src/controllers/user.controller.js` – Xử lý logic route
```js
const userService = require('../services/user.service');

exports.getAllUsers = (req, res) => {
  const users = userService.getAll();
  res.json(users);
};

exports.createUser = (req, res) => {
  const newUser = userService.create(req.body);
  res.status(201).json(newUser);
};
```

🔸 `src/services/user.service.js` – Tầng logic (business)
```js
let users = [];

exports.getAll = () => users;

exports.create = (data) => {
  const newUser = { id: Date.now(), ...data };
  users.push(newUser);
  return newUser;
};
```

## 🔹 4. **Kết luận**

- `express` giúp xây dựng server nhanh chóng, dễ bảo trì, mở rộng.
- Dùng đúng cấu trúc và tách các phần rõ ràng (route, controller, service) sẽ giúp dự án của bạn:
    - **Dễ scale**
    - **Dễ test**
    - **Dễ làm việc nhóm**