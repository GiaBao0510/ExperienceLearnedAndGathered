# Experience Learned And Gathered

> Kho lưu trữ kiến thức, ghi chú kỹ thuật và các project thực hành trong quá trình học Backend Engineering.

Repository này ghi lại những gì tôi đã học và tích lũy về:

- Backend Development
- Go và C#
- API Architecture
- Database
- Distributed Systems
- System Design
- Cloud và DevOps
- Software Design & Architecture
- Security
- Data Structures & Algorithms

Một số nội dung là ghi chú cá nhân, một số nội dung là bài thực hành hoặc tổng hợp từ quá trình học. Vì vậy, repository này nên được xem như một **learning archive**, không phải toàn bộ nội dung đều là production-ready code.

---

## Mục lục

- [Lộ trình nên đọc](#lộ-trình-nên-đọc)
- [Bản đồ kiến thức](#bản-đồ-kiến-thức)
- [Các project thực hành nổi bật](#các-project-thực-hành-nổi-bật)
- [Interview Preparation](#interview-preparation)
- [Roadmap hiện tại](#roadmap-hiện-tại)
- [Quy ước ghi chép](#quy-ước-ghi-chép)
- [Feedback](#feedback)

---

## Lộ trình nên đọc

Nếu bạn mới bắt đầu xem repository, nên đi theo thứ tự sau:

### Phase 1 — Programming Fundamentals

Nắm các khái niệm nền tảng trước khi học backend:

1. [Go Tutorial](ProgrammingLanguages/Go/Go_Tutorial)
2. [C# Tutorial](ProgrammingLanguages/c-sharp/hocC%23)
3. [Object-Oriented Programming](ProgrammingParadigm/OOP_Object%20Oriented%20Programming)
4. [Functional Programming](ProgrammingParadigm/FP_Functional%20Programming)
5. [Data Structures & Algorithms](DataStructureAlgorithm)

---

### Phase 2 — Backend Fundamentals

Sau khi đã nắm ngôn ngữ lập trình, học các thành phần cốt lõi của backend:

1. [RESTful API](API_Architectures/RESTful_API)
2. [gRPC](API_Architectures/gRPC)
3. [Authentication](Authentication)
4. [Asynchronous Programming](Experience/Asynchronou)
5. [SOLID Principles](Experience/SOLID)
6. [Linux Fundamentals](Linux)

---

### Phase 3 — Database

Database là một trong những trọng tâm chính của repository:

1. [Database Fundamentals](Database/01.DatabaseFundamentals)
2. [Database Design](Database/02.DatabaseDesign)
3. [SQL Fundamentals](Database/03.SQL%20Fundamentals)
4. [Advanced Database](Database/04.AdvancedDatabase)
5. [PostgreSQL](Database/05.Relational%20DBMS/PostgreSQL)
6. [MySQL](Database/05.Relational%20DBMS/MySql)
7. [MongoDB](Database/06.NoSQL/MongoDB)
8. [Redis](Database/06.NoSQL/Redis)
9. [Query Optimization](Database/10.QueryOptimization)
10. [Distributed Database](Database/12.DistributedDatabase)

---

### Phase 4 — Software Design & Architecture

Học cách tổ chức code và thiết kế hệ thống:

1. [Dependency Injection](Software%20Design%20%26%20Architecture/01-Design_Principles/Dependency_Injection-DI)
2. [Design Patterns](Software%20Design%20%26%20Architecture/02-GoF_Design_Pattern)
3. [Layered Architecture](Software%20Design%20%26%20Architecture/03-Architectual_Pattern/Layered-Architecture)
4. [MVC](Software%20Design%20%26%20Architecture/03-Architectual_Pattern/MVC)
5. [Clean Architecture](Software%20Design%20%26%20Architecture/03-Architectual_Pattern/CleanArchitecture)
6. [DDD và Data Patterns](Software%20Design%20%26%20Architecture/04-Data_And_Domain_Pattern)
7. [Microservices](Software%20Design%20%26%20Architecture/System-wide%20Patterns/Microservice)
8. [System Design](systemDesign)

---

### Phase 5 — Distributed Systems và Infrastructure

Sau khi có nền tảng backend và database, tìm hiểu cách xây dựng hệ thống phân tán:

1. [Kafka](Technologies/Kafka)
2. [Elasticsearch](Technologies/Elasticsearch)
3. [Docker](Technologies/Docker)
4. [Loki và Grafana](Technologies/LokiGrafana)
5. [HashiCorp Vault](Technologies/HashingCorpVault)
6. [API Gateway](Technologies/Kong%20API%20Gateway)
7. [AWS](Cloud/AWS)
8. [CI/CD](DevOps/CICD_Guide)

---

## Bản đồ kiến thức

| Nhóm | Nội dung chính |
|---|---|
| [Programming Languages](ProgrammingLanguages) | Go, C#, Node.js, Rust |
| [API Architectures](API_Architectures) | REST, SOAP, gRPC, WebSocket, Webhooks |
| [Authentication](Authentication) | API Key, JWT, OAuth2, token management |
| [Database](Database) | SQL, PostgreSQL, MySQL, MongoDB, Redis, optimization |
| [Data Structures & Algorithms](DataStructureAlgorithm) | Data structures, searching, sorting, string algorithms |
| [Software Design & Architecture](Software%20Design%20%26%20Architecture) | SOLID, GoF patterns, Clean Architecture, DDD, Microservices |
| [Cloud](Cloud) | AWS, Azure, cloud computing |
| [DevOps](DevOps) | Git, GitHub Actions, CI/CD, deployment |
| [Technologies](Technologies) | Docker, Kafka, Elasticsearch, Vault, Grafana, API Gateway |
| [Security](Security) | Cybersecurity, cryptography, data security |
| [Interview Questions](InterviewQuestions) | Backend, database, Go, C#, system design, workflows |

---

## Các project thực hành nổi bật

Đây là những thư mục có source code thực hành đáng xem nhất:

### Go Backend

- [Real-time Backend với WebSocket và SSE](ProgrammingLanguages/Go/GO_web/go_RealTime)
- [REST API với Go](ProgrammingLanguages/Go/GO_web/go_RESTFUL_API)
- [gRPC với Go](ProgrammingLanguages/Go/GO_web/RPC_gRPC)
- [Go kết nối PostgreSQL](ProgrammingLanguages/Go/GO_connect/Go_PostgreSQL)
- [Go kết nối MongoDB](ProgrammingLanguages/Go/GO_connect/Go_MongoDB)
- [Go kết nối Redis](ProgrammingLanguages/Go/GO_connect/Go_Redis)
- [Kafka Producer và Consumer](ProgrammingLanguages/Go/GO_connect/Go_kafka)
- [Elasticsearch REST API](ProgrammingLanguages/Go/GO_connect/GO_Elasticsearch)
- [SQLC](ProgrammingLanguages/Go/GO_lib/1.SQLC)
- [Goose Database Migration](ProgrammingLanguages/Go/GO_lib/2.Goose)

### Authentication và Security

- [API Key Authentication bằng Go](Authentication/API_Key/Demo_api_key_Golang)
- [JWT và Token Management](Authentication/JWT)
- [Password Encryption với Argon2id và bcrypt](ProgrammingLanguages/Go/GO_web/Validation/Go-password-encryption)

### Design Patterns

- [Factory Method bằng Go](Software%20Design%20%26%20Architecture/02-GoF_Design_Pattern/practice/Golang/CreationalPattern_FactoryMethod)
- [Singleton bằng Go](Software%20Design%20%26%20Architecture/02-GoF_Design_Pattern/practice/Golang/CreationalPattern_Singleton)
- [Adapter bằng Go](Software%20Design%20%26%20Architecture/02-GoF_Design_Pattern/practice/Golang/StructualPattern_Adapter)
- [Facade bằng Go](Software%20Design%20%26%20Architecture/02-GoF_Design_Pattern/practice/Golang/StructuralPattern_Facade)

---

## Interview Preparation

Các nội dung chuẩn bị phỏng vấn được phân loại tại:

- [Backend Interview](InterviewQuestions/BackEnd)
- [Go Interview](ProgrammingLanguages/Go/Go_Interview)
- [C# Interview](InterviewQuestions/c-sharp)
- [Database Interview](Database/15.InterviewQuestions)
- [Technical Questions](InterviewQuestions/Technical)
- [Workflow Questions](InterviewQuestions/WorkFlow)
- [Microservices Interview](Software%20Design%20%26%20Architecture/System-wide%20Patterns/Microservice/interview)
- [Git Interview](Git/Interview)

> Interview notes chỉ dùng để ôn tập. Để hiểu sâu hơn, nên đọc thêm phần lý thuyết và source code thực hành tương ứng.

---

## Quy ước ghi chép

Mỗi chủ đề nên cố gắng bao gồm:

1. Khái niệm chính.
2. Vấn đề chủ đề giải quyết.
3. Cách hoạt động bên trong.
4. Ví dụ code hoặc diagram.
5. Trade-off và giới hạn.
6. Trường hợp nên và không nên sử dụng.
7. Link tới project thực hành nếu có.

Các thư mục hoặc file chưa hoàn thiện có thể được tiếp tục cập nhật trong roadmap.

---

## Ghi chú về repository

Repository này được xây dựng chủ yếu cho mục đích:

- Ghi lại quá trình học.
- Ôn tập kiến thức.
- Thực hành công nghệ mới.
- Hệ thống hóa kinh nghiệm cá nhân.
- Chia sẻ kiến thức với cộng đồng.

Không phải tất cả source code trong repository đều được thiết kế để sử dụng trực tiếp trong production. Khi sử dụng lại code, hãy kiểm tra kỹ security, error handling, testing và performance requirements.

---

## Feedback

Nếu bạn phát hiện:

- Nội dung chưa chính xác.
- Link bị hỏng.
- Ví dụ code có lỗi.
- Cách giải thích chưa rõ.
- Chủ đề cần bổ sung.

Bạn có thể mở một Issue hoặc gửi Pull Request. Mọi feedback mang tính xây dựng đều được hoan nghênh.

![](https://www.icegif.com/wp-content/uploads/2023/09/icegif-808.gif)

Thank you for stopping by, and I hope you find something useful here 😎
