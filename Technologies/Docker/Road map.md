### **1. Tổng quan về Docker**

- Docker là gì? So sánh với Virtual Machine (VM)
- Kiến trúc của Docker (Docker Engine, Daemon, CLI, API)
- Các thành phần chính: Images, Containers, Volumes, Networks
    

---

### **2. Cài đặt và Cấu hình Docker**

- Cài đặt Docker trên Windows, macOS, Linux
- Kiểm tra phiên bản và xác nhận Docker hoạt động
- Cấu hình Docker Daemon (file `daemon.json`)
    

---

### **3. Làm việc với Docker Images**

- Docker Image là gì? Cấu trúc của Image
- Pull, Search, List Docker Images
- Dockerfile: Viết và Build Image từ Dockerfile
- Docker Hub và Private Registry (Tạo và push Image)

---

### **4. Làm việc với Docker Containers**

- Container là gì? Cách tạo, khởi chạy và dừng Container
- Quản lý Container (start, stop, restart, remove)
- Xem logs, attach và exec vào Container
- Cách lưu trạng thái Container bằng Checkpoint
    

---

### **5. Quản lý Dữ liệu trong Docker**

- Bind Mount vs Volume vs tmpfs mount
- Tạo và sử dụng Volumes trong Docker
- Backup và Restore dữ liệu từ Volume
- Chia sẻ dữ liệu giữa Containers
    

---

### **6. Docker Networking**

- Các loại mạng trong Docker (Bridge, Host, None, Overlay, Macvlan)
- Quản lý mạng với `docker network`
- Kết nối nhiều Container với nhau qua mạng Bridge
- Sử dụng `docker-compose` để tạo mạng riêng cho nhiều Container
    

---

### **7. Docker Compose**

- Giới thiệu về Docker Compose
- Viết file `docker-compose.yml`
- Quản lý nhiều Container với `docker-compose up/down`
- Biến môi trường trong Docker Compose
- Mạng và Volumes trong Docker Compose

---

### **8. Docker Security**

- Người dùng và quyền trong Docker
- Chạy Docker với quyền hạn chế (`--user`)
- Docker Content Trust (DCT)
- Seccomp, AppArmor, SELinux trong Docker
- Quét lỗ hổng bảo mật trong Docker Image (`docker scan`)

---

### **9. Docker Swarm (Orchestration)**

- Giới thiệu về Docker Swarm
- Tạo Swarm Cluster và quản lý Nodes
- Triển khai Service trên Swarm (`docker service`)
- Load Balancing trong Docker Swarm
- Quản lý Persistent Storage trong Swarm

---

### **10. Kubernetes (Tích hợp với Docker)**

- Giới thiệu Kubernetes (K8s) và so sánh với Docker Swarm
- Chạy Docker Container trên Kubernetes
- Kubernetes Pod, Deployment, Service, ConfigMap, Secret
- Sử dụng Minikube và Kubectl

---

### **11. CI/CD với Docker**

- Tích hợp Docker vào quy trình CI/CD
- Sử dụng Docker với GitHub Actions, GitLab CI/CD
- Xây dựng Pipeline tự động hóa với Docker
- Đẩy Image lên Docker Hub hoặc Private Registry

---

### **12. Debug và Tối ưu Docker**

- Xem logs và debug Container
- Giảm dung lượng Docker Image (`multi-stage builds`, `--squash`)
- Tối ưu bộ nhớ và CPU cho Container
- Dọn dẹp Docker (Images, Containers, Volumes không dùng đến)